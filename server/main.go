package main

import (
  "fmt"
  "log"
  "net"

  "../api"
  "google.golang.org/grpc"
)

func main(){
  lis ,err := net.Listen("tcp", fmt.Sprintf(":%d", 5678))
  if err != nil {
    log.Fatalf("failed to listen :%v",err)
  }

  s:= api.Server{}
  grpcServer := grpc.NewServer()

  api.RegisterPingServer(grpcServer, &s)

  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }

}
