package utils

import (
	"bytes"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"indivest-engine/constants"
	"math"
	"time"
)

func GenerateID() string {
	//"github.com/segmentio/ksuid"
	return ksuid.New().String()
}
func GenerateUserID() string {
	//"github.com/segmentio/ksuid"
	return ksuid.New().String()
}

func GenerateSipID() string {
	return "SIP_" + GenerateID()
}

func GenerateTransactionID() string {
	return "TRANS_" + GenerateID()
}

func GenerateWithdrawalId() string {
	return "WITH_" + GenerateID()
}

func GetCurrentDate() int {
	currentTime := time.Now()
	return currentTime.Day()
}

func GetCurrentDateTime() time.Time {
	return time.Now()
}

func RoundOfTo2Decimal(f float64) float64 {
	if f < constants.DefaultfloatPrecissionAccepted {
		return 0
	}

	return math.Round(f*100) / 100
}
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
func GeneratePartnerTransactionID() string {
	return "PART_TRANS_" + GenerateID()
}

func Transcode(in, out interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) string {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(pwd)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
