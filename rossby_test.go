package rossby_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"

	. "github.com/kansaslabs/rossby"
	pb "github.com/kansaslabs/rossby/pb"
	"github.com/stretchr/testify/require"
)

var (
	testDir string
	replica *Replica
	mu      sync.Mutex
)

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

// Ensures that the replica object is a singleton, otherwise we would end up trying to
// create multiple databases on the same path, which will fail because badger locks the
// directory at the os level. Any test that requires a new replica should either use
// this function or use a different database path.
func makeReplica(t *testing.T) *Replica {
	mu.Lock()
	defer mu.Unlock()

	if replica == nil {
		var err error
		replica, err = New(nil)
		require.NoError(t, err)
	}

	return replica
}

// Check if a path exists, returns an error if permission denied or other OS error.
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

// Use require to check that a path exists (asserts no OS error)
func requirePathExists(t *testing.T, path string) {
	exists, err := pathExists(path)
	require.NoError(t, err)
	require.True(t, exists)
}

// Use require to ensure that a path does not exist (asserts no OS error)
func requirePathNotExists(t *testing.T, path string) {
	exists, err := pathExists(path)
	require.NoError(t, err)
	require.False(t, exists)
}

//===========================================================================
// Rossby Tests
//===========================================================================

func TestReplicaInitialization(t *testing.T) {
	// Test preconditions
	opts := &Config{}
	require.NoError(t, opts.Load())

	// NOTE: must use new database path to prevent conflicts with other tests.
	opts.Database = filepath.Join(testDir, "altdb")
	requirePathNotExists(t, opts.Database)

	// Test correct instantiation of the replica
	replica, err := New(opts)
	require.NoError(t, err)
	require.Implements(t, (*pb.RossbyServer)(nil), replica)
	require.NotEmpty(t, replica)

	// Test correct initialization of the replica
	require.Equal(t, LogLevel(), "caution")
	requirePathExists(t, opts.Database)
}
