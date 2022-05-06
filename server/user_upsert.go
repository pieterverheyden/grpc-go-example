package main

import (
	"context"
	"fmt"

	pb "github.com/pieterverheyden/grpc-go-example/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *UserManagementServer) CreateOrUpdateUser(ctx context.Context, in *pb.User) (*emptypb.Empty, error) {
	if !IsEmailValid(in.GetEmail()) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Bad request: invalid email %v", in.GetEmail()),
		)

	}

	upsertSql := `INSERT INTO users (email, first_name, last_name, gender) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO UPDATE SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, gender = EXCLUDED.gender`
	_, err := server.conn.Exec(context.Background(), upsertSql, in.GetEmail(), in.GetFirstName(), in.GetLastName(), in.Gender.String())

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal server error: %v", err),
		)
	}

	return &emptypb.Empty{}, nil
}
