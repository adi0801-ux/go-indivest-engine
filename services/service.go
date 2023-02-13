package services

import (
	"indivest-engine/repositories"
	"indivest-engine/utils"
)

type RiskCalculatorService struct {
	UserRep *repositories.UserDetailsRepository
}

type SandboxServiceConfig struct {
	SandboxRep *repositories.SandboxRepository
	RedisRep   *repositories.RedisRepository
}

type MFService struct {
	TSAClient *repositories.TSAClient
	SavvyRepo *repositories.SavvyRepository
	//ShowAccountRepo *repositories.AccountRepository
	config *utils.Config
}
