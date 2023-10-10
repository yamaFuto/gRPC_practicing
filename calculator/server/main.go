package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	//structの中のパッケージ(interface)→空のメソッドが作られる(nil)
	// structの中に入っているstructやinterfaceのメンバーはその親のstructの中で広げられた状態で呼ぶことができる(継承)
	// protoがserverのinterfaceを作るときにmethodが存在するとそれに合わせてinterfaceに新たなメソッドが挿入される→それを補うためにここでメソッドを宣言している(あくまでinterfaceとの整合性を合わせるためなので空のメソッドでもよい)
	// unimplemented→その時点で実装されていないmethod(メソッドの中身がnil)が入っている→後で使う(のちにoverrideされる)
	// method→付け忘れ防止、どうせ(のちにoverrideされる)＋nilで初期化しておくことでunimplementedでのちに使用することができる + error検出しやすい
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	//http2通信をするgrpc専用に作られたserver(シリアライズ、ディスシリアライズをprotoファイルに基づいて行う)
	s := grpc.NewServer()
	// implement server with server instance
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

// CPUを圧迫するためstreamの通信路を閉じる
// 常に送る側が閉じなくてはならない(送られる側は常に受け身であり送られてこなかったら次の操作に移行する)