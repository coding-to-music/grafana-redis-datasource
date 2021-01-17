package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryTMScan(t *testing.T) {
	t.Parallel()
	t.Run("should process cursor", func(t *testing.T) {
		t.Parallel()
		client := testClient{
			rcv: []interface{}{
				[]byte("24"),
				[]interface{}{
					[]byte("test:string"),
					[]byte("test:stream"),
					[]byte("test:set"),
					[]byte("test:list"),
					[]byte("test:float"),
					[]byte("test:hash"),
				},
			},
			batchRcv: [][]interface{}{
				{
					"string",
					"stream",
					"set",
					"list",
					"string",
					"hash",
				},
				{
					int64(59),
					int64(612),
					int64(265),
					int64(140),
					int64(59),
					int64(108),
				},
			},
			err: nil,
		}
		resp := queryTMScan(queryModel{Command: "tmscan", Match: "test:*", Count: 100, Cursor: "0"}, &client)
		require.Len(t, resp.Frames, 2)
		require.Len(t, resp.Frames[0].Fields, 3)
		require.Len(t, resp.Frames[1].Fields, 1)
		require.Equal(t, 1, resp.Frames[1].Fields[0].Len())
		require.Equal(t, 6, resp.Frames[0].Fields[0].Len())
		require.Equal(t, 6, resp.Frames[0].Fields[1].Len())
		require.Equal(t, 6, resp.Frames[0].Fields[2].Len())
		require.Equal(t, "24", resp.Frames[1].Fields[0].At(0))

		require.Equal(t, "test:string", resp.Frames[0].Fields[0].At(0))
		require.Equal(t, "test:stream", resp.Frames[0].Fields[0].At(1))
		require.Equal(t, "test:set", resp.Frames[0].Fields[0].At(2))
		require.Equal(t, "test:list", resp.Frames[0].Fields[0].At(3))
		require.Equal(t, "test:float", resp.Frames[0].Fields[0].At(4))
		require.Equal(t, "test:hash", resp.Frames[0].Fields[0].At(5))

		require.Equal(t, "string", resp.Frames[0].Fields[1].At(0))
		require.Equal(t, "stream", resp.Frames[0].Fields[1].At(1))
		require.Equal(t, "set", resp.Frames[0].Fields[1].At(2))
		require.Equal(t, "list", resp.Frames[0].Fields[1].At(3))
		require.Equal(t, "string", resp.Frames[0].Fields[1].At(4))
		require.Equal(t, "hash", resp.Frames[0].Fields[1].At(5))

		require.Equal(t, int64(59), resp.Frames[0].Fields[2].At(0))
		require.Equal(t, int64(612), resp.Frames[0].Fields[2].At(1))
		require.Equal(t, int64(265), resp.Frames[0].Fields[2].At(2))
		require.Equal(t, int64(140), resp.Frames[0].Fields[2].At(3))
		require.Equal(t, int64(59), resp.Frames[0].Fields[2].At(4))
		require.Equal(t, int64(108), resp.Frames[0].Fields[2].At(5))

	})

	t.Run("should handle error during CURSOR", func(t *testing.T) {
		t.Parallel()
		client := testClient{
			rcv:      nil,
			batchRcv: nil,
			err:      errors.New("error when call cursor")}
		resp := queryTMScan(queryModel{Command: "tmscan", Match: "test:*", Count: 100}, &client)
		require.EqualError(t, resp.Error, "error when call cursor")
	})

	t.Run("should handle error during first batch", func(t *testing.T) {
		t.Parallel()
		client := testClient{
			rcv: []interface{}{
				[]byte("24"),
				[]interface{}{
					[]byte("test:string"),
					[]byte("test:stream"),
					[]byte("test:set"),
					[]byte("test:list"),
					[]byte("test:float"),
					[]byte("test:hash"),
				},
			},
			batchRcv: [][]interface{}{
				{
					"string",
					"stream",
					"set",
					"list",
					"string",
					"hash",
				},
				{
					int64(59),
					int64(612),
					int64(265),
					int64(140),
					int64(59),
					int64(108),
				},
			},
			batchErr: []error{errors.New("error when batch types"), nil},
			err:      nil,
		}
		resp := queryTMScan(queryModel{Command: "tmscan", Match: "test:*", Count: 100}, &client)
		require.EqualError(t, resp.Error, "error when batch types")
	})
	t.Run("should handle error during second batch", func(t *testing.T) {
		t.Parallel()
		client := testClient{
			rcv: []interface{}{
				[]byte("24"),
				[]interface{}{
					[]byte("test:string"),
					[]byte("test:stream"),
					[]byte("test:set"),
					[]byte("test:list"),
					[]byte("test:float"),
					[]byte("test:hash"),
				},
			},
			batchRcv: [][]interface{}{
				{
					"string",
					"stream",
					"set",
					"list",
					"string",
					"hash",
				},
				{
					int64(59),
					int64(612),
					int64(265),
					int64(140),
					int64(59),
					int64(108),
				},
			},
			batchErr: []error{nil, errors.New("error when batch memory")},
			err:      nil,
		}
		resp := queryTMScan(queryModel{Command: "tmscan", Match: "test:*", Count: 100}, &client)
		require.EqualError(t, resp.Error, "error when batch memory")
	})
}