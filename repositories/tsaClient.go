package repositories

import (
	"bytes"
	"encoding/json"
	"indivest-engine/models"
	"indivest-engine/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type TSAClient struct {
	Client  *http.Client
	BaseUrl string
	LogRep  *ApiLogsRepository
	Token   *string
}

func CreateHttpClient() *http.Client {
	client := &http.Client{Timeout: 40 * time.Second}
	return client
}
func (h *TSAClient) SendGetRequest(endpoint string, params url.Values) (response *http.Response, errResp error) {
	method := http.MethodGet
	URL := h.BaseUrl + endpoint

	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.Encode()

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Method:    method,
		Endpoint:  endpoint,
		Params:    params.Encode(),
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	response, errResp = h.callTSA(req, apiLog.RequestId)

	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))

	//updates in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)
	}
	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp

}
func (h *TSAClient) SendPostRequest(endpoint string, body interface{}) (response *http.Response, errResp error) {
	method := http.MethodPost
	URL := h.BaseUrl + endpoint
	//marshal - converts go objects to json
	jsonBody, err := json.Marshal(body)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Method:    method,
		Params:    req.URL.RawQuery,
		Payload:   string(jsonBody),
		Endpoint:  endpoint,
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	response, errResp = h.callTSA(req, apiLog.RequestId)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))

	//updates in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)
	}
	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp
}

func (h *TSAClient) SendPutRequest(endpoint string, body interface{}) (response *http.Response, errResp error) {
	method := http.MethodPut
	URL := h.BaseUrl + endpoint
	//marshal - converts go objects to json
	jsonBody, err := json.Marshal(body)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Params:    req.URL.RawQuery,
		Payload:   string(jsonBody),
		Method:    method,
		Endpoint:  endpoint,
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	response, errResp = h.callTSA(req, apiLog.RequestId)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))

	//updates in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)
	}
	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp
}

func (h *TSAClient) SendPostRequestWithoutResponseLog(endpoint string, body interface{}) (response *http.Response, errResp error) {
	method := http.MethodPost
	URL := h.BaseUrl + endpoint
	jsonBody, err := json.Marshal(body)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	//save to db
	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Params:    req.URL.RawQuery,
		Payload:   string(jsonBody),
		Method:    method,
		Endpoint:  endpoint,
	}
	err = h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		return nil, err
	}

	response, errResp = h.callTSA(req, apiLog.RequestId)

	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp
}

func (h *TSAClient) callTSA(req *http.Request, RequestId string) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Request-ID", RequestId)
	req.Header.Add("Authorization", "Bearer "+*h.Token)
	res, err := h.Client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return res, err
	}
	return res, nil
}

func (h *TSAClient) callPostFormTSA(req *http.Request, RequestId string, header string) (*http.Response, error) {
	req.Header.Add("Content-Type", header)
	req.Header.Add("Request-Id", RequestId)
	req.Header.Add("Authorization", "Bearer "+*h.Token)
	res, err := h.Client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return res, err
	}
	return res, err
}

func (h *TSAClient) SendPostFormRequest(endpoint string, body *bytes.Buffer, header string) (response *http.Response, errResp error) {
	URL := h.BaseUrl + endpoint
	method := http.MethodPost

	apiLog := &models.APILog{
		RequestId: utils.GenerateID(),
		CreatedAt: time.Time{},
		Method:    method,
		Endpoint:  endpoint,
	}
	err := h.LogRep.CreateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return nil, err
	}
	req, err := http.NewRequest(method, URL, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}
	response, errResp = h.callPostFormTSA(req, apiLog.RequestId, header)
	respBytes, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(respBytes))
	//	update in db
	apiLog.Response = string(respBytes)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		apiLog.Error = string(respBytes)
	}

	apiLog.ResponseCode = response.StatusCode
	apiLog.ResponseDate = time.Now().String()
	err = h.LogRep.UpdateApiLog(apiLog)
	if err != nil {
		utils.Log.Error(err)
		return response, err
	}
	return response, errResp

}
