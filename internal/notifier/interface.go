package notifier

import "time"

type ActivityProvider interface {
	UpdateLastSeen(userID uint, seenAt time.Time) error
	GetLastSeenForUsers(userIDs []uint) (map[uint]time.Time, error)
}

type MatchProvider interface {
	GetMatchUserIDs(uint) ([]uint, error)
}
