package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

// 関数が終わることでserver側のsendStreamを閉じる
// clientからrequestとchannelを受け取っている
// 受け取ったchannelを操作する関数が詰め込まれたstructが返ってくる
func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("primes function was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {

			// goroutine(関数内で送信)
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	// sendStreamを閉じる(send(response)を終える→serviceを終了する(serverの役割はrequestをもとにresponseを返すこと))
	return nil
}

///handlerとしてすべて返したら、sendstreamを閉じる仕組みを持っている(server側に備わっている)