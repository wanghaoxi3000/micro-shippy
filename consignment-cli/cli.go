package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	pb "./proto/consignment"
	"github.com/micro/go-micro"
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

	// 初始化 gRPC 客户端
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())

	// 解析货物信息
	infoFile := DEFAULT_INFO_FILE
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 调用 RPC
	// 将货物存储到我们自己的仓库里
	log.Printf("created: %v", consignment)
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	// 新货物是否托运成功
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