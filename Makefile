.PHONY: build
build:
	go build -o bin/kvserver ./cmd/kvserver
	go build -o bin/kvsctl ./cmd/kvsctl

PB_GO_PACKAGES=`find . -type f -name '*.pb.go'`
.PHONY: clean
clean:
	for pb in $(PB_GO_PACKAGES); do \
		rm -rf $$pb; \
	done

GO_PACKAGES=`go list ./... | grep -v -e test -e proto`
.PHONY: test
test:
	go test $(GO_PACKAGES) -v

.PHONY: proto
proto: clean
	bash ./genproto.sh