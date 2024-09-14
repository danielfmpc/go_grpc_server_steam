package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/danielfmpc/client_go_rpc_server_stream/server/src/pb/department"
	"google.golang.org/grpc"
)

type Server struct {
	department.DepartmentServiceServer
}

func (s *Server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersonServer) error {
	file, err := os.Open("./data.csv")
	if err != nil {
		return fmt.Errorf("error on open", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ";")
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.Atoi(data[3])
		departmentId, _ := strconv.Atoi(data[4])

		if int32(departmentId) == req.GetDepartmentId() {
			if err := srv.Send(&department.ListPersonResponse{
				Id:           int32(id),
				Name:         name,
				Email:        email,
				Income:       int32(income),
				DepartmentId: int32(departmentId),
			}); err != nil {
				return fmt.Errorf("error on send %v", err)
			}
		}
	}
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
