package notification

// Repository user_notification的数据仓库接口
type Repository interface {
	Create(notificationType NotificationType, notificationKey string, notificationParam string, userId uint64, createBy uint64) bool
	Update(id uint64, notificationType NotificationType, notificationKey string, notificationParam string, userId uint64, updatedBy uint64) bool
}
