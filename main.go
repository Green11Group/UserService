package main

import (
	"GreenThumbUser/genproto"
	"GreenThumbUser/serviceuser"
	"GreenThumbUser/storage/postgres"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	postgres.Config()

	db,err:=postgres.Connection()
	if err !=nil{
		log.Fatal("Postgresni connection qilishda xatolik?",err)
	}

	jwtSecretKey:=os.Getenv("JWT_SECRET_KEY")
	user:=postgres.NewUserRepo(db)

	userserver:=serviceuser.NewUserServer(db,jwtSecretKey,user)

	listener,err:=net.Listen("tcp",":50051")
	if err!=nil{
		log.Fatal("Tinglashda xatolik?",err)
	}

	server:=grpc.NewServer()

	genproto.RegisterUserServer(server,userserver)

	if err:=server.Serve(listener); err!=nil{
		log.Fatal("Serverni ishga tushirishda xatolik?",err)
	}
}