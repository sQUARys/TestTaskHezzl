package services

import (
	"bytes"
	"fmt"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"golang.org/x/net/context"
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

	bufError bytes.Buffer
	bufUser  bytes.Buffer

	loggerError = log.New(&bufError, "Error: ", log.Lshortfile)
	loggerUser  = log.New(&bufUser, "User : ", log.Lshortfile)

	loggingError = func(body string) {
		loggerError.Output(2, body)
	}
	loggingUser = func(body string) {
		loggerUser.Output(3, body)
	}

	ctx = context.Background()
)

type Service struct {
	repo  ordersRepository
	Cache cache
	Kafka kafka
}

type kafka interface {
	Start(wg sync.WaitGroup)
	ReadText(ctx context.Context) error
	WriteText(log string, ctx context.Context) error
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

func New(repo ordersRepository, cache cache, kfk kafka) *Service {
	return &Service{
		repo:  repo,
		Cache: cache,
		Kafka: kfk,
	}
}

// This func was added to demonstate working all methods with Postgres and Redis
func (serv *Service) Start(wg sync.WaitGroup) {
	defer wg.Done()

	serv.AddUser(newUser1)
	serv.AddUser(newUser2)

	loggingUser(fmt.Sprintf("Id : %d  , Name : %s", newUser1.Id, newUser1.Name))
	serv.Kafka.WriteText("Add "+bufUser.String(), ctx)

	user, err := serv.Cache.GetUser(newUser1.Name)
	if err != nil {
		loggingError(err.Error())
		serv.Kafka.WriteText(bufError.String(), ctx)
	}

	fmt.Println(fmt.Sprintf("GetUser method. ID : %d , Name : %s", user.Id, user.Name))

	users, err := serv.Cache.GetUsers()
	if err != nil {
		loggingError(err.Error())
		serv.Kafka.WriteText(bufError.String(), ctx)
	}
	fmt.Println("GetUsers method before deleting: ", users)

	serv.DeleteUser(newUser1.Name)

	users, err = serv.Cache.GetUsers()
	if err != nil {
		loggingError(err.Error())
		serv.Kafka.WriteText(bufError.String(), ctx)
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
