package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// insertだけはinsertする上で情報に_idが存在しなくてはならないため、bson化、decodeを内部ですることでもし_idがなかった場合には_idを追加している
func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", in)

	data := BlogItem {
		AuthorId: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}

	// 内部でbsonに変換している
	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	// IDを持ったinterface→関数を使えるようにprimitive.ObjectIDでキャストしている
	// どのようとでもIDを使えるように関数が用意されている
	// interfaceのdecodeは特別な作業が必要だから、decodeも済ませてある(関数が含まれているとencode, decodeが大変だからinterfaceを渡してある)
	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}