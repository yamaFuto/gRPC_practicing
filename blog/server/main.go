package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	//structの中のパッケージ(interface)→空のメソッドが作られる(nil)
	// structの中に入っているstructやinterfaceのメンバーはその親のstructの中で広げられた状態で呼ぶことができる(継承)
	// protoがserverのinterfaceを作るときにmethodが存在するとそれに合わせてinterfaceに新たなメソッドが挿入される→それを補うためにここでメソッドを宣言している(あくまでinterfaceとの整合性を合わせるためなので空のメソッドでもよい)
	// unimplemented→その時点で実装されていないmethod(メソッドの中身がnil)が入っている→後で使う(のちにoverrideされる)
	// method→付け忘れ防止、どうせ(のちにoverrideされる)＋nilで初期化しておくことでunimplementedでのちに使用することができる + error検出しやすい
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	//http2通信をするgrpc専用に作られたserver(シリアライズ、ディスシリアライズをprotoファイルに基づいて行う)
	s := grpc.NewServer()
	// implement server with server instance
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}