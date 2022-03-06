package main

import (
	"context"
	"fmt"
	pb "github.com/PudgeKim/go-holdem-protos/protos"
	"google.golang.org/grpc"
	"time"
)

const address = "localhost:6060"

func main() {
	clientMain()
}

func clientMain() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("client did not connect: ", err.Error())
	}
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := client.GetUser(ctx, &pb.UserId{Id: "107848443863300905065"})
	if err != nil {
		fmt.Println("client_main_user: ", user)
	}
	fmt.Println("client_main_user: ", user)
}
