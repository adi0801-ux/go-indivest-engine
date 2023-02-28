package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"strings"
	"time"
)

func (p *UserSrv) UserSignUp(signUp *models.UserSignup) (int, interface{}, error) {
	signUp.EmailId = strings.ToLower(signUp.EmailId)

	_, err := p.UserRepo.ReadUserByEmail(signUp.EmailId)
	if err == nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf("user email already exists")
	}
	if err != nil && err.Error() != constants.UserNotFound {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	user := &models.User{
		CreatedAt:     time.Time{},
		UserId:        utils.GenerateUserID(),
		FirstName:     signUp.FirstName,
		LastName:      signUp.LastName,
		EmailId:       signUp.EmailId,
		PhoneNumber:   signUp.PhoneNumber,
		Password:      utils.HashPassword(signUp.Password),
		DeviceToken:   signUp.DeviceToken,
		DeviceType:    signUp.DeviceType,
		ProfileStatus: "0",
	}

	err = p.UserRepo.CreateUser(user)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//create jwt
	token, err := p.CreateJwt(user)

	user.Password = ""
	return http.StatusOK, map[string]interface{}{"token": token, "user": user}, nil
}

func (p *UserSrv) UserLogin(login *models.UserLogin) (int, interface{}, error) {
	login.EmailId = strings.ToLower(login.EmailId)
	user, err := p.UserRepo.ReadUserByEmail(login.EmailId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	match := utils.CheckPasswordHash(login.Password, user.Password)
	if !match {
		return http.StatusBadRequest, nil, fmt.Errorf("password is incorrect")
	}

	//create jwt
	token, err := p.CreateJwt(user)

	user.Password = ""
	return http.StatusOK, map[string]interface{}{"token": token, "user": user}, nil

}

func (p *UserSrv) UserVerifyToken(token string) (int, interface{}, error) {

	// Initialize a new instance of `Claims`
	claims := &models.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.Config.JwtSecretKey), nil
	})
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, map[string]interface{}{"user_id": claims.UserId, "exp": claims.ExpiresAt}, nil

}

func (p *UserSrv) CreateJwt(user *models.User) (string, error) {
	expire, _ := time.ParseDuration(p.Config.TokenExpire)

	// Create the JWT claims, which includes the username and expiry time
	claims := &models.Claims{
		Email:  user.EmailId,
		UserId: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(p.Config.JwtSecretKey))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		utils.Log.Error(err)
		return tokenString, fmt.Errorf("error creating token")
	}
	return tokenString, nil

}

func (p *UserSrv) SaveOnBoardingQuestionnaire(questions *models.UserQuestioner) (int, interface{}, error) {
	questionDb := &models.UserLeads{
		UserId:            questions.UserId,
		InvestingInterest: questions.InvestingInterest,
		Profession:        questions.Profession,
	}

	err := p.UserRepo.UpdateOrCreateUserLeads(questionDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	user, err := p.UserRepo.ReadUser(questions.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	user.ProfileStatus = "1"

	err = p.UserRepo.UpdateOrCreateUser(user)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, nil, nil
}
