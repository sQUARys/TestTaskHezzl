package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"time"
)

type Cache struct {
	Client *redis.Client
}

func New() *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return &Cache{
		Client: client,
	}
}

func (c *Cache) GetUser(key string) (pb.User, error) {
	valJSON, err := c.Client.Get(key).Result()
	if err != nil {
		return pb.User{}, err
	}
	var val pb.User
	json.Unmarshal([]byte(valJSON), &val)

	return val, nil
}

func (c *Cache) SetUser(user pb.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = c.Client.Set(user.Name, userJSON, time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetUsers() ([]*pb.User, error) {
	iter := c.Client.Scan(0, "", 0).Iterator()

	var users []*pb.User

	for iter.Next() {
		user, err := c.GetUser(iter.Val())
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func (c *Cache) DeleteUser(key string) {
	c.Client.Del(key)
}
