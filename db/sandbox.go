package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateUserWallet_(w *models.UserWallet) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserWallet_(userId string) (*models.UserWallet, error) {
	u := &models.UserWallet{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateUserWallet_(u *models.UserWallet) error {

	result := d.store.Where("user_id = ?", u.UserID).Updates(u)

	return result.Error
}

//Transaction

func (d *Database) CreateMFTransaction_(w *models.UserMFTransactions) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadMFTransactions_(userId string, schemeCode string) (*[]models.UserMFTransactions, error) {
	u := &[]models.UserMFTransactions{}
	err := d.store.Where("user_id = ? and scheme_code = ?", userId, schemeCode).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoTransactionFound)
	}
	return u, nil
}

func (d *Database) ReadAllMFTransactions_(userId string) (*[]models.UserMFTransactions, error) {
	u := &[]models.UserMFTransactions{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoTransactionFound)
	}
	return u, nil
}

//User SIP's

func (d *Database) CreateUserSIP_(w *models.UserMFActiveSIP) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) UpdateOrCreateUserHoldings_(w *models.UserMFHoldings) error {
	result := d.store.Model(&w).Where("user_id = ? AND scheme_code = ?", w.UserID, w.SchemeCode).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) ReadUserHolding_(userId string, schemeCode string) (*models.UserMFHoldings, error) {
	u := &models.UserMFHoldings{}
	err := d.store.Where("user_id = ? and scheme_code = ?", userId, schemeCode).Find(u).Error
	if err != nil {
		return u, err
	}
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.NoHoldingsFound)
	}
	return u, nil
}

func (d *Database) ReadUserHoldings_(userId string) (*[]models.UserMFHoldings, error) {
	u := &[]models.UserMFHoldings{}
	err := d.store.Where("user_id = ? ", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {
		return u, fmt.Errorf(constants.NoHoldingsFound)
	}
	return u, nil
}
