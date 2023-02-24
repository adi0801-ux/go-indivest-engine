package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"strconv"
)

func (d *Database) CreateUserWallet_(w *models.UserWallet) error {
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserWallet_(userId string) (*models.UserWallet, error) {
	u := &models.UserWallet{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	//if err != nil {
	//	return u, err
	//}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllUserWallets_() (*[]models.UserWallet, error) {
	u := &[]models.UserWallet{}
	err := d.Store.Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoTransactionFound)
	}
	return u, nil
}

func (d *Database) UpdateUserWallet_(u *models.UserWallet) error {

	result := d.Store.Where("user_id = ?", u.UserID).Updates(u)

	return result.Error
}

//Transaction

func (d *Database) CreateMFTransaction_(w *models.UserMFTransactions) error {
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) ReadMFTransactions_(userId string, schemeCode string) (*[]models.UserMFTransactions, error) {
	u := &[]models.UserMFTransactions{}
	err := d.Store.Where("user_id = ? and scheme_code = ?", userId, schemeCode).Find(u).Error
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
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
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
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) FindAllUserSIPWithDate_(SipDate int) (*[]models.UserMFActiveSIP, error) {
	u := &[]models.UserMFActiveSIP{}

	result := d.Store.Where("sip_date = ? and active = ?", SipDate, constants.DefaultSIPActiveSatus).Find(u)
	return u, result.Error
}

func (d *Database) UpdateOrCreateUserHoldings_(w *models.UserMFHoldings) error {
	result := d.Store.Model(&w).Where("user_id = ? AND scheme_code = ?", w.UserID, w.SchemeCode).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) ReadUserHolding_(userId string, schemeCode string) (*models.UserMFHoldings, error) {
	u := &models.UserMFHoldings{}
	err := d.Store.Where("user_id = ? and scheme_code = ?", userId, schemeCode).Find(u).Error
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
	err := d.Store.Where("user_id = ? ", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {
		return u, fmt.Errorf(constants.NoHoldingsFound)
	}
	return u, nil
}

//daily report

func (d *Database) CreateMFDailyReport_(w *models.UserMFDailyReport) error {
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) ReadAllMFDailyReport_(userId string, daysLimit int) (*[]models.UserMFDailyReport, error) {
	u := &[]models.UserMFDailyReport{}
	queryParam := strconv.Itoa(daysLimit)
	err := d.Store.Where("user_id = ? and created_at >= (now() - INTERVAL '"+queryParam+" days')", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if len(*u) == 0 {

		return u, fmt.Errorf(constants.NoTransactionFound)
	}
	return u, nil
}
