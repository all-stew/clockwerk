package user

type Repository interface {
	Create(username string, nickname string, password string, email string, phone string, createBy uint64) bool
	Update(id uint64, nickname string, email string, phone string, updatedBy uint64) bool
}
