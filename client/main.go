package main

import (
  "log"

  "../api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

func main() {
  var conn *grpc.ClientConn

  conn, err := grpc.Dial(":5678", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  c := api.NewPingClient(conn)

  response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "hello, it's me"})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Greeting)
}
