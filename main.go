package main

import (
	"github.com/vongphachan/funny-store-backend/src/initiators"
)

func main() {
	db := initiators.StartMongoDB()

	initiators.StartGin(db)
}
