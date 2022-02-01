package service

import (
	"cache-service/z_generated/redisService"
	"context"
	"errors"
)

type RedisServiceServer struct{
	redisService.UnimplementedRedisServiceServer
}
func NewRedisServiceServer() *RedisServiceServer{
	return &RedisServiceServer{};
}
func (server *RedisServiceServer)Set(ctx context.Context, req *redisService.SetRequest)(*redisService.SetResponse,error){
	return nil,errors.New("not implemented yet. <yourname> will implement me")

}

func (server *RedisServiceServer)Get(ctx context.Context, req *redisService.GetRequest)(*redisService.GetResponse,error){
	return nil,errors.New("not implemented yet. <yourname> will implement me")

}