package main

import (
	"fmt"
	"github.com/pleum/viabusgo/v1"
	"log"
)

func main() {
	client, err := viabusgo.New()
	if err != nil {
		log.Fatalln(err)
	}

	client.Anonymous()

	res, err := client.RegisterAnonymous()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
