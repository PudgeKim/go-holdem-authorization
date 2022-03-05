package main

import (
	"fmt"
	"github.com/Pudgekim/db"
	"github.com/Pudgekim/handlers"
	"github.com/Pudgekim/infrastructure/persistence"
	pb "github.com/Pudgekim/protos"
	"github.com/Pudgekim/protoserver"
	"google.golang.org/grpc"
	"net"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(50) PRIMARY KEY,
	name VARCHAR(50) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	balance BIGSERIAL NOT NULL
);
`

const RestServerPort = ":3000"
const ProtoServerPort = ":6060"

func main() {
	conn, err := db.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("postgres connection fail: %s", err.Error()))
	}

	err = conn.Ping()
	if err != nil {
		panic(fmt.Sprintf("postgres ping error: %s", err.Error()))
	}

	fmt.Println("ping success!")

	conn.MustExec(schema)

	userRepo := persistence.NewUserRepository(conn)

	protoAuthServer := protoserver.NewAuthServer(userRepo)
	go RunProtoServer(protoAuthServer)

	handler := handlers.NewHandler(userRepo)
	router := handler.Routes()

	router.Run(RestServerPort)

}

func RunProtoServer(protoAuthServer *protoserver.Auth) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost%s", ProtoServerPort))
	if err != nil {
		fmt.Println("proto server failed to listen: ", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, protoAuthServer)
	grpcServer.Serve(lis)
}
