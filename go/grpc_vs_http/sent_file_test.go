package grpc_vs_http

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"github.com/bigpigeon/Test/go/grpc_vs_http/proto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"sync"
	"testing"
)

type FileStorage struct {
	download []byte
}

func (fs *FileStorage) Upload(s info.FileStorage_UploadServer) error {
	d := make([]byte, 0, 100*1024*1024)
	for {
		data, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		d = append(d, data.Data...)
	}
	key := sha1.Sum(d)

	return s.SendAndClose(&info.UploadResponse{
		Key: base64.RawURLEncoding.EncodeToString(key[:]),
	})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

const chunkSize = 4096

func (fs *FileStorage) Download(req *info.DownloadRequest, s info.FileStorage_DownloadServer) error {
	for i := 0; i < len(fs.download); i += chunkSize {
		end := min(i+chunkSize, len(fs.download))
		err := s.Send(&info.DownloadStream{
			Data: fs.download[i:end],
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func get100MData(t *testing.T) []byte {
	d100M := make([]byte, 1024*1024*100)
	readEnd, err := rand.Read(d100M)
	require.NoError(t, err)
	require.Equal(t, readEnd, len(d100M))
	return d100M
}

func TestGrpcSentFile(t *testing.T) {
	li, err := net.Listen("tcp", "localhost:50001")
	assert.NoError(t, err)

	server := grpc.NewServer()

	info.RegisterFileStorageServer(server, &FileStorage{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Serve(li); err != nil {
			panic(err)
		}
	}()
	defer server.GracefulStop()

	d100M := get100MData(t)

	dialer, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	require.NoError(t, err)
	client := info.NewFileStorageClient(dialer)

	memStat := runtime.MemStats{}
	runtime.ReadMemStats(&memStat)
	t.Logf("total memory %d current memory %d", memStat.TotalAlloc, memStat.Alloc)
	{
		stream, err := client.Upload(context.Background())
		require.NoError(t, err)
		for i := 0; i < len(d100M); i += chunkSize {
			end := min(i+chunkSize, len(d100M))
			err := stream.Send(&info.UploadStream{
				Data: d100M[i:end],
			})
			require.NoError(t, err)
		}
		rsp, err := stream.CloseAndRecv()
		require.NoError(t, err)
		t.Log(rsp.Key)
	}
	runtime.ReadMemStats(&memStat)
	t.Logf("total memory %d current memory %d", memStat.TotalAlloc, memStat.Alloc)
}

func TestHttpSentFile(t *testing.T) {
	g := gin.New()
	g.POST("/", func(ctx *gin.Context) {
		var length int
		//_, err := fmt.Scan(ctx.GetHeader("content-length"), &length)
		//require.NoError(t, err)
		t.Log(length)
		data, err := ioutil.ReadAll(ctx.Request.Body)
		require.NoError(t, err)

		key := sha1.Sum(data)
		ctx.JSON(200, gin.H{"Key": base64.RawURLEncoding.EncodeToString(key[:])})
	})

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()

	server := http.Server{Addr: "localhost:50002", Handler: g}
	defer server.Close()

	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if err != nil {
			t.Log(err)
		}
		//require.NoError(t, err)
	}()

	d100M := get100MData(t)
	memStat := runtime.MemStats{}
	runtime.ReadMemStats(&memStat)
	t.Logf("total memory %d current memory %d", memStat.TotalAlloc, memStat.Alloc)

	{
		rsp, err := http.Post("http://localhost:50002/", "application/octet-stream", bytes.NewReader(d100M))
		require.NoError(t, err)
		rspData, err := ioutil.ReadAll(rsp.Body)
		require.NoError(t, err)
		t.Log(string(rspData))
	}
	runtime.ReadMemStats(&memStat)
	t.Logf("total memory %d current memory %d", memStat.TotalAlloc, memStat.Alloc)
}
