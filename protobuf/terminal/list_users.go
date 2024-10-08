package terminal

import (
	"context"
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
)

var users = map[string]*pb.User{
	"25bdd851-631b-4cff-a8c2-0d090b485bed": &pb.User{
		Id:        "25bdd851-631b-4cff-a8c2-0d090b485bed",
		FirstName: "Alexey",
		LastName:  "Kuldoshin",
	},
	"7c659e53-9871-4297-9084-d1ce251af68d": &pb.User{
		Id:        "7c659e53-9871-4297-9084-d1ce251af68d",
		FirstName: "Andrey",
		LastName:  "Orel",
	},
	"e8906f30-c13b-457f-a142-dac4fd429656": &pb.User{
		Id:        "e8906f30-c13b-457f-a142-dac4fd429656",
		FirstName: "Ildar",
		LastName:  "Lutfullin",
	},
	"784fc6d5-5fdd-4b90-baee-71df6caa4c59": &pb.User{
		Id:        "784fc6d5-5fdd-4b90-baee-71df6caa4c59",
		FirstName: "Evgeny",
		LastName:  "Tushkanov",
	},
	"f833ba18-5e01-449c-a98b-ff5602f4c812": &pb.User{
		Id:        "f833ba18-5e01-449c-a98b-ff5602f4c812",
		FirstName: "Nadezhda",
		LastName:  "Filippova",
	},
	"73878c88-a90d-451d-bfe5-2572acc8f5f6": &pb.User{
		Id:        "73878c88-a90d-451d-bfe5-2572acc8f5f6",
		FirstName: "Polina",
		LastName:  "Romanchenko",
	},
	"aec366b3-12b6-48e5-9175-134ba7bceff3": &pb.User{
		Id:        "aec366b3-12b6-48e5-9175-134ba7bceff3",
		FirstName: "Jo",
		LastName:  "",
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
				UserIds: []string{"f833ba18-5e01-449c-a98b-ff5602f4c812", "73878c88-a90d-451d-bfe5-2572acc8f5f6", "aec366b3-12b6-48e5-9175-134ba7bceff3"},
			}, nil
		default:
			return nil, nil
		}

	}
	return &pb.ListUsersOut{
		UserIds: []string{"25bdd851-631b-4cff-a8c2-0d090b485bed", "7c659e53-9871-4297-9084-d1ce251af68d"},
		NextKey: &pb.BatchKey{
			Key: "7c659e53-9871-4297-9084-d1ce251af68d",
		},
	}, nil
}
