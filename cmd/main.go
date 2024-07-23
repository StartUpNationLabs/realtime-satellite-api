package main

import (
	"context"
	"github.com/NYTimes/gziphandler"
	v1 "github.com/StartUpNationLabs/react-flight-tracker-satellite/gen/go/proto/satellite/v1"
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/service"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func startDebug() {
	log.Println("Starting pprof server on " + ":10001")
	err := http.ListenAndServe(":10001", nil)
	if err != nil {
		log.Fatal("failed to start pprof server: ", err)
	}

}

var isReady = false

func monitoring(host string) {
	log.Println("Starting monitoring server on " + ":10002")
	mux := http.NewServeMux()
	mux.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Error("Error writing response livez: ", err)
		}
	})
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if isReady {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("OK"))
			if err != nil {
				log.Error("Error writing response readyz: ", err)
			}
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, err := w.Write([]byte("Not ready"))
			if err != nil {
				log.Error("Error writing response readyz: ", err)
			}
		}

	})
	err := http.ListenAndServe(host+":10002", mux)

	if err != nil {
		log.Fatal("failed to start monitoring server: ", err)
	}
}

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Error loading .env file")
	}
	if os.Getenv("DEBUG") == "true" {
		go startDebug()
	}
	host := "localhost"
	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}
	// start monitoring server
	go monitoring(host)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start the gRPC server
	lis, err := net.Listen("tcp", host+":5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterSatelliteServiceServer(grpcServer, service.NewSatelliteService())
	log.Println("gRPC server ready on ", host+":5566")
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
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Error(err)

		}
	}(conn)

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

	withGzip, _ := gziphandler.GzipHandlerWithOpts(gziphandler.CompressionLevel(9))

	// mount the gRPC HTTP gateway to the root
	mux.Handle("/", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(withGzip(rmux)))

	// mount a path to expose the generated OpenAPI specification on disk
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving OpenAPI specification...")
		http.ServeFile(w, r, "./gen/v1/service.swagger.json")
	})

	// mount the Swagger UI that uses the OpenAPI specification path above
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))

	log.Println("HTTP server ready on " + host + ":8080")
	isReady = true
	err = http.ListenAndServe(host+":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(mux))
	if err != nil {
		log.Fatal(err)
	}
}
