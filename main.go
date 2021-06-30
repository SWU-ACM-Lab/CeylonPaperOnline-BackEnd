package main

import (
	"CeylonPaperOnline-BackEnd/Controller"
	"CeylonPaperOnline-BackEnd/Middleware"
	"fmt"
	_ "fmt"
)

func main() {
	var test Middleware.QueryConsole
	var config Middleware.DatabaseConfig
	Middleware.Console.SetStatus(true)
	config.LoadConfig("/Users/sunist/Projects/2021/CeylonPaperOnline-BackEnd/Bin/dbconfig.json")
	test.Connect(config)

	var Auth Controller.AuthGateway
	Auth.Init(test)
	fmt.Println(Auth.GetUserWithAuth("100000000000000", "114514"))
}
