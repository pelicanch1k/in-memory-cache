package main

import (
	"log"
	"time"

	"github.com/pelicanch1k/in-memory-cache/pkg/cache"
)

func main() {
	memory := cache.New()
	memory.Set("test", "gg", time.Second*2)

	value, _ := memory.Get("test")
	println(value.(string))

	time.Sleep(time.Second)

	value, _ = memory.Get("test")
	println(value.(string))

	time.Sleep(time.Second)

	_, err := memory.Get("test")
	if err {
		log.Fatal(err)
	}

}
