package main // import "github.com/wanghaoxi3000/micro-shippy"

import (
	"context"
	"errors"
	"log"
	"os"

	pb "./proto/consignment"
	// pb "github.com/wanghaoxi3000/micro-shippy/proto/consignment"

	vesselPb "./proto/vessel"
	// vesselPb "github.com/wanghaoxi3000/micro-shippy/proto/vessel"

	userPb "./proto/user"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		// consignment-service 独立测试时不进行认证
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]

		// Auth here
		authClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

func main() {
	repo, err := CreateStore()
	defer repo.Close()
	if err != nil {
		log.Fatalf("create store error: %v\n", err)
	}

	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	server.Init()
	log.Printf("micro start: go.micro.srv.consignment\n")

	vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{repo, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
