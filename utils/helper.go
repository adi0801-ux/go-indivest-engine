package utils

import (
	"github.com/segmentio/ksuid"
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

func GetCurrentDate() int {

	currentTime := time.Now()

	return currentTime.Day()

}

func GetCurrentDateTime() time.Time {

	return time.Now()

}
