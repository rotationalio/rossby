package rossby_test

import (
	"os"
	"testing"

	. "github.com/kansaslabs/rossby"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	config := new(Config)
	require.Empty(t, config)

	dbpath, ok := os.LookupEnv("ROSSBY_DATABASE")
	require.True(t, ok)

	require.NoError(t, config.Load())
	require.NotEmpty(t, config)
	require.NotEqual(t, "fixtures/db", config.Database)
	require.Equal(t, config.Database, dbpath)

	dbopts := config.DatabaseOptions()
	require.NotEmpty(t, dbopts)
	require.Equal(t, config.Database, dbopts.Dir)
}
