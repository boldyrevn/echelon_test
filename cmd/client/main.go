package main

import (
    "context"
    "echelon_task/internal/model"
    grpcimpl "echelon_task/internal/ports/grpc"
    "flag"
)

var (
    asyncFlag = flag.Bool("async", false, "do requests asynchronously")
    quality   = flag.String("quality", model.LowQuality.String(), "specifies thumbnail resolution")
)

func writeToFile()

func downloadAsynchronously()

func main() {
    flag.Parse()
    if *asyncFlag {

    } else {
        err := grpcimpl.DownloadFiles(context.Background())
    }
}
