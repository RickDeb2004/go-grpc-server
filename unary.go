package main
import(
	"context"
	"log"
	"time"
	pb "github.com/RickDeb2004/go-grpc-server/proto"

)
func callSayHello(client pb.GreetServiceClient){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	res,err:=client.SayHello(ctx,&pb.NoParam{})
	if err!=nil{
		log.Fatalf("could not greet: %v",err)
	}
	log.Printf("Greeting: %s",res.Message)
}