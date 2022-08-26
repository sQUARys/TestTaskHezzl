package main

import (
	"github.com/sQUARys/TestTaskHezzl/cache"
	"github.com/sQUARys/TestTaskHezzl/controllers"
	"github.com/sQUARys/TestTaskHezzl/kafka"
	"github.com/sQUARys/TestTaskHezzl/repositories"
	"github.com/sQUARys/TestTaskHezzl/services"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	db := repositories.New()
	c := cache.New()
	kfk := kafka.New()

	wg.Add(1)
	go kfk.Start(wg)

	service := services.New(db, c, kfk)

	wg.Add(1)
	go service.Start(wg)

	controller := controllers.New(service)
	controller.Start()

	wg.Wait()
}
