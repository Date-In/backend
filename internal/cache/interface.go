package cache

type IReferenceCache interface {
	IsValidSexID(id uint) bool

	IsValidEducationID(id uint) bool

	IsValidZodiacSignID(id uint) bool

	IsValidWorldviewID(id uint) bool

	IsValidTypeOfDatingID(id uint) bool

	IsValidAttitudeToAlcoholID(id uint) bool

	IsValidAttitudeToSmokingID(id uint) bool

	IsValidStatusID(id uint) bool

	IsValidInterestIDs(ids []uint) bool

	IsValidInterest(id uint) bool
}
