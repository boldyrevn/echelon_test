package main

import (
    grpcimpl "echelon_task/internal/ports/grpc"
    "echelon_task/pkg/app"
    "echelon_task/pkg/proto"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "google.golang.org/grpc"
)

func init() {
    _ = os.Setenv("REDIS_HOST", "localhost")
    _ = os.Setenv("REDIS_PORT", "6379")
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.ServicePort))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    exit := make(chan os.Signal)
    signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

    serverImpl := grpcimpl.NewServer()
    err = serverImpl.Init()
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    proto.RegisterThumbnailServiceServer(s, serverImpl)

    go func() {
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()
    log.Printf("server listening at %v", lis.Addr())

    <-exit
    s.GracefulStop()
    if err := serverImpl.Shutdown(); err != nil {
        log.Printf("failed to shutdown server: %e", err)
    }
    log.Printf("server stopped")
}
