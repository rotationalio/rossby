package rossby

import (
	"context"

	pb "github.com/kansaslabs/rossby/pb"
)

// Register implements the Rossby server interface.
func (r *Replica) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return nil, nil
}

// Authorize implements the Rossby server interface.
func (r *Replica) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (out *pb.AuthorizeReply, err error) {
	nyi := make(pb.Errors, 0, 1)
	nyi = append(nyi, pb.Errorf(99, "the authorize rpc is not implemented yet"))

	// Authorize is current not implemented so return a not implemented error.
	out = &pb.AuthorizeReply{
		Success: false,
		Token:   "",
		Inbox:   "",
		Errors:  nyi.Serialize(),
	}

	return out, nil
}

// Contact implements the Rossby server interface.
func (r *Replica) Contact(ctx context.Context, in *pb.ContactRequest) (out *pb.ContactReply, err error) {
	nyi := make(pb.Errors, 0, 1)
	nyi = append(nyi, pb.Errorf(99, "the contact rpc is not implemented yet"))

	// Authorize is current not implemented so return a not implemented error.
	out = &pb.ContactReply{
		Success: false,
		Pubkey:  "",
		Errors:  nyi.Serialize(),
	}

	return out, nil
}

// Fetch implements the Rossby server interface.
func (r *Replica) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.Messages, error) {
	return nil, nil
}

// Deliver implements the Rossby server interface.
func (r *Replica) Deliver(ctx context.Context, in *pb.Messages) (*pb.DeliverResponse, error) {
	return nil, nil
}

// Chat implements the Rossby server interface.
func (r *Replica) Chat(stream pb.Rossby_ChatServer) error {
	return pb.Errorf(99, "the streaming chat rpc is not implemented yet")
}
