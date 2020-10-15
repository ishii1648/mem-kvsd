package app

import (
	"context"
	// "fmt"
	"strconv"
	"testing"
	"time"
	"unsafe"

	pb "github.com/ishii1648/mem-kvsd/cmd/kvserver/app/kvserverpb"
	"github.com/ishii1648/mem-kvsd/pkg/kvs"
	"google.golang.org/grpc"
)

var (
	testKVServerFlags = &KVServerFlags{
		Host:      defaultHost,
		Port:      defaultListenPort,
		DebugMode: true,
	}
	testKey   = "test-key"
	testValue = "test-value"
)

func TestPutAndGetViaRpc(t *testing.T) {
	s := &KVServer{
		host: defaultHost,
		port: defaultListenPort,
		kv:   kvs.GetKvsInstance(),
	}
	go s.Start(SetupSignalHandler())

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	addr := defaultHost + ":" + strconv.Itoa(defaultListenPort)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewKVClient(conn)

	_, p_err := client.Put(ctx, &pb.PutRequest{Key: []byte(testKey), Value: []byte(testValue)})
	if p_err != nil {
		t.Fatalf("fail to Put: %v", err)
	}

	resp, err := client.Get(ctx, &pb.GetRequest{Key: []byte(testKey)})
	if err != nil {
		t.Fatalf("fail to Get: %v", err)
	}

	if toString(resp.PrevKv.Key) != testKey {
		t.Fatalf("no matach key <resp: %s, testKey: %s>", toString(resp.PrevKv.Key), testKey)
	}

	if toString(resp.PrevKv.Value) != testValue {
		t.Fatalf("no matach value <resp: %s, testValue: %s>", toString(resp.PrevKv.Value), testValue)
	}
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
