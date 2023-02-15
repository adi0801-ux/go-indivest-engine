package models

import "time"

type FundHousesSupported struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	AMCID     int       `gorm:"column:amc_id;not null" json:"amc_id"`
	AMCCode   string    `gorm:"column:amc_code;not null" json:"amc_code"`
	Name      string    `gorm:"column:name;not null" json:"same"`
	Logo      string    `gorm:"column:logo;not null" json:"logo"`
	Active    int       `gorm:"column:active;not null" json:"active"`
}

type FundsSupported struct {
	ID                         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt                  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                  time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	SavvyCode                  string    `gorm:"column:savvy_code" json:"savvy_code"`
	AMFICode                   string    `gorm:"column:amfi_code" json:"amfi_code"`
	Name                       string    `gorm:"column:name" json:"name"`
	Category                   string    `gorm:"column:category" json:"category"`
	Active                     int       `gorm:"column:active" json:"active"`
	MinimumFirstTimeInvestment string    `gorm:"column:minimum_first_time_investment" json:"minimum_first_time_investment"`
	MinimumOngoingInvestment   string    `gorm:"column:minimum_ongoing_investment" json:"minimum_ongoing_investment"`
	MinimumRedemptionAmount    string    `gorm:"column:minimum_redemption_amount" json:"minimum_redemption_amount"`
	SettlementDays             string    `gorm:"column:settlement_days" json:"settlement_days"`
	MinimumSipAmount           string    `gorm:"column:minimum_sip_amount" json:"minimum_sip_amount"`
	MinimumSwpAmount           string    `gorm:"column:minimum_swp_amount" json:"minimum_swp_amount"`
	MinimumStpAmount           string    `gorm:"column:minimum_stp_amount" json:"minimum_stp_amount"`
	CagrY1                     float64   `gorm:"column:cagr_y1" json:"cagr_y1"`
	CagrY3                     float64   `gorm:"column:cagr_y2" json:"cagr_y2"`
	CagrY5                     float64   `gorm:"column:cagr_y3" json:"cagr_y3"`
	AMCID                      int       `gorm:"column:amc_id" json:"amc_id"`
	AMCCode                    string    `gorm:"column:amc_code" json:"amc_code"`
	NAV                        int       `gorm:"column:nav" json:"nav"`
}

type FundHousesList struct {
	Amcs []struct {
		Name     string `json:"name"`
		Code     string `json:"code"`
		Id       int    `json:"id"`
		Branding struct {
			Id        int         `json:"id"`
			FundId    interface{} `json:"fund_id"`
			PartnerId interface{} `json:"partner_id"`
			Logo      string      `json:"logo"`
			OccStyles struct {
				HeaderBackgroundColor     string `json:"headerBackgroundColor"`
				HeaderTitleColor          string `json:"headerTitleColor"`
				HeaderTitleFontSize       string `json:"headerTitleFontSize"`
				MobileHeaderTitleFontSize string `json:"mobileHeaderTitleFontSize"`
				ButtonColor               string `json:"buttonColor"`
				ButtonBackgoundColor      string `json:"buttonBackgoundColor"`
				InputLabelColor           string `json:"inputLabelColor"`
				FontsColor                string `json:"fontsColor"`
				CloseButtonColor          string `json:"closeButtonColor"`
				FontFamily                string `json:"font-family"`
				TableHeaderTextColor      string `json:"tableHeaderTextColor"`
				TableBackgroundColor      string `json:"tableBackgroundColor"`
			} `json:"occ_styles"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			AmcId     int       `json:"amc_id"`
		} `json:"branding"`
	} `json:"amcs"`
}

type FundDetails struct {
	Funds []struct {
		Name                       string `json:"name"`
		Active                     bool   `json:"active"`
		Code                       string `json:"code"`
		AmfiCode                   string `json:"amfi_code"`
		MinimumFirstTimeInvestment string `json:"minimum_first_time_investment"`
		MinimumOngoingInvestment   string `json:"minimum_ongoing_investment"`
		MinimumRedemptionAmount    string `json:"minimum_redemption_amount"`
		SettlementDays             int    `json:"settlement_days"`
		MinimumSipAmount           string `json:"minimum_sip_amount"`
		MinimumSwpAmount           string `json:"minimum_swp_amount"`
		MinimumStpAmount           string `json:"minimum_stp_amount"`
		FactsheetLink              string `json:"factsheet_link"`
		Category                   string `json:"category"`
		AmcId                      int    `json:"amc_id"`
		FundInfo                   struct {
			Nav         int     `json:"nav"`
			ReturnYear1 float64 `json:"return_year_1"`
			ReturnYear3 float64 `json:"return_year_3"`
			ReturnYear5 float64 `json:"return_year_5"`
		} `json:"fund_info"`
		RiskRating   *int        `json:"risk_rating"`
		ExpenseRatio *string     `json:"expense_ratio"`
		FundManagers interface{} `json:"fund_managers"`
	} `json:"funds"`
}
