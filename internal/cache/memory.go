package cache

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"log"
	"sync"
)

type referenceCache struct {
	mu sync.RWMutex

	sexIDs               map[uint]struct{}
	educationIDs         map[uint]struct{}
	zodiacSignIDs        map[uint]struct{}
	worldviewIDs         map[uint]struct{}
	typeOfDatingIDs      map[uint]struct{}
	attitudeToAlcoholIDs map[uint]struct{}
	attitudeToSmokingIDs map[uint]struct{}
	statusIDs            map[uint]struct{}
	interestIDs          map[uint]struct{}
}

func NewReferenceCache(db *db.Db) (IReferenceCache, error) {

	var sexes []model.Sex
	if err := db.PgDb.Find(&sexes).Error; err != nil {
		return nil, err
	}
	sexMap := make(map[uint]struct{}, len(sexes))
	for _, item := range sexes {
		sexMap[item.ID] = struct{}{}
	}

	var educations []model.Education
	if err := db.PgDb.Find(&educations).Error; err != nil {
		return nil, err
	}
	eduMap := make(map[uint]struct{}, len(educations))
	for _, item := range educations {
		eduMap[item.ID] = struct{}{}
	}

	var zodiacSigns []model.ZodiacSign
	if err := db.PgDb.Find(&zodiacSigns).Error; err != nil {
		return nil, err
	}
	zodiacMap := make(map[uint]struct{}, len(zodiacSigns))
	for _, item := range zodiacSigns {
		zodiacMap[item.ID] = struct{}{}
	}

	var worldviews []model.Worldview
	if err := db.PgDb.Find(&worldviews).Error; err != nil {
		return nil, err
	}
	worldviewMap := make(map[uint]struct{}, len(worldviews))
	for _, item := range worldviews {
		worldviewMap[item.ID] = struct{}{}
	}

	var typesOfDating []model.TypeOfDating
	if err := db.PgDb.Find(&typesOfDating).Error; err != nil {
		return nil, err
	}
	datingMap := make(map[uint]struct{}, len(typesOfDating))
	for _, item := range typesOfDating {
		datingMap[item.ID] = struct{}{}
	}

	var alcoholAttitudes []model.AttitudeToAlcohol
	if err := db.PgDb.Find(&alcoholAttitudes).Error; err != nil {
		return nil, err
	}
	alcoholMap := make(map[uint]struct{}, len(alcoholAttitudes))
	for _, item := range alcoholAttitudes {
		alcoholMap[item.ID] = struct{}{}
	}

	var smokingAttitudes []model.AttitudeToSmoking
	if err := db.PgDb.Find(&smokingAttitudes).Error; err != nil {
		return nil, err
	}
	smokingMap := make(map[uint]struct{}, len(smokingAttitudes))
	for _, item := range smokingAttitudes {
		smokingMap[item.ID] = struct{}{}
	}

	var statuses []model.Status
	if err := db.PgDb.Find(&statuses).Error; err != nil {
		return nil, err
	}
	statusMap := make(map[uint]struct{}, len(statuses))
	for _, item := range statuses {
		statusMap[item.ID] = struct{}{}
	}

	var interests []model.Interest
	if err := db.PgDb.Find(&interests).Error; err != nil {
		return nil, err
	}
	interestMap := make(map[uint]struct{}, len(interests))
	for _, item := range interests {
		interestMap[item.ID] = struct{}{}
	}

	log.Println("The directory cache has been loaded successfully")

	return &referenceCache{
		sexIDs:               sexMap,
		educationIDs:         eduMap,
		zodiacSignIDs:        zodiacMap,
		worldviewIDs:         worldviewMap,
		typeOfDatingIDs:      datingMap,
		attitudeToAlcoholIDs: alcoholMap,
		attitudeToSmokingIDs: smokingMap,
		statusIDs:            statusMap,
		interestIDs:          interestMap,
	}, nil
}

func (c *referenceCache) IsValidSexID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.sexIDs[id]
	return exists
}

func (c *referenceCache) IsValidEducationID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.educationIDs[id]
	return exists
}

func (c *referenceCache) IsValidZodiacSignID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.zodiacSignIDs[id]
	return exists
}

func (c *referenceCache) IsValidWorldviewID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.worldviewIDs[id]
	return exists
}

func (c *referenceCache) IsValidTypeOfDatingID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.typeOfDatingIDs[id]
	return exists
}

func (c *referenceCache) IsValidAttitudeToAlcoholID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.attitudeToAlcoholIDs[id]
	return exists
}

func (c *referenceCache) IsValidAttitudeToSmokingID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.attitudeToSmokingIDs[id]
	return exists
}

func (c *referenceCache) IsValidStatusID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.statusIDs[id]
	return exists
}

func (c *referenceCache) IsValidInterestIDs(ids []uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, id := range ids {
		if _, exists := c.interestIDs[id]; !exists {
			return false
		}
	}

	return true
}
