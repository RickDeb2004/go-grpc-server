package main

import ("log"
		"io"
		"context"
		pb "github.com/RickDeb2004/go-grpc-server/proto"
	)

func callSayHelloServerStream(client pb.GreetServiceClient,names *pb.NameList){
	log.Printf("Calling SayHelloServerStreaming")
	stream,err:=client.SayHelloServerStreaming(context.Background(),names)
	if err!=nil{
		log.Fatalf("could not send names  to server: %v",err)
	}
	for{
		message,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("could not receive message: %v",err)
		}
		log.Printf("Received: %s",message.Message)
	}
	log.Printf("Finished streaming")
}