package terminal

import (
	"context"
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
)

func (h *Handler) GetUserById(ctx context.Context, in *pb.GetUserByIdIn) (*pb.GetUserByIdOut, error) {
	if user, ok := users[in.Id]; ok {
		return &pb.GetUserByIdOut{User: user}, nil
	} else {
		return nil, nil
	}
}
