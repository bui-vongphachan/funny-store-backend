package main

import (
	"github.com/vongphachan/funny-store-backend/src/initiators"
)

func main() {
	db := initiators.StartMongoDB()

	initiators.StartRedis()

	initiators.StartGin(db)
}
