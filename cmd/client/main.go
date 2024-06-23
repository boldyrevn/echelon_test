package main

import (
    "context"
    "echelon_task/internal/model"
    grpcimpl "echelon_task/internal/ports/grpc"
    "echelon_task/pkg/proto"
    "flag"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "os"
)

var cliQualityToModel = map[string]model.ThumbnailQuality{
    "low":    model.LowQuality,
    "medium": model.MediumQuality,
    "high":   model.HighQuality,
}

var (
    address    = flag.String("address", "", "gRPC server address in format \"host:port\"")
    asyncFlag  = flag.Bool("async", false, "do requests asynchronously")
    cliQuality = flag.String(
        "quality",
        "low",
        "specifies thumbnail resolution, can be one of: low, medium, high",
    )
)

func main() {
    flag.Parse()
    if *address == "" {
        fmt.Println("address parameter must be specified")
        flag.PrintDefaults()
        return
    }

    quality, ok := cliQualityToModel[*cliQuality]
    if !ok {
        fmt.Println("incorrect quality")
        flag.PrintDefaults()
        return
    }

    grpcClient, err := grpc.NewClient("localhost:54320", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        fmt.Printf("error occurred during client creation: %e\n", err)
        return
    }

    client := proto.NewThumbnailServiceClient(grpcClient)

    dir, err := os.Getwd()
    if err != nil {
        fmt.Printf("can't get name of current directory: %e\n", err)
        return
    }

    urls := flag.Args()
    if *asyncFlag {
        if err := grpcimpl.DownloadFilesAsynchronously(context.Background(), client, quality, dir, urls...); err != nil {
            fmt.Printf("failed to download some files: %s", err)
            return
        }
    } else {
        if err := grpcimpl.DownloadFiles(context.Background(), client, quality, dir, urls...); err != nil {
            fmt.Printf("failed to download some files: %s", err)
            return
        }
    }
    fmt.Println("thumbnails were successfully downloaded")
}
