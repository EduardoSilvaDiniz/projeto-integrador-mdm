package repositories

// type GormAssociatedRepository struct {
// 	DB *gorm.DB
// }
//
// func NewGormAssociatedRepository(db *gorm.DB) *GormAssociatedRepository {
// 	return &GormAssociatedRepository{DB: db}
// }

// func (r *GormAssociatedRepository) Save(a *entities.Associated) error {
// 	return r.DB.Create(a).Error
// }
//
// func (r *GormAssociatedRepository) DeleteByCPF(cpf string) error {
// 	return r.DB.Where("cpf = ?", cpf).Delete(&entities.Associated{}).Error
// }
//
// func (r *GormAssociatedRepository) GetAll() ([]dto.Associated, error) {
// 	var assoc []dto.Associated
// 	return assoc, errors.New("falta implementação")
// }
