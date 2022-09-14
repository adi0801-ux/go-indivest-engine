package repositories

import "indivest-engine/redis"

type RedisRepository struct {
	Db *redis.Client
}

func (r *RedisRepository) GetKeyValue(key string) (string, error) {

	return r.Db.GetKeyValue_(key)
}
