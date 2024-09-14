package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/danielfmpc/client_go_rpc_server_stream/client/src/pb/department"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("erro on connect server ", err)
	}

	defer conn.Close()

	client := department.NewDepartmentServiceClient(conn)
	stream, err := client.ListPerson(context.Background(), &department.ListPersonRequest{
		DepartmentId: 5,
	})

	if err != nil {
		log.Fatal("erro on get channel to steam ", err)
	}

	for {
		response, err := stream.Recv()
		if response == nil {
			fmt.Println("Departament empty")
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("erro on recv", err)
		}

		fmt.Println("response: ", response)
	}
}
