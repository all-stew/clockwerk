package notification

import (
	. "clockwerk/config/mysql"
	"clockwerk/pkg/logger"
	"encoding/json"
)

type Store struct {
}

var _ Repository = (*Store)(nil)

func (s Store) Create(notificationType NotificationType, notificationKey string, notificationParam string, userId uint64, createBy uint64) bool {

	logger.Logf("user_notification_store Create notificationType:%d notificationKey:%s notificationParam:%s userId:%d createBy:%d",
		notificationType, notificationKey, notificationParam, userId, createBy)

	if json.Valid([]byte(notificationParam)) {
		return false
	}

	userNotification := UserNotification{
		NotificationType:  notificationType,
		NotificationKey:   notificationKey,
		NotificationParam: notificationParam,
		UserId:            userId,
		CreatedBy:         createBy,
	}

	result := GetDb().Create(&userNotification)

	if result.Error != nil {
		logger.Logf("create error %s\n", result.Error.Error())
	}

	return result.RowsAffected == 1
}

func (s Store) Update(id uint64, notificationType NotificationType, notificationKey string, notificationParam string, userId uint64, updatedBy uint64) bool {

	logger.Logf("user_notification_store Update id:%d notificationType:%d notificationKey:%s notificationParam:%s userId:%d updatedBy:%d",
		id, notificationType, notificationKey, notificationParam, userId, updatedBy)

	if json.Valid([]byte(notificationParam)) {
		return false
	}

	userNotification := UserNotification{}
	GetDb().Where("id = ?", id).Take(&userNotification)
	userNotification.NotificationType = notificationType
	userNotification.NotificationKey = notificationKey
	userNotification.NotificationParam = notificationParam
	userNotification.UserId = userId
	userNotification.UpdatedBy = updatedBy

	result := GetDb().Save(&userNotification)
	if result.Error != nil {
		logger.Logf("update error %s\n", result.Error.Error())
	}

	return result.RowsAffected == 1
}
