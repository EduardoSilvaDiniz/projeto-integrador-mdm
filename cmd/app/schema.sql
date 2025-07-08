-- necessario para o go:embed
CREATE TABLE associated (
	id INTEGER PRIMARY KEY, 
	cpf TEXT NOT NULL,
  name TEXT NOT NULL,
  date_birth TEXT NOT NULL,
  marital_status TEXT NOT NULL
);
