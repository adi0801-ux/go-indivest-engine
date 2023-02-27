package services

import (
	"indivest-engine/repositories"
	"indivest-engine/utils"
)

type RiskCalculatorService struct {
	UserRepo *repositories.UserRepository
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
	Config   *utils.Config
	UserRepo *repositories.UserRepository
}
