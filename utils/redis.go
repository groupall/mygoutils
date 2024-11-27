package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(server, port string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", server, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func RedisRemoveValue(rdb *redis.Client, key string) (string) {
	// expiration, err := strconv.Atoi(expire)
	// if err != nil {
	// 	log.Print("No se pudo convertir el expiration del redis")
	// 	return "", err
	// }
	// ctx := context.Background()

	return rdb.Del(context.Background(),key).String()
	
}
func RedisSetValue(rdb *redis.Client, key string, value interface{}, expire int) (string) {
	p, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return rdb.Set(context.Background(), key, p, time.Duration(expire)*time.Second).String()


}

func RedisGetValue(rdb *redis.Client, key string) (string,error) {
	log.Printf("redis- leyendo %s", key)
	rdo, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("error al obtener la key %v", err)
		return "", err
	}
	return rdo, nil

}
