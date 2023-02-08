package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type AuthVerificationResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Exp    int    `json:"exp"`
		UserId string `json:"user_id"`
	} `json:"data"`
	Error interface{} `json:"error"`
}

func (s *HTTPServer) WebhookAuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(c.GetReqHeaders())
		apiKey := c.GetReqHeaders()["X-Api-Key"]

		if apiKey == "" {

			errorResponse(c, http.StatusUnauthorized, fmt.Errorf("no x-api-key in header provided"))

			return nil
		}
		if apiKey == s.config.XApiKey {
			err := c.Next()
			if err != nil {
				utils.Log.Warn(err)
				return nil
			}
		} else {
			errorResponse(c, http.StatusUnauthorized, fmt.Errorf("incorrect x-api-key in header provided"))
			return nil
		}

		return nil
	}
}
func (s *HTTPServer) AuthorizeMiddleware(AuthApi string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.GetReqHeaders()["Authorization"]

		if clientToken == "" {

			errorResponse(c, http.StatusBadRequest, fmt.Errorf("no Authorization header provided"))

			return nil
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])

		} else {

			errorResponse(c, http.StatusBadRequest, fmt.Errorf("incorrect format of authorization token"))

			return nil
		}

		AuthResp, err := CheckTokenAuth(clientToken, AuthApi)

		if err != nil {
			// oops!! something went wrong

			// log
			utils.Log.Error(err)
			errorResponse(c, http.StatusBadRequest, err)
			return nil
		}
		if AuthResp.Status != 1 {
			// unauthorized

			errorResponse(c, http.StatusBadRequest, fmt.Errorf("bearer token unauthorized"))
			return nil
		} else {
			c.Locals("userId", AuthResp.Data.UserId)
			err := c.Next()
			if err != nil {
				utils.Log.Warn(err)
				return nil
			}

		}

		err = c.Next()
		if err != nil {
			utils.Log.Debug(err)
			return nil
		}

		return nil
	}
}

func CheckTokenAuth(token string, AuthApi string) (AuthVerificationResp, error) {
	var resp AuthVerificationResp

	url := AuthApi
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		utils.Log.Error(err)

		return resp, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		utils.Log.Error(err)
		return resp, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		utils.Log.Error(err)

		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		utils.Log.Error(err)
		return AuthVerificationResp{}, err
	}

	return resp, nil
}
