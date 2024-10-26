package main

import (
	"fmt"
	"livecode/livecode/internal/pkg/storage"
	"log"
)

func main() {
	s, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	s.Set("egor", "zxc")
	fmt.Println(*s.Get("egor"))
}
