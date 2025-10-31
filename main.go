package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/redis/go-redis/v9"
)

var commands = map[string]func([]string){
	"--help":showHelp,
	"--store":store,
}

var (
	ctx = context.Background()

	// connect to redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
)

func main() {
	
	defer rdb.Close()

	err:=rdb.Set(ctx,"name","sa3232jad11",0).Err()

	if err != nil{
		fmt.Println("Error On storing something in redis", err)
	}

	args := os.Args
	command := args[1]

	if fn, ok := commands[command]; ok {
		fn(args[2:])
	}
}

func set(key string, value string, ttl time.Duration) bool{
	err := rdb.Set(ctx,key,value,ttl).Err()

	if err != nil{
		return false
	}

	return true
}

func showHelp(args []string){
    fmt.Println("  --version Show version")
}

func store(args []string){
	if len(args) > 2{
		color.Red("only gives, key/value")
	}

	key := args[0]
	value := args[1]
	result := set(key,value,0)	

	if !result{
		color.Red("Something wen wrong :(")
		return
	}
	
	color.Green("Stored successfully")
}
