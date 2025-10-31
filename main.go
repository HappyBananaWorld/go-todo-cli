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
	"--store":storeTodo,
	"--get":getTodo,
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

func get(key string) (string,error){
	val, err := rdb.Get(ctx,key).Result()

	if err != nil{
		return "",err
	}

	return val,nil
}

func showHelp(args []string){
    fmt.Println("  --version Show version")
}

func storeTodo(args []string){
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


func getTodo(args[]string){
	key := args[0]

	val,err := get(key)

	if err != nil{
		color.Red("Something wen wrong :(")
		return
	}

	color.Green(val)
}