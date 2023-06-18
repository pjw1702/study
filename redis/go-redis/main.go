package main

import (
	"fmt"

	"github.com/pjw1702/study/redis/go-redis/repositories"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	products, err := productRepo.GetProducts()
	if err != nil {
		fmt.Errorf("Fail to find record for infinitas DB: %w", err)
		return
	}
	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/infinitas")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
