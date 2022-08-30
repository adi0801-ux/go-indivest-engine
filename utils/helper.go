package utils

import (
	"github.com/segmentio/ksuid"
)

func GenerateID() string {
	//"github.com/segmentio/ksuid"
	return ksuid.New().String()
}
