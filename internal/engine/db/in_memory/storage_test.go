package inmemory

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AnnDutova/in-memory-db/internal/engine/db"
)

var (
	testKey   = "test-key"
	testValue = "test-value"
)

func TestSetGet(t *testing.T) {
	testCases := []struct {
		name        string
		prepareFn   func(s db.Storage)
		outputValue string
		outputFound bool
		input       string
	}{
		{
			name: "valid get",
			prepareFn: func(s db.Storage) {
				s.Set(testKey, testValue)
			},
			input:       testKey,
			outputValue: testValue,
			outputFound: true,
		},
		{
			name:        "not found",
			prepareFn:   func(_ db.Storage) {},
			input:       testKey,
			outputValue: "",
			outputFound: false,
		},
	}

	for _, tc := range testCases {
		storage, err := New()
		require.Nil(t, err)
		require.NotNil(t, storage)

		tc.prepareFn(storage)
		actual, ok := storage.Get(tc.input)
		require.Equal(t, tc.outputValue, actual)
		require.Equal(t, tc.outputFound, ok)
	}
}

func TestDel(t *testing.T) {
	delKey := "del-key"
	testCases := []struct {
		name      string
		prepareFn func(s db.Storage)
		outputLen int
		input     string
	}{
		{
			name: "valid del",
			prepareFn: func(s db.Storage) {
				s.Set(testKey, testValue)
				s.Set(delKey, testValue)
			},
			input:     delKey,
			outputLen: 1,
		},
		{
			name: "not found",
			prepareFn: func(s db.Storage) {
				s.Set(delKey, testValue)
			},
			input:     testKey,
			outputLen: 1,
		},
		{
			name: "valid del",
			prepareFn: func(s db.Storage) {
				s.Set(delKey, testValue)
			},
			input:     delKey,
			outputLen: 0,
		},
	}

	for _, tc := range testCases {
		storage, err := New()
		require.Nil(t, err)
		require.NotNil(t, storage)

		tc.prepareFn(storage)
		storage.Delite(tc.input)

		require.Equal(t, storage.Length(), tc.outputLen)
	}
}
