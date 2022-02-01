package main

import (
	"cache-service/z_generated/redisService"
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
)

func main(){
	serverAddress := flag.String("address","","the Server Address");
	flag.Parse();

	log.Printf("dial server %s",*serverAddress);

	conn,err := grpc.Dial(*serverAddress,grpc.WithInsecure());

	if err!=nil{
		log.Fatal("could not dial Server:", err);
	}

	client := redisService.NewRedisServiceClient(conn);

	req := &redisService.SetRequest{
		Key: "Mayank",
		Value:[]byte("hello"),
	};

	res,err := client.Set(context.Background(),req);
	
	if err!=nil{
		log.Fatal("cannot set:",err);
	}

	log.Println(res);

}