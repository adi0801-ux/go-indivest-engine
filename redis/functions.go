package redis

import (
	"context"
	"fmt"
	"indivest-engine/constants"
	"time"
)

func (c *Client) GetKeyValue_(key string) (string, error) {
	result := c.store.Get(context.Background(), key)
	//c.store.HGetAll()

	value, err := result.Result()
	if err != nil {
		return "", fmt.Errorf(constants.NoSuchSchemCodeExists)
	}
	return value, err
}

func (c *Client) SetKeyValue_(key string, value interface{}) error {
	result := c.store.Set(context.Background(), key, value, 6*time.Hour)

	err := result.Err()
	return err
}
