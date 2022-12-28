package services

import (
	"indivest-engine/repositories"
)

type ServiceConfig struct {
	UserRep                    *repositories.UserDetailsRepository
	TSAClient                  *repositories.TSAClient
	FullKycRepo                *repositories.AddFullKyc
	ReadPanCardRepo            *repositories.ReadPanCardRepository
	ReadAddressProofRepo       *repositories.ReadAddressProofReposotry
	StartVideoVerificationRepo *repositories.StartVideoVerificationRepository
}

type SandboxServiceConfig struct {
	SandboxRep *repositories.SandboxRepository
	RedisRep   *repositories.RedisRepository
}
