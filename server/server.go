package main

import (
	"cache-service/service"
	"cache-service/z_generated/redisService"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	port := flag.Int("port",0,"the server port");

	flag.Parse();

	log.Printf("starting server on port %d",*port);

	server := service.NewRedisServiceServer();
	grpcServer := grpc.NewServer();

	redisService.RegisterRedisServiceServer(grpcServer,server);

	address := fmt.Sprintf("0.0.0.0:%d",*port);

	listener,err := net.Listen("tcp",address);
	
	if err!=nil{
		log.Fatal("cannot start the redis service server:",err);
	}

	err = grpcServer.Serve(listener);

	if err!=nil{
		log.Fatal("cannot start the redis service server:",err);
	}

}