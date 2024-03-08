package main

import (
	"github.com/vongphachan/funny-store-backend/src/initiators"
	initiatesmongodb "github.com/vongphachan/funny-store-backend/src/initiators/mongodb"
)

func main() {
	db := initiatesmongodb.Start()

	initiators.StartGin(db)
}
