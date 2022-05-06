package main

import (
	"context"
	"log"
	"net"
	"os"
	"regexp"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	pb "github.com/pieterverheyden/grpc-go-example/proto"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	conn *pgx.Conn
	pb.UnimplementedUserManagementServer
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("Server listening at %v\n", lis.Addr())

	return s.Serve(lis)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	var userManagementServer *UserManagementServer = NewUserManagementServer()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	userManagementServer.conn = conn

	if err := userManagementServer.Run(); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
