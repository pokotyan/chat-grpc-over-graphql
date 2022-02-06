package main

import (
	"google.golang.org/grpc/reflection"
	"grpc-server/service"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"grpc-server/handler"
	"grpc-server/pb"

	"google.golang.org/grpc"
)

func register(server *grpc.Server) {
	var ms service.MathService
	ms = service.NewMathService()
	mathServer := handler.NewMathHandler(ms)
	pb.RegisterMathServiceServer(server, mathServer)

	var cs service.ChatService
	cs = service.NewChatService()
	chatServer := handler.NewChatHandler(cs)
	pb.RegisterChatServiceServer(server, chatServer)

	reflection.Register(server)
}

func main() {
	grpcServer := grpc.NewServer()
	register(grpcServer)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		grpcServer.GracefulStop()
		// grpc.Stop() // leads to error while receiving stream response: rpc error: code = Unavailable desc = transport is closing
		wg.Done()
	}()

	log.Println("starting grpc server")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
	wg.Wait()
	log.Println("clean shutdown")
}
