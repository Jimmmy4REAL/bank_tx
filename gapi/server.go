package gapi

import (
	"fmt"

	db "github.com/Jimmmy4REAL/bank_tx/db/sqlc"
	"github.com/Jimmmy4REAL/bank_tx/pb"
	"github.com/Jimmmy4REAL/bank_tx/token"
	"github.com/Jimmmy4REAL/bank_tx/util"
)

// Server serves gRPC requests for our banking service. - here init unimplmented-sever (called this way :))
type Server struct {
	pb.UnimplementedBanktxServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server. - no require for routing - generated in protoBuff
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
