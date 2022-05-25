package account

import "gorm.io/gorm"

type Repository interface {
	Create(account string, accountType AccountType, parentId uint64, accountParam string, userId uint64, createdBy uint64) bool
	Update(id uint64, accountParam string, updatedBy uint64) bool
	ListByStatus(status Status) []Account
	FindByAccountAndAccountType(account string, accountType AccountType) *gorm.DB
	UpdateStatus(id uint64, status Status)
}
