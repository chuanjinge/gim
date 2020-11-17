package main

import (
	"context"
	"fmt"
	"gim/pkg/pb"
	"google.golang.org/grpc"
)

/**
从服务端获取设备ID
 */
func main() {
	connect, err := grpc.Dial("192.168.33.20:50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建客户端调用对象
	client := pb.NewLogicExtClient(connect)

	registerDeviceReq := pb.RegisterDeviceReq{
		Type:2,
		Brand:"小米3",
		Model:"青春版3",
		SystemVersion:"14.53",
		SdkVersion:"2.2.13",
	}

	response, err := client.RegisterDevice(context.Background(), &registerDeviceReq)

	fmt.Println(response)

}



