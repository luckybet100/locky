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
	"e8906f30-c13b-457f-a142-dac4fd429656": &pb.User{
		FirstName: "Ildar",
		LastName:  "Lutfullin",
	},
	"784fc6d5-5fdd-4b90-baee-71df6caa4c59": &pb.User{
		FirstName: "Evgeny",
		LastName:  "Tushkanov",
	},
	"f833ba18-5e01-449c-a98b-ff5602f4c812": &pb.User{
		FirstName: "Nadezhda",
		LastName:  "Filippova",
	},
	"73878c88-a90d-451d-bfe5-2572acc8f5f6": &pb.User{
		FirstName: "Polina",
		LastName:  "Romanchenko",
	},
}

func (h *Handler) ListUsers(ctx context.Context, in *pb.ListUsersIn) (*pb.ListUsersOut, error) {
	if in == nil {
		return nil, nil
	}
	if in.Key != nil {
		switch in.Key.Key {
		case "7c659e53-9871-4297-9084-d1ce251af68d":
			return &pb.ListUsersOut{
				UserIds: []string{"e8906f30-c13b-457f-a142-dac4fd429656", "784fc6d5-5fdd-4b90-baee-71df6caa4c59"},
				NextKey: &pb.BatchKey{
					Key: "784fc6d5-5fdd-4b90-baee-71df6caa4c59",
				},
			}, nil
		case "784fc6d5-5fdd-4b90-baee-71df6caa4c59":
			return &pb.ListUsersOut{
			UserIds: []string{"e8906f30-c13b-457f-a142-dac4fd429656", "73878c88-a90d-451d-bfe5-2572acc8f5f6"},
		}, nil
		default:
			return nil, nil
		}

	}
	return &pb.ListUsersOut{
		UserIds: []string{"f833ba18-5e01-449c-a98b-ff5602f4c812", "7c659e53-9871-4297-9084-d1ce251af68d"},
		NextKey: &pb.BatchKey{
			Key: "7c659e53-9871-4297-9084-d1ce251af68d",
		},
	}, nil
}
