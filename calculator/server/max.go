package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

// 関数が終わることでserver側のsendStreamを閉じる
// clientからrequestとchannelを受け取っている
// 受け取ったchannelを操作する関数が詰め込まれたstructが返ってくる
func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max functionwas invoked")
	var maximum int32 = 0

	for {
		// goroutine(関数内で受信する)
		req, err := stream.Recv()

		// client側から渡ってきたcontextを利用している
		// serverからのstreamが途絶えてしまったらcontextがerrorをはいてRecv()を中断さえる
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while readin client stream: %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}