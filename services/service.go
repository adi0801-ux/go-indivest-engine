package services

import (
	"indivest-engine/repositories"
)

type ServiceConfig struct {
	UserRep *repositories.UserDetailsRepository
}

type SandboxServiceConfig struct {
	SandboxRep *repositories.SandboxRepository
	RedisRep   *repositories.RedisRepository
}
