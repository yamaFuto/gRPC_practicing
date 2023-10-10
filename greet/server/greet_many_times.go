package main

import (
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

// 関数が終わることでserver側のsendStreamを閉じる
// clientからchannelを受け取っている
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
