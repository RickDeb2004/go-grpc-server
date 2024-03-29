package main

import ("log"
		
		"context"
		pb "github.com/RickDeb2004/go-grpc-server/proto"
		"time"
	)

	func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList){
		log.Printf("Calling SayHelloClientStreaming")
		stream,err:=client.SayHelloClientStreaming(context.Background())
		if err!=nil{
			log.Fatalf("could not send names to server: %v",err)
		}
		for _,name:=range names.Names{
			req:=&pb.HelloRequest{Name:name}
			if err:=stream.Send(req);err!=nil{
				log.Fatalf("could not send request: %v",err)
			}
			log.Printf("Sent: %s",name)
			time.Sleep(2*time.Second)
		}
		res,err:=stream.CloseAndRecv()
		log.Printf("Client streaming finished")
		if err!=nil{
			log.Fatalf("could not receive response: %v",err)
		}
		log.Printf("Received: %v",res.Messages)

	
	}