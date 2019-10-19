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

// New creates a Rossby server with the specified config or loads the config from the
// environment if no configuration is passed. The server is initialized and validated
// once created but is not running.
func New(options *Config) (r *Server, err error) {
	// Load config from environment if not specified by user.
	if options == nil {
		options = new(Config)
		if err = options.Load(); err != nil {
			return nil, err
		}
	}

	// Set the logging level from the configuration
	SetLogLevel(uint8(options.LogLevel))

	// Create and initialize the server
	r = &Server{config: options}
	if r.db, err = badger.Open((options.DatabaseOptions())); err != nil {
		return nil, err
	}
	return r, nil
}

// Server objects contain the state to run a single rossby server instance that handles
// messages from clients and manages a local database instance.
type Server struct {
	config *Config
	db     *badger.DB
}

// Listen for messages from clients and respond to them.
func (s *Server) Listen() error {
	// Open TCP socket to listen for messages
	sock, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %s", s.config.Address, err)
	}
	defer sock.Close()
	info("listening for requests on %s", s.config.Address)

	return nil
}
