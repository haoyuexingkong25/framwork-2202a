package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GetGrpc(port int64, server func(s *grpc.Server)) error {
	//创建一个tcp链接 打印端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	//初始化一个新的grpc服务
	s := grpc.NewServer()
	//调用闭包函数
	server(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
