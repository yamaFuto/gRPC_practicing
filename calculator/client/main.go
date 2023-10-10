package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// 生成したコードでラッピングすることでシリアライズ、ディスシリアライズなどをする
	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, 10)
	doSqrt(c, -2)
}


// CPUを圧迫するためstreamの通信路を閉じる
// 常に送る側が閉じなくてはならない(送られる側は常に受け身であり送られてこなかったら次の操作に移行する)
// server側は呼ばれた関数が終わると同時にstreamを破棄するのに対して、client側はpb.CalculatorServiceClient(pointer)で常にstreamを保持してしまうため、コードで明示的にstream通信を閉じる必要がある