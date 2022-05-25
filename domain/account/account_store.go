package account

import (
	"clockwerk/api/genshin"
	. "clockwerk/config/mysql"
	"clockwerk/pkg/logger"
	"clockwerk/pkg/util"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Store struct {
}

var _ Repository = (*Store)(nil)

func (s Store) Create(account string, accountType AccountType, parentId uint64, accountParam string, userId uint64, createdBy uint64) bool {
	logger.Logf("account_store Create genshin_account:%s accountType:%d, parentId:%d, accountParam:%s createdBy:%d",
		account, accountType, parentId, accountParam, createdBy)

	if !util.IsJsonString(accountParam) {
		return false
	}

	result := s.FindByAccountAndAccountType(account, accountType)
	if result.Error == nil {
		logger.Logf("account %s accountType %d exists", account, accountType)
		return false
	}

	ac := create(account, accountType, parentId, accountParam, ACCOUNT_DISABLE, userId, createdBy)

	var status Status
	switch accountType {
	case ACCOUNT_TYPE_MIHOYO_BBS:
		status = createGenshinAccount(ac, accountParam, userId)
	}
	s.UpdateStatus(ac.ID, status)

	return true
}

func create(account string, accountType AccountType, parentId uint64, accountParam string, status Status, userId uint64, createdBy uint64) Account {
	ac := Account{
		Account:      account,
		AccountType:  accountType,
		ParentId:     parentId,
		AccountParam: accountParam,
		UserId:       userId,
		Status:       status,
		CreatedBy:    createdBy,
	}

	result := GetDb().Create(&ac)

	if result.Error != nil {
		logger.Logf("create error %s\n", result.Error.Error())
		return Account{}
	}

	return ac
}

func createGenshinAccount(account Account, accountParam string, userId uint64) Status {
	var status Status

	cookieObj, err := genshin.GetGenshinCookie(accountParam)
	if err != nil {
		logger.Logf("get genshin cookie error %s\n", err.Error())
		status = ACCOUNT_MIHOYO_ERROR
	} else {
		cookieStr := fmt.Sprintf("cookie_token=%s; account_id=%s", cookieObj.CookieToken, cookieObj.AccountId)
		resp, err := genshin.GetUserGameRoleByCookie(cookieStr)
		if err != nil {
			logger.Logf("get genshin genshin_account error %s\n", err.Error())
			status = ACCOUNT_MIHOYO_ERROR
		} else {
			respStr := resp.String()
			var genshinAccountResp genshin.GenshinAccountResp
			err = json.Unmarshal([]byte(respStr), &genshinAccountResp)
			if err != nil {
				logger.Logf("get genshin genshin_account json error %s\n", err.Error())
				status = ACCOUNT_MIHOYO_ERROR
			}
			if len(genshinAccountResp.Data.List) > 0 {
				for _, genshinAccount := range genshinAccountResp.Data.List {
					b, _ := json.Marshal(genshinAccount)
					create(genshinAccount.GameUid, ACCOUNT_TYPE_GENSHIN, account.ID, string(b), ACCOUNT_GENSHIN_VALID, userId, 0)
					logger.Logf("get genshin genshin_account %s\n", b)
				}
				status = ACCOUNT_ENABLE
			} else {
				logger.Log("genshin genshin_account size is 0\n")
				status = ACCOUNT_MIHOYO_DO_NOT_HAVE_GENSHIN
			}
		}
	}
	return status
}

func (s Store) Update(id uint64, accountParam string, updatedBy uint64) bool {
	logger.Logf("account_store Update id:%d accountParam:%s updatedBy:%d",
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

func (s Store) UpdateStatus(id uint64, status Status) {
	ac := Account{}
	ac.ID = id
	GetDb().Model(&ac).Update("status", status)
}

func (s Store) ListByStatus(status Status) []Account {
	logger.Logf("job_store ListByStatus status:%d", status)

	var accounts []Account

	result := GetDb().Where("status = ?", status).Find(&accounts)

	if result.Error != nil {
		logger.Logf("ListByStatus error %s\n", result.Error.Error())
		return accounts
	}
	return accounts
}

func (s Store) FindByAccountAndAccountType(account string, accountType AccountType) *gorm.DB {
	var ac Account
	result := GetDb().Where(&Account{Account: account, AccountType: accountType}).First(&ac)
	return result
}
