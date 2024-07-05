package main

import (
	"GreenThumb/User-Service/genproto"
	"GreenThumb/User-Service/serviceuser"
	"GreenThumb/User-Service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	postgres.Config()

	db, err := postgres.Connection()
	if err != nil {
		log.Fatal("Postgres connection failed?", err)
	}

	user := serviceuser.NewUserServer(db, postgres.NewUserRepo(db))

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Listening failed", err)
	}

	server := grpc.NewServer()

	genproto.RegisterUserServer(server, user)

	if err := server.Serve(listener); err != nil {
		log.Fatal("Server run failed?", err)
	}
}
