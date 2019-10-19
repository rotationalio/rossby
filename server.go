package rossby

import (
	"context"

	pb "github.com/kansaslabs/rossby/pb"
)

// Register implements the Rossby server interface.
func (s *Replica) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return nil, nil
}

// Authorize implements the Rossby server interface.
func (s *Replica) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (*pb.AuthorizeReply, error) {
	return nil, nil
}

// Contact implements the Rossby server interface.
func (s *Replica) Contact(ctx context.Context, in *pb.ContactRequest) (*pb.ContactReply, error) {
	return nil, nil
}

// Fetch implements the Rossby server interface.
func (s *Replica) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.Messages, error) {
	return nil, nil
}

// Deliver implements the Rossby server interface.
func (s *Replica) Deliver(ctx context.Context, in *pb.Messages) (*pb.DeliverResponse, error) {
	return nil, nil
}

// Chat implements the Rossby server interface.
func (s *Replica) Chat(stream pb.Rossby_ChatServer) error {
	return nil
}
