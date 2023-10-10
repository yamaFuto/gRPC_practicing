package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// 生成したコードでラッピングすることでシリアライズ、ディスシリアライズなどをする
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id) // valid
	// readBlog(c, "aNonExistingId") // invalid
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}