// Copyright 2023 The YockCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// #include "../cloud.h"
import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unsafe"
	"yockcloud/cgo"

	clientv3 "go.etcd.io/etcd/client/v3"
)

//export Dial
func Dial(config *C.char) *C.component {
	conf := clientv3.Config{}
	err := json.Unmarshal([]byte(C.GoString(config)), &conf)
	if err != nil {
		panic(err)
	}
	cli, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}
	ret := cgo.Malloc(C.component{})
	ret.ptr = unsafe.Pointer(cli)
	return ret
}

type WatchOption struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

//export Watch
func Watch(comp *C.component, c *C.char) {
	cli := cgo.CastPtr[clientv3.Client](comp.ptr)
	opt := WatchOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		panic(err)
	}
	rch := cli.Watch(context.Background(), opt.Key)
	for wresp := range rch {
		for _, e := range wresp.Events {
			tmp := []byte(fmt.Sprintf(`{"key": "%s", "value": "%s"}`, e.Kv.Key, e.Kv.Value))
			http.Post(opt.URL, "application/json;charset=UTF-8", bytes.NewReader(tmp))
		}
	}
}

//export Put
func Put(comp *C.component, key, value *C.char) *C.char {
	cli := cgo.CastPtr[clientv3.Client](comp.ptr)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ret, err := cli.Put(ctx, C.GoString(key), C.GoString(value))
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"msg": nil,
			"err": err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"msg": ret,
		"err": nil,
	}))
}

type KV struct {
	// key is the key in bytes. An empty key is not allowed.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the key. A deletion resets
	// the version to zero and any modification of the key
	// increases its version.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// value is the value held by the key, in bytes.
	Value string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	// lease is the ID of the lease that attached to key.
	// When the attached lease expires, the key will be deleted.
	// If lease is 0, then no lease is attached to the key.
	Lease int64 `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}

//export Get
func Get(comp *C.component, key *C.char) *C.char {
	cli := cgo.CastPtr[clientv3.Client](comp.ptr)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ret, err := cli.Get(ctx, C.GoString(key))
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"msg": nil,
			"err": err.Error(),
		}))
	}
	kvs := []KV{}
	for _, kv := range ret.Kvs {
		kvs = append(kvs, KV{
			Key:            string(kv.Key),
			Value:          string(kv.Value),
			Version:        kv.Version,
			CreateRevision: kv.CreateRevision,
			ModRevision:    kv.ModRevision,
			Lease:          kv.Lease,
		})
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"msg": kvs,
		"err": nil,
	}))
}

//export Close
func Close(comp *C.component) {
	cli := cgo.CastPtr[clientv3.Client](comp.ptr)
	cli.Close()
}

func Ping(comp *C.component) {
	// cli := cgo.CastPtr[clientv3.Client](comp.ptr)
}

func main() {}
