package activity

import "time"

type ActivityStorage interface {
	UpdateLastSeen(userID uint, seenAt time.Time) error
	GetLastSeenForUsers(userIDs []uint) (map[uint]time.Time, error)
}

type MatchProvider interface {
	GetMatchUserIDs(uint) ([]uint, error)
}
