package main

import (
	"log"

	"github.com/GOLANG-NINJA/grpc-app/proto/notification"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{Message: "Я на пути к становлению ниндзей!"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("STATUS:", response.Status)
}
