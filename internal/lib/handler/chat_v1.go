package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

	desc "github.com/AlekseiAndriushin/go_chat_server/pkg/chat_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatRPCServerV1 struct {
	desc.UnimplementedChat_V1Server
	log *log.Logger
}

func NewChatRPCServerV1(log *log.Logger) *ChatRPCServerV1 {
	return &ChatRPCServerV1{
		log: log,
	}
}

func (s *ChatRPCServerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	buf := strings.Builder{}
	buf.WriteString("Received Create:\n")

	for i, user := range req.Usernames {
		buf.WriteString(fmt.Sprintf("\t#%d Username: %s\n", i, user))
	}

	s.log.Println(color.BlueString(buf.String()))

	if dline, ok := ctx.Deadline(); ok {
		s.log.Println(color.BlueString("Deadline: %v", dline))
	}

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<62))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.log.Println(color.GreenString(fmt.Sprintf("Response Create:\n\tID: %v\n", randInt64.Int64())))

	return &desc.CreateResponse{
		Id: randInt64.Int64(),
	}, nil
}

func (s *ChatRPCServerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	s.log.Println(color.BlueString("Received Delete:\n\tID: %v", req.GetId()))

	if dline, ok := ctx.Deadline(); ok {
		s.log.Println(color.BlueString("Deadline: %v", dline))
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatRPCServerV1) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	s.log.Println(color.BlueString("Received SendMessage:\n\tFrom: %v\n\tText: %v\n\tTimestamp: %v",
		req.GetFrom(),
		req.GetText(),
		req.GetTimestamp(),
	))

	if dline, ok := ctx.Deadline(); ok {
		s.log.Println(color.BlueString("Deadline: %v", dline))
	}

	return &emptypb.Empty{}, nil
}