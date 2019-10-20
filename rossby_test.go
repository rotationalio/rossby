package rossby_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/kansaslabs/rossby"
	pb "github.com/kansaslabs/rossby/pb"
	"github.com/stretchr/testify/require"
)

var testDir string

// TestMain runs the main goroutine and can run setup and teardown functions. In this
// case we configure the environment and the database for all tests.
func TestMain(m *testing.M) {
	// Setup the tests
	var err error
	if testDir, err = ioutil.TempDir("", "rossby_tests"); err != nil {
		fmt.Fprintln(os.Stderr, "could not create temp test directory")
		os.Exit(1)
	}

	os.Setenv("ROSSBY_DATABASE", filepath.Join(testDir, "db"))

	// Execute the tests
	code := m.Run()

	// Teardown the tests
	os.RemoveAll(testDir)
	os.Exit(code)
}

//===========================================================================
// Test Helpers
//===========================================================================

// Check if a path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func requirePathExists(t *testing.T, path string) {
	exists, err := pathExists(path)
	require.NoError(t, err)
	require.True(t, exists)
}

func requirePathNotExists(t *testing.T, path string) {
	exists, err := pathExists(path)
	require.NoError(t, err)
	require.False(t, exists)
}

//===========================================================================
// Rossby Tests
//===========================================================================

func TestReplicaService(t *testing.T) {
	// Test preconditions
	dbpath, ok := os.LookupEnv("ROSSBY_DATABASE")
	require.True(t, ok)
	requirePathNotExists(t, dbpath)

	// Test correct instantiation of the replica
	replica, err := New(nil)
	require.NoError(t, err)
	require.Implements(t, (*pb.RossbyServer)(nil), replica)
	require.NotEmpty(t, replica)

	// Test correct initialization of the replica
	require.Equal(t, LogLevel(), "caution")
	requirePathExists(t, dbpath)
}
