package main

import (
	"context"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/skipio-corp/gr/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:2323", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := service.NewContactServiceClient(conn)
	opts := []grpc.CallOption{}
	upsertClient, err := client.UpsertContacts(context.Background(), opts...)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10_000; i++ {
		contact := service.Contact{
			FirstName:   gofakeit.FirstName(),
			LastName:    gofakeit.LastName(),
			Email:       gofakeit.Email(),
			MobilePhone: gofakeit.PhoneFormatted(),
		}
		err = upsertClient.Send(&contact)
		if err != nil {
			panic(err)
		}
	}
}
