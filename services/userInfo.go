package services

import (
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (p *UserSrv) SaveOnBoardingQuestionnaire(questions *models.UserQuestioner) (int, interface{}, error) {
	questionDb := &models.UserLeads{
		UserId:            questions.UserId,
		InvestingInterest: questions.InvestingInterest,
		Profession:        questions.Profession,
	}

	err := p.UserRepo.CreateUserLeads(questionDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, nil, nil
}
