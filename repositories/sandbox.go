package repositories

import (
	"indivest-engine/constants"
	"indivest-engine/db"
	"indivest-engine/models"
	"indivest-engine/utils"
)

type SandboxRepository struct {
	Db *db.Database
}

func (s *SandboxRepository) CreateUserWallet(userID string) error {
	wallet := models.UserWallet{
		UserID:    userID,
		CreatedAt: utils.GetCurrentDateTime(),
		INR:       constants.DefaultSandboxWalletHoldings,
	}

	err := s.Db.CreateUserWallet_(&wallet)
	return err
}

func (s *SandboxRepository) CreateGetUserWallet(userID string) (*models.UserWallet, error) {
	wallet := models.UserWallet{
		UserID:    userID,
		CreatedAt: utils.GetCurrentDateTime(),
		INR:       constants.DefaultSandboxWalletHoldings,
	}

	err := s.Db.CreateUserWallet_(&wallet)
	return &wallet, err
}

func (s *SandboxRepository) ReadUserWallet(userId string) (*models.UserWallet, error) {

	return s.Db.ReadUserWallet_(userId)
}

func (s *SandboxRepository) UpdateUserWallet(u *models.UserWallet) error {

	return s.Db.UpdateUserWallet_(u)
}

//Transaction

func (s *SandboxRepository) CreateMFTransaction(w *models.UserMFTransactions) error {
	return s.Db.CreateMFTransaction_(w)
}

func (s *SandboxRepository) ReadMFTransaction(userId string, schemeCode string) (*[]models.UserMFTransactions, error) {
	return s.Db.ReadMFTransactions_(userId, schemeCode)
}

func (s *SandboxRepository) ReadAllMFTransactions(userId string) (*[]models.UserMFTransactions, error) {
	return s.Db.ReadAllMFTransactions_(userId)
}

// CreateUserSIP create user SIP
func (s *SandboxRepository) CreateUserSIP(w *models.UserMFActiveSIP) error {
	return s.Db.CreateUserSIP_(w)
}

// UpdateOrCreateUserHoldings create or update user holdings
func (s *SandboxRepository) UpdateOrCreateUserHoldings(w *models.UserMFHoldings) error {
	return s.Db.UpdateOrCreateUserHoldings_(w)
}

func (s *SandboxRepository) ReadUserHolding(userId string, schemeCode string) (*models.UserMFHoldings, error) {
	return s.Db.ReadUserHolding_(userId, schemeCode)
}

func (s *SandboxRepository) ReadUserHoldings(userId string) (*[]models.UserMFHoldings, error) {
	return s.Db.ReadUserHoldings_(userId)
}
