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

func (c *Cache) Get(key string) pb.User {
	valJSON, err := c.Client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var val pb.User
	json.Unmarshal([]byte(valJSON), &val)
	return val
}

func (c *Cache) Set(user pb.User) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	err = c.Client.Set(user.Name, userJSON, 10*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Cache) GetAll() []pb.User {
	iter := c.Client.Scan(0, "", 0).Iterator()

	var users []pb.User

	for iter.Next() {
		user := c.Get(iter.Val())
		users = append(users, user)
	}

	if err := iter.Err(); err != nil {
		fmt.Println(err)
	}
	return users

}
