// Copyright 2023 The YockCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// #include "../cloud.h"
import "C"
import (
	"context"
	"encoding/json"
	"time"
	"unsafe"
	"yockcloud/cgo"

	"github.com/go-redis/redis/v8"
)

func main() {}

type RClient struct {
	*redis.Client
	ctx context.Context
}

//export Dial
func Dial(conf *C.char) *C.component {
	opt := &redis.Options{}
	err := json.Unmarshal([]byte(C.GoString(conf)), opt)
	if err != nil {
		panic(err)
	}
	rdb := &RClient{
		Client: redis.NewClient(opt),
		ctx:    context.Background(),
	}
	ret := cgo.Malloc(C.component{})
	ret.ptr = unsafe.Pointer(rdb)
	return ret
}

//export Ping
func Ping(comp *C.component) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	_, err := rc.Ping(rc.ctx).Result()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Close
func Close(comp *C.component) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Close()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Set
func Set(comp *C.component, key *C.char, value *C.char, expire C.double) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Set(rc.ctx, C.GoString(key), C.GoString(value), time.Duration(expire)).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Get
func Get(comp *C.component, key *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	val, err := rc.Get(rc.ctx, C.GoString(key)).Result()
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString(val)
}

//export Del
func Del(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	keys := []string{}
	err := json.Unmarshal([]byte(C.GoString(c)), &keys)
	if err != nil {
		return C.CString(err.Error())
	}
	err = rc.Client.Del(rc.ctx, keys...).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export SetNX
func SetNX(comp *C.component, key *C.char, value *C.char, expire C.double) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Client.SetNX(rc.ctx, C.GoString(key), C.GoString(value), time.Duration(expire)).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Do
func Do(comp *C.component, cmd *C.char) *C.char {
	args := []any{}
	err := json.Unmarshal([]byte(C.GoString(cmd)), &args)
	if err != nil {
		return C.CString(err.Error())
	}
	rc := cgo.CastPtr[RClient](comp.ptr)
	res, err := rc.Client.Do(rc.ctx, args...).Result()
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"err":  err.Error(),
			"data": res,
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

type EvalOption struct {
	Script string   `json:"script"`
	Keys   []string `json:"keys"`
	Args   []any    `json:"args"`
}

//export Eval
func Eval(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	opt := &EvalOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), opt)
	if err != nil {
		return C.CString(err.Error())
	}
	res, err := rc.Client.Eval(rc.ctx, opt.Script, opt.Keys, opt.Args...).Result()
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"err":  err.Error(),
			"data": res,
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

type HSetOption struct {
	Key    string `json:"key"`
	Fields []any  `json:"field"`
}

//export HSet
func HSet(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	opt := HSetOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = rc.Client.HSet(rc.ctx, opt.Key, opt.Fields...).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export HGet
func HGet(comp *C.component, key *C.char, value *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	val, err := rc.Client.HGet(rc.ctx, C.GoString(key), C.GoString(value)).Result()
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": val,
		"err":  nil,
	}))
}

type HDelOption struct {
	Key    string   `json:"key"`
	Fields []string `json:"field"`
}

//export HDel
func HDel(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	opt := HDelOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = rc.Client.HDel(rc.ctx, opt.Key, opt.Fields...).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

type LPushOption struct {
	Key    string `json:"key"`
	Values []any  `json:"values"`
}

//export LPush
func LPush(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	opt := LPushOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = rc.Client.LPush(rc.ctx, opt.Key, opt.Values...).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

type RPushOption struct {
	Key    string `json:"key"`
	Values []any  `json:"values"`
}

//export RPush
func RPush(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	opt := RPushOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = rc.Client.RPush(rc.ctx, opt.Key, opt.Values...).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

type LRangeOption struct {
	Key   string `json:"key"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
}

//export LRange
func LRange(comp *C.component, c *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	opt := LRangeOption{}
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	res := rc.Client.LRange(rc.ctx, opt.Key, opt.Start, opt.End).Val()
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": res,
		"err":  nil,
	}))
}

//export Incr
func Incr(comp *C.component, key *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Client.Incr(rc.ctx, C.GoString(key)).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Decr
func Decr(comp *C.component, key *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Client.Decr(rc.ctx, C.GoString(key)).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export LPop
func LPop(comp *C.component, key *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Client.LPop(rc.ctx, C.GoString(key)).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export RPop
func RPop(comp *C.component, key *C.char) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	err := rc.Client.RPop(rc.ctx, C.GoString(key)).Err()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Scan
func Scan(comp *C.component, prefix *C.char, count C.longlong) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	var cursor uint64
	match := "*" + C.GoString(prefix) + "*"
	keys, _, err := rc.Client.Scan(rc.ctx, cursor, match, int64(count)+1).Result()
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": keys,
		"err":  nil,
	}))
}

//export SScan
func SScan(comp *C.component, key *C.char, prefix *C.char, count C.longlong) *C.char {
	rc := cgo.CastPtr[RClient](comp.ptr)
	var cursor uint64
	match := "*" + C.GoString(prefix) + "*"
	keys, _, err := rc.Client.SScan(rc.ctx, C.GoString(key), cursor, match, int64(count)+1).Result()
	if err != nil {
		return (*C.char)(cgo.JSON(cgo.MAP{
			"data": nil,
			"err":  err.Error(),
		}))
	}
	return (*C.char)(cgo.JSON(cgo.MAP{
		"data": keys,
		"err":  nil,
	}))
}
