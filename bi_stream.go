package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/RickDeb2004/go-grpc-server/proto"
)
func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Calling SayHelloBidirectionalStreaming")
	stream,err:=client.SayHelloBidirectionalStreaming(context.Background())
	if err!=nil{
		log.Fatalf("could not send names to server: %v",err)
	}
	waitc:=make(chan struct{})
//sending names one by one
	go func(){
		for{
			message,err:=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("could not receive message: %v",err)
			}	
			log.Println(message)
		}
		close(waitc)
	}()
	for _,name:=range names.Names{
		req:=&pb.HelloRequest{Name:name}
		if err:=stream.Send(req);err!=nil{
			log.Fatalf("could not send request: %v",err)
		}
		time.Sleep(2*time.Second)
		<-waitc //block until goroutine is done with
		log.Printf("Streaming finished")
	}
			
}