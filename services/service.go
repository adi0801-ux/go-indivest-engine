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
	config    *utils.Config
}

type UserSrv struct {
	config   *utils.Config
	UserRepo *repositories.UserRepository
}
