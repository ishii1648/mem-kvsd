package kvs

import (
	// "fmt"
	"testing"
)

var (
	testKey   = "test-key"
	testValue = "test-value"
)

func TestPutAndGet(t *testing.T) {
	kvs := GetKvsInstance()

	_, err := kvs.Put([]byte(testKey), []byte(testValue))
	if err != nil {
		t.Fatalf("fail to Put: %v", err)
	}

	resp, err := kvs.Get([]byte(testKey))
	if err != nil {
		t.Fatalf("fail to Get: %v", err)
	}

	if toString(resp.Key) != testKey {
		t.Fatalf("no matach key <resp: %s, testKey: %s>", toString(resp.Key), testKey)
	}

	if toString(resp.Value) != testValue {
		t.Fatalf("no matach value <resp: %s, testValue: %s>", toString(resp.Value), testValue)
	}

	// fmt.Printf("<key: %s value: %s>", toString(resp.Key), toString(resp.Value))
}
