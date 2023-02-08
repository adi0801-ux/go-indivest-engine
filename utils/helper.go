package utils

import (
	"github.com/segmentio/ksuid"
	"indivest-engine/constants"
	"math"
	"time"
)

func GenerateID() string {
	//"github.com/segmentio/ksuid"
	return ksuid.New().String()
}

func GenerateSipID() string {
	return "SIP_" + GenerateID()
}

func GenerateTransactionID() string {
	return "TRANS_" + GenerateID()
}

//	func GeneratePartnerTransactionID() string {
//		return "TRANS_" + GenerateID()
//	}
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

func GeneratePartnerTransactionID() string {
	return "PART_TRANS_" + GenerateID()
}
