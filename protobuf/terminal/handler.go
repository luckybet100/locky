package terminal

import pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"

type Handler struct {
	pb.UnimplementedGatewayServer
}
