package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Update fund houses list
func (p *MFService) UpdateFundHouses() error {
	param := url.Values{}
	response, err := p.TSAClient.SendGetRequest(constants.ListAMCEndpoint, param)
	var data models.FundHousesList
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return err
	}

	//	create or update query
	listData := data.Amcs

	for _, fundHouse := range listData {
		fundHousesDb := &models.FundHousesSupported{
			CreatedAt: time.Time{},
			AMCID:     fundHouse.Id,
			AMCCode:   fundHouse.Code,
			Name:      fundHouse.Name,
			Logo:      fundHouse.Branding.Logo,
			Active:    1,
		}

		err := p.SavvyRepo.CreateOrUpdateFundHousesList(fundHousesDb)
		if err != nil {
			utils.Log.Error(err)
			return err
		}
	}
	return nil
}

// get list of fund houses available
func (p *MFService) GetListOfFundHouses() (int, interface{}, error) {

	fundHouseList, err := p.SavvyRepo.ReadAllFundHousesList()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, fundHouseList, nil
}

// Update funds
func (p *MFService) UpdateFunds() error {

	fundHouseList, err := p.SavvyRepo.ReadAllFundHousesList()
	if err != nil {
		utils.Log.Error(err)
		return err
	}

	for _, fundHouse := range *fundHouseList {

		fundHouseDetail := fundHouse

		func() {
			param := url.Values{}
			param.Add("amc_code", fundHouseDetail.AMCCode)
			response, err := p.TSAClient.SendGetRequest(constants.FundDetailsEndpoint, param)
			var data models.FundDetails
			//converting struct to []bytes
			err = json.NewDecoder(response.Body).Decode(&data)
			if err != nil {
				utils.Log.Error(err)
			}
			p.CreateOrUpdateFundHouse(data, fundHouseDetail.AMCCode, fundHouseDetail.AMCID)
		}()

	}

	return nil

}

func (p *MFService) CreateOrUpdateFundHouse(fundDetails models.FundDetails, amcCode string, amcId int) {
	for _, fundDetail := range fundDetails.Funds {
		fund := fundDetail
		func() {

			fundSupported := &models.FundsSupported{
				CreatedAt:                  time.Time{},
				SavvyCode:                  fund.Code,
				AMFICode:                   fund.AmfiCode,
				Name:                       fund.Name,
				Category:                   fund.Category,
				Active:                     1,
				MinimumFirstTimeInvestment: fund.MinimumFirstTimeInvestment,
				MinimumOngoingInvestment:   fund.MinimumOngoingInvestment,
				MinimumRedemptionAmount:    fund.MinimumRedemptionAmount,
				SettlementDays:             strconv.Itoa(fund.SettlementDays),
				MinimumSipAmount:           fund.MinimumSipAmount,
				MinimumSwpAmount:           fund.MinimumSwpAmount,
				MinimumStpAmount:           fund.MinimumStpAmount,
				CagrY1:                     fund.FundInfo.ReturnYear1,
				CagrY3:                     fund.FundInfo.ReturnYear3,
				CagrY5:                     fund.FundInfo.ReturnYear5,
				AMCID:                      amcId,
				AMCCode:                    amcCode,
				NAV:                        fundDetail.FundInfo.Nav,
			}
			err := p.SavvyRepo.CreateOrUpdateFundDetails(fundSupported)
			if err != nil {
				utils.Log.Error(err)
			}

		}()
	}
}

// get list of funds avaialble
func (p *MFService) GetListOfFunds() (int, interface{}, error) {

	fundHouseList, err := p.SavvyRepo.ReadAllFundDetails()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, fundHouseList, nil
}

func (p *MFService) GetFundDetail(AMFICode string) (int, interface{}, error) {

	fundHouseList, err := p.SavvyRepo.ReadFundDetails(AMFICode)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, fundHouseList, nil
}

func (p *MFService) ReturnsInterestCalculator(fundDtls *models.ReturnsCalc) (int, interface{}, error) {
	returnsDtls, err := p.SavvyRepo.ReadFundDetails(fundDtls.FundCode)
	if err != nil {
		utils.Log.Info(err)
		return http.StatusBadRequest, nil, err
	}
	var cagr float64
	if fundDtls.Tenure == 1 {
		cagr = returnsDtls.CagrY1
	} else if fundDtls.Tenure == 3 {
		cagr = returnsDtls.CagrY3
	} else if fundDtls.Tenure == 5 {
		cagr = returnsDtls.CagrY5
	}
	interest := (fundDtls.Amount * cagr * fundDtls.Tenure) / 100
	return http.StatusOK, interest, nil
}

func (p *MFService) Recommendations() (int, interface{}, error) {
	_, funds, err := p.GetListOfFunds()
	if err != nil {
		utils.Log.Info(err)
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, funds, err
}
func (p *MFService) PopularFunds() (int, interface{}, error) {
	_, funds, err := p.GetListOfFunds()
	if err != nil {
		utils.Log.Info(err)
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, funds, err
}
func (p *MFService) DistinctFunds() (int, interface{}, error) {
	distinctFund, err := p.SavvyRepo.ReadAllFundDetails()
	if err != nil {
		utils.Log.Info(err)
	}
	var str []string
	for _, fund := range *distinctFund {
		str = append(str, fund.Category)
	}
	return http.StatusOK, str, nil
}
func (p *MFService) FundCategories() (int, interface{}, error) {
	//uniqueFundCategory, err := p.SavvyRepo.ReadFundCategory()
	//if err != nil {
	//	utils.Log.Info(err)
	//	return http.StatusBadRequest, nil, err
	//}
	allFunds, err := p.SavvyRepo.ReadAllFundDetails()

	data := map[string][]models.FundsSupported{}
	//for _, category := range *uniqueFundCategory {
	//
	//	data[strings.ToLower(category.Category)] = []models.FundsSupported{}
	//}

	for _, fund := range *allFunds {
		data[strings.ToLower(fund.Category)] = append(data[strings.ToLower(fund.Category)], fund)
	}
	return http.StatusOK, data, err
}
