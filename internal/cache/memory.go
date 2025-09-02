package cache

import (
	"dating_service/internal/model"
	"sync"
)

type ReferenceCache struct {
	mu sync.RWMutex

	sexIDs               map[uint]model.Sex
	educationIDs         map[uint]model.Education
	zodiacSignIDs        map[uint]model.ZodiacSign
	worldviewIDs         map[uint]model.Worldview
	typeOfDatingIDs      map[uint]model.TypeOfDating
	attitudeToAlcoholIDs map[uint]model.AttitudeToAlcohol
	attitudeToSmokingIDs map[uint]model.AttitudeToSmoking
	statusIDs            map[uint]model.Status
	interestIDs          map[uint]model.Interest
}
