package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

// 関数が終わることでserver側のsendStreamを閉じる
// clientからrequestとchannelを受け取っている
// 受け取ったchannelを操作する関数が詰め込まれたstructが返ってくる
func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		// goroutine(関数内で受信する)
		req, err := stream.Recv()

		// client側から渡ってきたcontextを利用している
		// clientからのstreamが途絶えてしまったらcontextがerrorをはいてRecv()を中断さえる
		if err == io.EOF {
			// 名前は違うがほぼsendと同じ
			// responseの返却もgoroutine通信を利用している
			// goroutine通信(内部で送信している)
			return stream.SendAndClose((&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			}))
		}

		if err != nil {
			log.Fatalf("Error while reding client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)
		sum += req.Number
		count++
	}
}