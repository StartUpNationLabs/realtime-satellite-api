package main

import (
	"context"
	v1 "github.com/StartUpNationLabs/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Error loading .env file")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start the gRPC server
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterSatelliteServiceServer(grpcServer, service.NewSatelliteService())
	log.Println("gRPC server ready on localhost:5566...")
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// dial the gRPC server above to make a client connection
	conn, err := grpc.NewClient("localhost:5566", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	client := v1.NewSatelliteServiceClient(conn)
	err = v1.RegisterSatelliteServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	// create a standard HTTP router
	mux := http.NewServeMux()

	// mount the gRPC HTTP gateway to the root
	mux.Handle("/", rmux)

	// mount a path to expose the generated OpenAPI specification on disk
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving OpenAPI specification...")
		http.ServeFile(w, r, "./gen/v1/service.swagger.json")
	})

	// mount the Swagger UI that uses the OpenAPI specification path above
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))

	log.Println("HTTP server ready on localhost:8080...")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
