package recommendations

import (
	"dating_service/internal/model"
	"sort"
)

type Service struct {
	userProvider   UserProvider
	filterProvider FilterProvider
}

func NewRecommendationService(userProvider UserProvider, filterProvider FilterProvider) *Service {
	return &Service{userProvider, filterProvider}
}

type ScoredUser struct {
	User  model.User
	Score float64
}

type ScoringWeights struct {
	HasBioBonus   float64
	SameCityBonus float64

	DatingGoalBaseWeight      float64
	WorldviewBaseWeight       float64
	EducationBaseWeight       float64
	AttitudeBaseWeight        float64
	SharedInterestsBaseWeight float64
	CandidateIsActive         float64
}

func DefaultWeights() ScoringWeights {
	return ScoringWeights{
		HasBioBonus:   30,
		SameCityBonus: 100,

		DatingGoalBaseWeight:      150,
		WorldviewBaseWeight:       30,
		EducationBaseWeight:       25,
		AttitudeBaseWeight:        80,
		SharedInterestsBaseWeight: 50,
		CandidateIsActive:         300,
	}
}

func calculateInterestsScore(userInterests, candidateInterests []*model.Interest, baseWeight float64) float64 {
	if len(userInterests) == 0 || len(candidateInterests) == 0 {
		return 0
	}

	userSet := make(map[uint]struct{})
	for _, interest := range userInterests {
		userSet[interest.ID] = struct{}{}
	}

	intersectionSize := 0
	unionSet := make(map[uint]struct{}, len(userSet))
	for k, v := range userSet {
		unionSet[k] = v
	}

	for _, interest := range candidateInterests {
		if _, found := userSet[interest.ID]; found {
			intersectionSize++
		}
		unionSet[interest.ID] = struct{}{}
	}

	if len(unionSet) == 0 {
		return 0
	}

	jaccardIndex := float64(intersectionSize) / float64(len(unionSet))

	return baseWeight * jaccardIndex
}

func CalculateMatchScore(currentUser, candidateUser *model.User, weights ScoringWeights) float64 {
	var score float64

	if candidateUser.SexID == 3 || candidateUser.SexID == 2 {
		score -= weights.CandidateIsActive
	}

	if candidateUser.Bio != nil && *candidateUser.Bio != "" {
		score += weights.HasBioBonus
	}

	if currentUser.City != nil && candidateUser.City != nil && *currentUser.City == *candidateUser.City {
		score += weights.SameCityBonus
	}

	datingMatrix := DatingGoalCompatibilityMatrix()
	worldviewMatrix := WorldviewCompatibilityMatrix()
	educationMatrix := EducationCompatibilityMatrix()
	attitudeMatrix := AttitudeCompatibilityMatrix()

	score += calculateCompatibilityScore(currentUser.TypeOfDatingID, candidateUser.TypeOfDatingID, weights.DatingGoalBaseWeight, datingMatrix)
	score += calculateCompatibilityScore(currentUser.WorldviewID, candidateUser.WorldviewID, weights.WorldviewBaseWeight, worldviewMatrix)
	score += calculateCompatibilityScore(currentUser.EducationID, candidateUser.EducationID, weights.EducationBaseWeight, educationMatrix)
	score += calculateCompatibilityScore(currentUser.AttitudeToSmokingID, candidateUser.AttitudeToSmokingID, weights.AttitudeBaseWeight, attitudeMatrix)
	score += calculateCompatibilityScore(currentUser.AttitudeToAlcoholID, candidateUser.AttitudeToAlcoholID, weights.AttitudeBaseWeight, attitudeMatrix)
	score += calculateInterestsScore(currentUser.Interests, candidateUser.Interests, weights.SharedInterestsBaseWeight)
	if score < 0 {
		return 0
	}
	return score
}

func (s *Service) GetRecommendations(currentUserID uint, page, pageSize int) ([]ScoredUser, error) {
	currentUser, err := s.userProvider.FindUserWithoutEntity(currentUserID)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, ErrUserNotFound
	}
	userFilter, err := s.filterProvider.GetFilter(currentUserID)
	if err != nil {
		return nil, err
	}
	if userFilter == nil {
		return nil, ErrFilterNotFound
	}
	users, _, err := s.userProvider.FindUsersWithFilter(userFilter, page, pageSize)
	if err != nil {
		return nil, err
	}
	filteredUsers := users
	scoredUsers := make([]ScoredUser, 0, len(filteredUsers))

	baseScoreWeight := DefaultWeights()
	for _, candidate := range filteredUsers {
		if candidate.ID == currentUserID {
			continue
		}
		score := CalculateMatchScore(currentUser, candidate, baseScoreWeight)
		scoredUsers = append(scoredUsers, ScoredUser{
			User:  *candidate,
			Score: score,
		})
	}
	sort.Slice(scoredUsers, func(i, j int) bool {
		return scoredUsers[i].Score > scoredUsers[j].Score
	})
	return scoredUsers, nil
}
