package main

import (
	"log"
	"test_sat/route"

	"github.com/joho/godotenv"
)

var err error

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
	r := route.SetupRouter()

	//running
	r.Run()
}
