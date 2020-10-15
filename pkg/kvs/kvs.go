package kvs

import (
	"fmt"
	"unsafe"

	kvpb "github.com/ishii1648/mem-kvsd/pkg/kvs/kvpb"
	// "github.com/ishii1648/mem-kvsd/pkg/log"
)

type KV interface {
	Get(key []byte) (*kvpb.KeyValue, error)
	Put(key []byte, value []byte) (*kvpb.KeyValue, error)
}

type kvs struct {
	keyList  []string
	keyValue []*kvpb.KeyValue
}

var sharedKvsInstance = &kvs{}

func GetKvsInstance() *kvs {
	return sharedKvsInstance
}

func (k *kvs) Get(key []byte) (*kvpb.KeyValue, error) {
	_key := toString(key)

	keyIndex := -1
	for i, key := range k.keyList {
		if _key == key {
			keyIndex = i
		}
	}

	if keyIndex == -1 {
		return nil, fmt.Errorf("no key list")
	}

	return k.keyValue[keyIndex], nil
}

func (k *kvs) Put(key []byte, value []byte) (*kvpb.KeyValue, error) {
	keyValue := &kvpb.KeyValue{
		Key:   key,
		Value: value,
	}
	k.keyValue = append(k.keyValue, keyValue)
	k.keyList = append(k.keyList, toString(key))

	return keyValue, nil
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
