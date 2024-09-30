package terminal

import (
	"context"
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
)

var users = map[string]*pb.User{
	"25bdd851-631b-4cff-a8c2-0d090b485bed": &pb.User{
		FirstName: "Alexey",
		LastName:  "Kuldoshin",
	},
	"7c659e53-9871-4297-9084-d1ce251af68d": &pb.User{
		FirstName: "Andrey",
		LastName:  "Orel",
	},
}

func (h *Handler) ListUsers(ctx context.Context, in *pb.ListUsersIn) (*pb.ListUsersOut, error) {
	if in == nil {
		return nil, nil
	}
	return &pb.ListUsersOut{
		UserIds: []string{},
	}, nil
}
