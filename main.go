package main

import (
	"github.com/sQUARys/TestTaskHezzl/cache"
	"github.com/sQUARys/TestTaskHezzl/repositories"
	"github.com/sQUARys/TestTaskHezzl/services"
)

//var (
//	buf    bytes.Buffer
//	logger = log.New(&buf, "INFO: ", log.Lshortfile)
//
//	infof = func(info string) {
//		logger.Output(2, info)
//	}
//)
//infof("Hello world")

//1. Описать proto файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
//2. Реализовать gRPC сервис на основе proto файла на Go
//3. Для хранения данных использовать PostgreSQL
//4. на запрос получения списка пользователей данные будут кешироваться в redis на минуту и брать из редиса
//5. При добавлении пользователя делать лог в clickHouse
//6. Добавление логов в clickHouse делать через очередь Kafka

func main() {
	db := repositories.New()
	c := cache.New()
	//kfk := kafka.New()

	service := services.New(db, c)
	service.Start()

}
