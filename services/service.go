package services

import (
	"indivest-engine/repositories"
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
}
