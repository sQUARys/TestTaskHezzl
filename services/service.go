package services

import (
	"fmt"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"log"
	"sync"
)

var (
	newUser1 = pb.User{
		Id:   1,
		Name: "John",
	}
	newUser2 = pb.User{
		Id:   2,
		Name: "Dmitrii",
	}
)

type Service struct {
	repo  ordersRepository
	Cache cache
}

type cache interface {
	GetUser(key string) (pb.User, error)
	AddUser(user pb.User) error
	GetUsers() ([]*pb.User, error)
	DeleteUser(key string)
}

type ordersRepository interface {
	AddUser(user *pb.User) error
	DeleteUser(key string) error
}

func New(repo ordersRepository, cache cache) *Service {
	return &Service{
		repo:  repo,
		Cache: cache,
	}
}

// This func was added to demonstate working all methods with Postgres and Redis
func (serv *Service) Start(wg sync.WaitGroup) {
	defer wg.Done()

	serv.AddUser(newUser1)
	serv.AddUser(newUser2)

	user, err := serv.Cache.GetUser(newUser1.Name)
	if err != nil {
		log.Println("Error in service : ", err)
	}

	fmt.Println(fmt.Sprintf("GetUser method. ID : %d , Name : %s", user.Id, user.Name))

	users, err := serv.Cache.GetUsers()
	if err != nil {
		log.Println("Error in service : ", err)
	}
	fmt.Println("GetUsers method before deleting: ", users)

	serv.DeleteUser(newUser1.Name)

	users, err = serv.Cache.GetUsers()
	if err != nil {
		log.Println("Error in service : ", err)
	}
	fmt.Println("GetUsers method after deleting: ", users)

}

func (serv *Service) AddUser(user pb.User) {
	serv.repo.AddUser(&user)
	serv.Cache.AddUser(user)
}

func (serv *Service) DeleteUser(name string) {
	serv.repo.DeleteUser(name)
	serv.Cache.DeleteUser(name)
}
