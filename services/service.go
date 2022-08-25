package services

import (
	"fmt"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"log"
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
	repo ordersRepository
	c    cache
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
		repo: repo,
		c:    cache,
	}
}

func (serv *Service) Start() {

	serv.AddUser(newUser1)
	serv.AddUser(newUser2)

	user, err := serv.c.GetUser(newUser1.Name)
	if err != nil {
		log.Println("Error in service : ", err)
	}

	fmt.Println(fmt.Sprintf("GetUser method. ID : %d , Name : %s", user.Id, user.Name))

	users, err := serv.c.GetUsers()
	if err != nil {
		log.Println("Error in service : ", err)
	}
	fmt.Println("GetUsers method before deleting: ", users)

	serv.DeleteUser(newUser1.Name)

	users, err = serv.c.GetUsers()
	if err != nil {
		log.Println("Error in service : ", err)
	}
	fmt.Println("GetUsers method after deleting: ", users)

}

func (serv *Service) AddUser(user pb.User) {
	serv.repo.AddUser(&user)
	serv.c.AddUser(user)
}

func (serv *Service) DeleteUser(name string) {
	serv.repo.DeleteUser(name)
	serv.c.DeleteUser(name)
}
