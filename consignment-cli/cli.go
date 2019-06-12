package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	pb "./proto/consignment"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

const (
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.srv.consignment.cli"))
	service.Init()

	// 在命令行中指定用户 token
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments, expecing token.")
	}
	token := os.Args[1]

	// 初始化 gRPC 客户端
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())

	// 解析货物信息
	infoFile := DEFAULT_INFO_FILE
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 创建带有用户 token 的 context
	// consignment-service 服务端将从中取出 token，解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// 调用 RPC
	// 将货物存储到指定用户的仓库里
	log.Printf("created: %v", consignment)
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	log.Printf("created: %t", resp.Created)

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}

	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}
}
