package gapi

import (
	db "github.com/Jimmmy4REAL/bank_tx/db/sqlc"
	"github.com/Jimmmy4REAL/bank_tx/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// change db to pb format

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
