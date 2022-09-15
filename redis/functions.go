package redis

import (
	"context"
	"fmt"
	"indivest-engine/constants"
)

func (c *Client) GetKeyValue_(key string) (string, error) {
	result := c.store.Get(context.Background(), key)

	value, err := result.Result()
	if err != nil {
		return "", fmt.Errorf(constants.NoSuchSchemCodeExists)
	}
	return value, err
}
