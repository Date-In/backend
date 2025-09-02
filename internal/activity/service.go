package activity

import "time"

type Service struct {
	activityStorage ActivityStorage
}

func NewService(activityStorage ActivityStorage) *Service {
	return &Service{activityStorage}
}

func (s *Service) UpdateLastSeen(userID uint, seenAt time.Time) error {
	return s.activityStorage.UpdateLastSeen(userID, seenAt)
}
func (s *Service) GetLastSeenForUsers(userIDs []uint) (map[uint]time.Time, error) {
	return s.activityStorage.GetLastSeenForUsers(userIDs)
}
