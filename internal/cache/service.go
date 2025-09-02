package cache

import (
	"dating_service/internal/model"
	"log"
)

func NewReferenceCache(repo *Repository) (*ReferenceCache, error) {
	sexMap, err := repo.LoadSexes()
	if err != nil {
		return nil, err
	}

	eduMap, err := repo.LoadEducations()
	if err != nil {
		return nil, err
	}

	zodiacMap, err := repo.LoadZodiacSigns()
	if err != nil {
		return nil, err
	}

	worldviewMap, err := repo.LoadWorldviews()
	if err != nil {
		return nil, err
	}

	datingMap, err := repo.LoadTypeOfDating()
	if err != nil {
		return nil, err
	}

	alcoholMap, err := repo.LoadAttitudeToAlcohol()
	if err != nil {
		return nil, err
	}

	smokingMap, err := repo.LoadAttitudeToSmoking()
	if err != nil {
		return nil, err
	}

	statusMap, err := repo.LoadStatuses()
	if err != nil {
		return nil, err
	}

	interestMap, err := repo.LoadInterests()
	if err != nil {
		return nil, err
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

// новые методы (списки)
func (c *ReferenceCache) GetSexes() ([]model.Sex, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.Sex, 0, len(c.sexIDs))
	for _, v := range c.sexIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetEducations() ([]model.Education, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.Education, 0, len(c.educationIDs))
	for _, v := range c.educationIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetZodiacSigns() ([]model.ZodiacSign, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.ZodiacSign, 0, len(c.zodiacSignIDs))
	for _, v := range c.zodiacSignIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetWorldViews() ([]model.Worldview, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.Worldview, 0, len(c.worldviewIDs))
	for _, v := range c.worldviewIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetTypeOfDating() ([]model.TypeOfDating, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.TypeOfDating, 0, len(c.typeOfDatingIDs))
	for _, v := range c.typeOfDatingIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetAttitudeToAlcohol() ([]model.AttitudeToAlcohol, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.AttitudeToAlcohol, 0, len(c.attitudeToAlcoholIDs))
	for _, v := range c.attitudeToAlcoholIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetAttitudeToSmoking() ([]model.AttitudeToSmoking, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.AttitudeToSmoking, 0, len(c.attitudeToSmokingIDs))
	for _, v := range c.attitudeToSmokingIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetInterests() ([]model.Interest, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.Interest, 0, len(c.interestIDs))
	for _, v := range c.interestIDs {
		list = append(list, v)
	}
	return list, nil
}

func (c *ReferenceCache) GetStatuses() ([]model.Status, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]model.Status, 0, len(c.statusIDs))
	for _, v := range c.statusIDs {
		list = append(list, v)
	}
	return list, nil
}
