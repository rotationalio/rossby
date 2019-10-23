package rossby_test

import (
	"context"
	"testing"

	pb "github.com/kansaslabs/rossby/pb"
	"github.com/stretchr/testify/require"
)

//===========================================================================
// Test Server RPC Calls
//===========================================================================

func TestRegister(t *testing.T) {
	req := &pb.RegisterRequest{
		Contacts: []*pb.Contact{
			{Type: pb.ContactType_EMAIL, Contact: "squirrel@example.com"},
			{Type: pb.ContactType_PHONE, Contact: "+14155552671"},
		},
	}

	r := makeReplica(t)
	rep, err := r.Register(context.Background(), req)
	require.NoError(t, err)
	require.False(t, rep.Success)
}

func TestAuthorize(t *testing.T) {
	req := &pb.AuthorizeRequest{
		Device:        "7991079d-5de9-5403-bd3b-3136e964ff14",
		Authorization: "331908",
	}

	r := makeReplica(t)
	rep, err := r.Authorize(context.Background(), req)
	require.NoError(t, err)

	// Check the reply
	require.False(t, rep.Success)
	require.Zero(t, rep.Token)
	require.Zero(t, rep.Inbox)

	// Check the errors
	errs := rep.Errors.Deserialize()
	require.Len(t, errs, 1)
	require.EqualError(t, errs, "[99] the authorize rpc is not implemented yet")
}

func TestContact(t *testing.T) {
	req := &pb.ContactRequest{
		Contacts: []*pb.Contact{
			{Type: pb.ContactType_EMAIL, Contact: "squirrel@example.com"},
			{Type: pb.ContactType_PHONE, Contact: "+14155552671"},
		},
	}

	r := makeReplica(t)
	rep, err := r.Contact(context.Background(), req)
	require.NoError(t, err)

	// Check the reply
	require.False(t, rep.Success)
	require.Zero(t, rep.Pubkey)

	// Check the errors
	errs := rep.Errors.Deserialize()
	require.Len(t, errs, 1)
	require.EqualError(t, errs, "[99] the contact rpc is not implemented yet")
}

func TestFetch(t *testing.T) {}

func TestDeliver(t *testing.T) {}

func TestChat(t *testing.T) {
	r := makeReplica(t)
	err := r.Chat(nil)
	require.EqualError(t, err, "[99] the streaming chat rpc is not implemented yet")
}
