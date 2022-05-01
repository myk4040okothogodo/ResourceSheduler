package main

import (
    "net"
    "os"
    "github.com/hashicorp/go-hclog"
    "google.golang.org/grpc"
    "github.com/myk4040okothogodo/ResourceSheduler/server"
     protos "github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos"
    "google.golang.org/grpc/reflection"
)


func main() {
    log := hclog.Default()
    gs := grpc.NewServer()
    cs := server.NewAssetAvailability(log)

    protos.RegisterAssetCheckServer(gs, cs)
    reflection.Register(gs)

    //create client instance
    //
    conn, err := grpc.Dial("localhost: 9092", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }

    defer conn.Close()

    protos.NewAssetCheckClient(conn)

    l, err := net.Listen("tcp", ":9092")
    if err != nil {
        log.Error("Unable to listen", "error", err)
        os.Exit(1)
    }

    gs.Serve(l)


}
