package activity

import "time"

type ActivityService struct {
	activityStorage ActivityStorage
}

func NewActivityService(activityStorage ActivityStorage) *ActivityService {
	return &ActivityService{activityStorage}
}

func (service *ActivityService) UpdateLastSeen(userID uint, seenAt time.Time) error {
	return service.activityStorage.UpdateLastSeen(userID, seenAt)
}
func (service *ActivityService) GetLastSeenForUsers(userIDs []uint) (map[uint]time.Time, error) {
	return service.activityStorage.GetLastSeenForUsers(userIDs)
}
