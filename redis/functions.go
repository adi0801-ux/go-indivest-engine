package redis

import "context"

func (c *Client) GetKeyValue_(key string) (string, error) {
	result := c.store.Get(context.Background(), key)

	return result.Result()
}
