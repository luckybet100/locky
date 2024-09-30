package terminal

import (
	"context"
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
)

func (h *Handler) SetUserFace(ctx context.Context, in *pb.SetUserFaceIn) (*pb.SetUserFaceOut, error) {
	if _, ok := users[in.UserId]; ok {
		users[in.UserId].Face = in.Face
		return &pb.SetUserFaceOut{User: users[in.UserId]}, nil
	} else {
		return nil, nil
	}
}
