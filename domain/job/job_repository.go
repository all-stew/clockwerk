package job

type Repository interface {
	Create(accountId uint64, userNotificationId uint64, userId uint64, timeWindow int, createdBy uint64) bool
}
