package main

import (
	"../web"
	"../persistence"
)

func main() {
	if err := persistence.InitDB(); err != nil{
		panic(err)
	}
	web.StartServer()
}

