package services

import (
	"log"
	"net"

	"github.com/akshay0074700747/projectandCompany_management_protofiles/pb/authpb"
	"google.golang.org/grpc"
)

type AuthEngine struct {
	Srv authpb.AuthServiceServer
}

func NewAuthEngine(srv authpb.AuthServiceServer) *AuthEngine {
	return &AuthEngine{
		Srv: srv,
	}
}
func (engine *AuthEngine) Start(addr string) {

	server := grpc.NewServer()
	authpb.RegisterAuthServiceServer(server, engine.Srv)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	log.Printf("Auth Server is listening...")

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to listen on port %s: %v", addr, err)
	}

}
