package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/ishii1648/mem-kvsd/cmd/kvserver/app/kvserverpb"
	"github.com/ishii1648/mem-kvsd/pkg/kvs"
	"github.com/ishii1648/mem-kvsd/pkg/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	componentKvs = "kvserver"
)

func NewKVServerCommand() *cobra.Command {
	kvserverFlags := NewKVServerFlags()

	cmd := &cobra.Command{
		Use:   componentKvs,
		Short: "run distributed-memory-kvs",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetLevel(GetLogLevel(kvserverFlags.DebugMode))

			return Run(kvserverFlags, SetupSignalHandler())
		},
	}

	fs := cmd.Flags()
	kvserverFlags.Set(fs)

	return cmd
}

func Run(f *KVServerFlags, stopCh <-chan struct{}) error {
	kvServer := NewKVServer(f)
	return kvServer.Start(stopCh)
}

func SetupSignalHandler() <-chan struct{} {
	shutdownHandler := make(chan os.Signal, 2)

	stop := make(chan struct{})
	signal.Notify(shutdownHandler, []os.Signal{syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		<-shutdownHandler
		close(stop)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

type KVServer struct {
	host string
	port int
	kv   kvs.KV
}

func NewKVServer(f *KVServerFlags) *KVServer {
	return &KVServer{
		host: f.Host,
		port: f.Port,
		kv:   kvs.GetKvsInstance(),
	}
}

func (s *KVServer) Start(stopCh <-chan struct{}) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKVServer(grpcServer, s)
	reflection.Register(grpcServer)

	go func() {
		<-stopCh
		grpcServer.GracefulStop()
	}()

	return grpcServer.Serve(lis)
}

func (s *KVServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	if err := checkGetRequest(r); err != nil {
		return nil, err
	}

	resp, err := s.kv.Get(r.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Header: makeHeader(200),
		PrevKv: resp,
	}, nil
}

func (s *KVServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	if err := checkPutRequest(r); err != nil {
		return nil, err
	}

	resp, err := s.kv.Put(r.Key, r.Value)
	if err != nil {
		return nil, err
	}

	return &pb.PutResponse{
		Header: makeHeader(200),
		PrevKv: resp,
	}, nil
}

func checkGetRequest(r *pb.GetRequest) error {
	return nil
}

func checkPutRequest(r *pb.PutRequest) error {
	return nil
}

func makeHeader(statusCode int64) *pb.ResponseHeader {
	return &pb.ResponseHeader{StatusCode: statusCode}
}
