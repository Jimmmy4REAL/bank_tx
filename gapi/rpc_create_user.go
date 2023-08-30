package gapi

import (
	"context"
	"time"

	db "github.com/Jimmmy4REAL/bank_tx/db/sqlc"
	"github.com/Jimmmy4REAL/bank_tx/worker"
	"github.com/hibiken/asynq"

	"github.com/Jimmmy4REAL/bank_tx/pb"
	"github.com/Jimmmy4REAL/bank_tx/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	// using internal method change to readable not-binary format
	arg := db.CreateUserParams{

		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}
	// retrieve create_user_and_email_sending tx_result
	user, err := server.store.CreateUser(ctx, arg) // async function -
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	// create db tx to avoid duplicate work
	taskPayload := &worker.PayloadSendVerifyEmail{
		Username: user.Username,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to distribute tasks to send email")
	}
	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
