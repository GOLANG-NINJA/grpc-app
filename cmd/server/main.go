package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/GOLANG-NINJA/grpc-app/proto/notification"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Notify(ctx context.Context, n *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	fmt.Println("RECEIVED NOTIFICATION:", n.Message)
	return &notification.NotificationResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
