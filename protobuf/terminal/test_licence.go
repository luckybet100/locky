package terminal

import (
	"context"
	pb "github.com/Locky-Inc/locky/internal/generated/protobuf/terminal"
)

func (h *Handler) TestLicence(ctx context.Context, in *pb.TestLicenceIn) (*pb.TestLicenceOut, error) {
	if in == nil {
		return &pb.TestLicenceOut{
			Status: pb.LiceneStatus_LICENCE_STATUS_FAILED,
		}, nil
	}
	return &pb.TestLicenceOut{
		Status: pb.LiceneStatus_LICENCE_STATUS_OK,
	}, nil
}
