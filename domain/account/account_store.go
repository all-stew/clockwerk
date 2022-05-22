package account

import (
	. "clockwerk/config/mysql"
	"clockwerk/pkg/logger"
	"clockwerk/pkg/util"
	"errors"
	"gorm.io/gorm"
)

type Store struct {
}

var _ Repository = (*Store)(nil)

func (s Store) Create(account string, accountType AccountType, parentId uint64, accountParam string, userId uint64, createdBy uint64) bool {
	logger.Logf("job_store Create account:%s accountType:%d, parentId:%d, accountParam:%s createdBy:%d",
		account, accountType, parentId, accountParam, createdBy)

	if !util.IsJsonString(accountParam) {
		return false
	}

	ac := Account{
		Account:      account,
		AccountType:  accountType,
		ParentId:     parentId,
		AccountParam: accountParam,
		UserId:       userId,
		CreatedBy:    createdBy,
	}

	result := GetDb().Create(&ac)

	if result.Error != nil {
		logger.Logf("create error %s\n", result.Error.Error())
		return false
	}

	return result.RowsAffected == 1
}

func (s Store) Update(id uint64, accountParam string, updatedBy uint64) bool {
	logger.Logf("job_store Update id:%d accountParam:%s updatedBy:%d",
		id, accountParam, updatedBy)

	if !util.IsJsonString(accountParam) {
		return false
	}

	ac := Account{}
	resultForFind := GetDb().First(&ac, id)
	if errors.Is(resultForFind.Error, gorm.ErrRecordNotFound) {
		logger.Logf("cannot find error %s\n", resultForFind.Error.Error())
		return false
	}

	ac.AccountParam = accountParam
	ac.UpdatedBy = updatedBy

	result := GetDb().Save(&ac)
	if result.Error != nil {
		logger.Logf("update error %s\n", result.Error.Error())
	}

	return result.RowsAffected == 1

}
