package main

import (
	"context"
	"fmt"

	pb "github.com/pieterverheyden/grpc-go-example/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *UserManagementServer) RemoveUserByEmail(ctx context.Context, in *pb.UserEmail) (*emptypb.Empty, error) {
	if !IsEmailValid(in.GetEmail()) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Bad request: invalid email %v", in.GetEmail()),
		)

	}

	deleteSql := `DELETE FROM users WHERE email = $1`
	res, err := server.conn.Exec(context.Background(), deleteSql, in.GetEmail())

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal server error: %v", err),
		)
	}

	if res.RowsAffected() != 1 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("User not found: %v", err),
		)
	}

	return &emptypb.Empty{}, nil
}
