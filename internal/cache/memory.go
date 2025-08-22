package cache

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"log"
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

func NewReferenceCache(db *db.Db) (*ReferenceCache, error) {
	var sexes []model.Sex
	if err := db.PgDb.Find(&sexes).Error; err != nil {
		return nil, err
	}
	sexMap := make(map[uint]model.Sex, len(sexes))
	for _, item := range sexes {
		sexMap[item.ID] = item
	}

	var educations []model.Education
	if err := db.PgDb.Find(&educations).Error; err != nil {
		return nil, err
	}
	eduMap := make(map[uint]model.Education, len(educations))
	for _, item := range educations {
		eduMap[item.ID] = item
	}

	var zodiacSigns []model.ZodiacSign
	if err := db.PgDb.Find(&zodiacSigns).Error; err != nil {
		return nil, err
	}
	zodiacMap := make(map[uint]model.ZodiacSign, len(zodiacSigns))
	for _, item := range zodiacSigns {
		zodiacMap[item.ID] = item
	}

	var worldviews []model.Worldview
	if err := db.PgDb.Find(&worldviews).Error; err != nil {
		return nil, err
	}
	worldviewMap := make(map[uint]model.Worldview, len(worldviews))
	for _, item := range worldviews {
		worldviewMap[item.ID] = item
	}

	var typesOfDating []model.TypeOfDating
	if err := db.PgDb.Find(&typesOfDating).Error; err != nil {
		return nil, err
	}
	datingMap := make(map[uint]model.TypeOfDating, len(typesOfDating))
	for _, item := range typesOfDating {
		datingMap[item.ID] = item
	}

	var alcoholAttitudes []model.AttitudeToAlcohol
	if err := db.PgDb.Find(&alcoholAttitudes).Error; err != nil {
		return nil, err
	}
	alcoholMap := make(map[uint]model.AttitudeToAlcohol, len(alcoholAttitudes))
	for _, item := range alcoholAttitudes {
		alcoholMap[item.ID] = item
	}

	var smokingAttitudes []model.AttitudeToSmoking
	if err := db.PgDb.Find(&smokingAttitudes).Error; err != nil {
		return nil, err
	}
	smokingMap := make(map[uint]model.AttitudeToSmoking, len(smokingAttitudes))
	for _, item := range smokingAttitudes {
		smokingMap[item.ID] = item
	}

	var statuses []model.Status
	if err := db.PgDb.Find(&statuses).Error; err != nil {
		return nil, err
	}
	statusMap := make(map[uint]model.Status, len(statuses))
	for _, item := range statuses {
		statusMap[item.ID] = item
	}

	var interests []model.Interest
	if err := db.PgDb.Find(&interests).Error; err != nil {
		return nil, err
	}
	interestMap := make(map[uint]model.Interest, len(interests))
	for _, item := range interests {
		interestMap[item.ID] = item
	}

	log.Println("The directory cache has been loaded successfully")

	return &ReferenceCache{
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

func (c *ReferenceCache) IsValidSexID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.sexIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidEducationID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.educationIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidZodiacSignID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.zodiacSignIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidWorldviewID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.worldviewIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidTypeOfDatingID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.typeOfDatingIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidAttitudeToAlcoholID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.attitudeToAlcoholIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidAttitudeToSmokingID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.attitudeToSmokingIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidStatusID(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.statusIDs[id]
	return exists
}

func (c *ReferenceCache) IsValidInterestIDs(ids []uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, id := range ids {
		if _, exists := c.interestIDs[id]; !exists {
			return false
		}
	}

	return true
}

func (c *ReferenceCache) IsValidInterest(id uint) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.interestIDs[id]
	return exists
}

func (c *ReferenceCache) GetSexByID(id uint) model.Sex {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.sexIDs[id]
}

func (c *ReferenceCache) GetEducationByID(id uint) model.Education {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.educationIDs[id]
}

func (c *ReferenceCache) GetZodiacSignByID(id uint) model.ZodiacSign {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.zodiacSignIDs[id]
}

func (c *ReferenceCache) GetWorldviewByID(id uint) model.Worldview {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.worldviewIDs[id]
}

func (c *ReferenceCache) GetTypeOfDatingByID(id uint) model.TypeOfDating {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.typeOfDatingIDs[id]
}

func (c *ReferenceCache) GetAttitudeToAlcoholByID(id uint) model.AttitudeToAlcohol {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.attitudeToAlcoholIDs[id]
}

func (c *ReferenceCache) GetAttitudeToSmokingByID(id uint) model.AttitudeToSmoking {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.attitudeToSmokingIDs[id]
}

func (c *ReferenceCache) GetStatusByID(id uint) model.Status {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.statusIDs[id]
}

func (c *ReferenceCache) GetInterestByID(id uint) model.Interest {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.interestIDs[id]
}
