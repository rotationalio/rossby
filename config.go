package rossby

import (
	"github.com/dgraph-io/badger"
	"github.com/kelseyhightower/envconfig"
)

// The prefix used to load environment variables with envconfig
const envvarPrefix = "rossby"

// Config stores the optional settings that modify the behavior of rossby at runtime.
// A config can be created either through default values specified on the struct tags,
// or fetched from environment variables prefixed with ROSSBY_.
//
// TODO: change the database default path after development
// TODO: allow configuration from cli context or from YAML file
type Config struct {
	Address  string `default:":1205"`                   // the host and port to bind the server to ($ROSSBY_ADDRESS)
	Database string `default:"fixtures/db"`             // the path to the directory of the badger database ($ROSSBY_DATABASE)
	LogLevel int    `default:"3" envconfig:"log_level"` // verbosity of logging, lower is more verbose ($ROSSBY_LOG_LEVEL)
}

// Load the configuration from the environment and set default values.
func (c *Config) Load() (err error) {
	return envconfig.Process(envvarPrefix, c)
}

// DatabaseOptions returns the badger configuration to open and access the database.
func (c *Config) DatabaseOptions() badger.Options {
	return badger.DefaultOptions(c.Database)
}
