package job

import (
	. "clockwerk/config/mysql"
	"clockwerk/pkg/logger"
)

type Store struct {
}

var _ Repository = (*Store)(nil)

func (s Store) Create(accountId uint64, userNotificationId uint64, userId uint64, timeWindow int, createdBy uint64) bool {
	logger.Logf("job_store Create accountId:%d userNotificationId:%d, userId:%d, timeWindow:%d createdBy:%d",
		accountId, userNotificationId, userId, timeWindow, createdBy)

	// todo 校验

	job := Job{
		AccountId:          accountId,
		UserNotificationId: userNotificationId,
		UserId:             userId,
		TimeWindow:         timeWindow,
		CreatedBy:          createdBy,
	}

	result := GetDb().Create(&job)

	if result.Error != nil {
		logger.Logf("create error %s\n", result.Error.Error())
	}

	return result.RowsAffected == 1
}
