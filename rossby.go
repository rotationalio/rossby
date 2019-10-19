/*
Package rossby implements a distributed, encrypted message service that is intended to
be geo-replicated. Rossby's goal is to facilitate the sending and receiving of messages
while also ensuring that as much work as possible is pushed to the client side.
Requiring clients to handle most of the encryption and message management ensures that
rossby has as little detail as possible to be exposed to any security vulnerabilities.
*/
package rossby

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/bbengfort/x/noplog"
	"github.com/dgraph-io/badger"
	pb "github.com/kansaslabs/rossby/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Initialize the package and random numbers, etc.
func init() {
	// Set the random seed to something different each time.
	rand.Seed(time.Now().UnixNano())

	// Initialize our debug logging with our prefix
	SetLogger(log.New(os.Stdout, "[rossby] ", log.Lmicroseconds))
	cautionCounter = new(counter)
	cautionCounter.init()

	// Stop the grpc verbose logging
	grpclog.SetLogger(noplog.New())
}

// New creates a Rossby replica with the specified config or loads the config from the
// environment if no configuration is passed. The replica is initialized and validated
// once created but is not running.
func New(options *Config) (r *Replica, err error) {
	// Load config from environment if not specified by user.
	if options == nil {
		options = new(Config)
		if err = options.Load(); err != nil {
			return nil, err
		}
	}

	// Set the logging level from the configuration
	SetLogLevel(uint8(options.LogLevel))

	// Create and initialize the replica
	r = &Replica{config: options}
	if r.db, err = badger.Open((options.DatabaseOptions())); err != nil {
		return nil, err
	}
	return r, nil
}

// Replica objects contain the state to run a single rossby replica instance which
// manages a local database instance and handles messages from clients and peers to
// distribute messages across the network. There is always one replica object per
// Rossby process.
type Replica struct {
	config *Config
	db     *badger.DB
}

// Listen for messages from clients and respond to them.
func (r *Replica) Listen() error {
	// Open TCP socket to listen for messages
	sock, err := net.Listen("tcp", r.config.Address)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %s", r.config.Address, err)
	}
	defer sock.Close()
	status("listening for requests on %s", r.config.Address)

	// Initialize and run the gRPC server
	srv := grpc.NewServer()
	pb.RegisterRossbyServer(srv, r)

	if err := srv.Serve(sock); err != nil {
		return err
	}

	return nil
}
