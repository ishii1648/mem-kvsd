#!/usr/bin/env bash
#
# Generate all etcd protobuf bindings.
# Run from repository root.
#
set -e

MEM_KVS_ROOT=${PWD}
GOGOPROTO_PATH="${GOPATH}/src/github.com/gogo/protobuf"
KV_PB_PATH="${PWD}/pkg/kv/kvpb"

# directories containing protos to be built
DIRS="./pkg/kvs/kvpb ./cmd/kvserver/app/kvserverpb"

for dir in ${DIRS}; do
	pushd "${dir}"
        protoc --gofast_out=plugins=grpc,import_prefix=github.com/ishii1648/mem-kvsd/:. -I=".:${GOGOPROTO_PATH}:${MEM_KVS_ROOT}" ./*.proto

        sed -i.bak -E 's/github\.com\/ishii1648\/mem-kvs\/(gogoproto|github\.com|golang\.org|google\.golang\.org)/\1/g' ./*.pb.go
		sed -i.bak -E 's/github\.com\/ishii1648\/mem-kvs\/(errors|fmt|io|math|context)/\1/g' ./*.pb.go

		rm -f ./*.bak
		gofmt -s -w ./*.pb.go
		goimports -w ./*.pb.go
    popd
done