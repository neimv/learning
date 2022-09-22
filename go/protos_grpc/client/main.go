package main

import (
	"context"
	"grpc-n/testpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)

	// DoUnary(c)
	// DoClientStreaming(c)
	// DoServerStreaming(c)
	DoBidireccionalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "1",
	}

	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GetTes: %v", err)
	}

	log.Printf("response from server: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q1t1",
			Answer:   "azul",
			Question: "choose color",
			TestId:   "1",
		},
		{
			Id:       "q2t1",
			Answer:   "red",
			Question: "flags",
			TestId:   "1",
		},
		{
			Id:       "q3t1",
			Answer:   "pokemon",
			Question: "rare animals",
			TestId:   "1",
		},
	}

	stream, err := c.SetQuestions(context.Background())

	if err != nil {
		log.Fatalf("error with calling SetQuestions: %v", err)
	}

	for _, question := range questions {
		log.Println("sending question: ", question.Id)
		stream.Send((question))
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error to close stream: %v", err)
	}

	log.Printf("response from server: %v", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GetStudentsPerTest: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("response from server: %v", msg)
	}
}

func DoBidireccionalStreaming(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{
		Answer: "42",
	}

	numberOfQuestions := 6

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())

	if err != nil {
		log.Fatalf("error while calling TakeTest: %v", err)
	}

	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reading stream: %v", err)
				break
			}

			log.Printf("response from server: %v", res)
		}

		close(waitChannel)
	}()

	<-waitChannel
}
