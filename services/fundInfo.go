package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
	"strconv"
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
