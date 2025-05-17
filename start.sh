#!/usr/bin/env bash

SERVER_CMD="air"
DOCKER_FILE="./deployments/docker-compose.yml"
SERVER_PID=0
MENU="
✅ Servidor iniciado com PID: ${SERVER_PID}
🕹️ Comandos disponíveis:
  [R] Reiniciar o servidor
  [Q] Fechar o servidor
"

start_server() {
	echo "🚀 Iniciando o servidor..."
	air &
	SERVER_PID=$!

	echo "[INFO] Aguardando servidor iniciar..."
	sleep 3

	if ! curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/ping | grep -q "200"; then
		echo "Servidor falhou na inicialização"
		docker-compose -f "$DOCKER_FILE" down
		stop_server
		exit 1
	fi
	echo "$MENU"
}

stop_server() {
	if [[ $SERVER_PID -ne 0 ]]; then
		echo "🛑 Parando o servidor com PID: $SERVER_PID"
		kill "$SERVER_PID"
		SERVER_PID=0
	fi
}

restart_server() {
	stop_server
	start_server
	echo "$MENU"
}

trap_ctrl_c() {
	echo "⛔ Encerrando por Ctrl+C..."
	stop_server
	exit 0
}

trap trap_ctrl_c INT

docker-compose -f "$DOCKER_FILE" up -d
sleep 4
start_server

while true; do
	read -n 1 -s key
	case $key in
	[Rr])
		stop_server
		start_server
		;;
	[Qq])
		stop_server
		docker-compose -f "$DOCKER_FILE" down
		break
		;;
	esac
done
