package services

import (
	"indivest-engine/repositories"
)

type ServiceConfig struct {
	UserRep  *repositories.UserDetailsRepository
}