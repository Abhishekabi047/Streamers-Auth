package main

import (
	"log"
	"service/pkg/config"
	"service/pkg/di"
)

func main() {
	c,err:=config.LoadConfig()
	if err != nil{
		log.Fatal("failed at loading config",err.Error())
	}
	server,err1:=di.InitializeServer(c)
	if err1 != nil{
		log.Fatal("failed to init server",err1.Error())
	}
	if err :=server.Start();err!= nil{
		log.Fatalf("coudnt start server")
	}
}