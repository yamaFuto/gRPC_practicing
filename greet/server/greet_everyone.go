package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

// 関数を終わらせることでserver側のsendStreamを閉じる
// clientから二つのchannelを受け取っている
func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		// clientからgoroutineを受信している
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		// clientにgoroutineで送信している
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client; %v\n", err)
		}
	}
}