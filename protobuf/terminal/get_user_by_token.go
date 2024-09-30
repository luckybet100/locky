package terminal

import (
	"context"
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
)

func (h *Handler) GetUserByToken(ctx context.Context, in *pb.GetUserByTokenIn) (*pb.GetUserByTokenOut, error) {
	if user, ok := users[in.Token]; ok {
		return &pb.GetUserByTokenOut{User: user}, nil
	} else {
		return nil, nil
	}
}
