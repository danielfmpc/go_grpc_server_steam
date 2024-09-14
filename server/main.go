package main

import (
	"fmt"
	"log"
	"net"

	"github.com/danielfmpc/client_go_rpc_server_stream/server/src/pb/department"
	"google.golang.org/grpc"
)

type Server struct {
	department.DepartmentServiceServer
}

func (s *Server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersonServer) error {
	return nil
}

func main() {
	fmt.Println("start")

	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("erro listner", err)
	}

	s := grpc.NewServer()
	department.RegisterDepartmentServiceServer(s, &Server{})
	if err := s.Serve(listner); err != nil {
		log.Fatal("error on server ", err)
	}
}
