// Copyright 2023 The YockCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// #include "../cloud.h"
import "C"
import (
	"bytes"
	"encoding/json"
	"net/http"
	"unsafe"
	"yockcloud/cgo"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

//export Dial
func Dial(url *C.char) *C.component {
	mq := &RabbitMQ{}
	var err error
	mq.conn, err = amqp.Dial(C.GoString(url))
	if err != nil {
		panic(err)
	}
	ret := cgo.Malloc(C.component{})
	ret.ptr = unsafe.Pointer(mq)
	return ret
}

type ConsumeOption struct {
	Queue     string `json:"queue"`
	Consumer  string `json:"consumer"`
	AutoAck   bool   `json:"auto_ack"`
	Exclusive bool   `json:"exclusive"`
	NoLocal   bool   `json:"nolocal"`
	NoWait    bool   `json:"nowait"`
	URL       string `json:"url"`
}

//export Consume
func Consume(comp *C.component, conf *C.char) {
	mq := cgo.CastPtr[RabbitMQ](comp.ptr)
	var err error
	if mq.channel == nil {
		mq.channel, err = mq.conn.Channel()
		if err != nil {
			panic(err)
		}
	}
	defer mq.channel.Close()
	opt := ConsumeOption{}
	err = json.Unmarshal([]byte(C.GoString(conf)), &opt)
	if err != nil {
		panic(err)
	}
	msgChan, err := mq.channel.Consume(opt.Queue, opt.Consumer, opt.AutoAck, opt.Exclusive, opt.NoLocal, opt.NoWait, nil)
	if err != nil {
		panic(err)
	}
	for msg := range msgChan {
		http.Post(opt.URL, "application/json;charset=UTF-8", bytes.NewReader(msg.Body))
	}
}

type QueueDeclareOption struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"auto_delete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"nowait"`
}

//export QueueDeclare
func QueueDeclare(comp *C.component, opt *C.char) *C.char {
	mq := cgo.CastPtr[RabbitMQ](comp.ptr)
	var err error
	if mq.channel == nil {
		mq.channel, err = mq.conn.Channel()
		if err != nil {
			panic(err)
		}
	}
	qopt := QueueDeclareOption{}
	err = json.Unmarshal([]byte(C.GoString(opt)), &qopt)
	if err != nil {
		return C.CString(err.Error())
	}
	_, err = mq.channel.QueueDeclare(qopt.Name, qopt.Durable, qopt.AutoDelete, qopt.Exclusive, qopt.NoWait, nil)
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

type ExchangeDeclareOption struct {
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"auto_delete"`
	Internal   bool   `json:"exclusive"`
	NoWait     bool   `json:"nowait"`
}

//export ExchangeDeclare
func ExchangeDeclare(comp *C.component, opt *C.char) *C.char {
	mq := cgo.CastPtr[RabbitMQ](comp.ptr)
	var err error
	if mq.channel == nil {
		mq.channel, err = mq.conn.Channel()
		if err != nil {
			panic(err)
		}
	}
	qopt := ExchangeDeclareOption{}
	err = json.Unmarshal([]byte(C.GoString(opt)), &qopt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = mq.channel.ExchangeDeclare(qopt.Name, qopt.Kind, qopt.Durable, qopt.AutoDelete, qopt.Internal, qopt.NoWait, nil)
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

type QueueBindOption struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	Exchange string `json:"exchange"`
	Nowait   bool   `json:"nowait"`
}

//export QueueBind
func QueueBind(comp *C.component, opt *C.char) *C.char {
	mq := cgo.CastPtr[RabbitMQ](comp.ptr)
	var err error
	if mq.channel == nil {
		mq.channel, err = mq.conn.Channel()
		if err != nil {
			panic(err)
		}
	}
	qopt := QueueBindOption{}
	err = json.Unmarshal([]byte(C.GoString(opt)), &qopt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = mq.channel.QueueBind(qopt.Name, qopt.Key, qopt.Exchange, qopt.Nowait, nil)
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

type Publishing struct {
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}

type PublishOption struct {
	Exchange  string     `json:"exchange"`
	Key       string     `json:"key"`
	Mandatory bool       `json:"mandatory"`
	Immediate bool       `json:"immediate"`
	Msg       Publishing `json:"msg"`
}

//export Publish
func Publish(comp *C.component, opt *C.char) *C.char {
	mq := cgo.CastPtr[RabbitMQ](comp.ptr)
	var err error
	if mq.channel == nil {
		mq.channel, err = mq.conn.Channel()
		if err != nil {
			panic(err)
		}
	}
	qopt := PublishOption{}
	err = json.Unmarshal([]byte(C.GoString(opt)), &qopt)
	if err != nil {
		return C.CString(err.Error())
	}
	err = mq.channel.Publish(qopt.Exchange, qopt.Key, qopt.Mandatory, qopt.Immediate, amqp.Publishing{
		ContentType: qopt.Msg.ContentType,
		Body:        []byte(qopt.Msg.Body),
	})
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

func main() {}
