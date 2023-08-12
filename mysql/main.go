// Copyright 2023 The YockCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// #include "../cloud.h"
import "C"
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"unsafe"
	"yockcloud/cgo"

	_ "github.com/go-sql-driver/mysql"
)

//export Dial
func Dial(dataSourceName *C.char) *C.component {
	// defer C.free(unsafe.Pointer(dataSourceName))
	conn, err := sql.Open("mysql", C.GoString(dataSourceName))
	if err != nil {
		panic(err)
	}
	ret := cgo.Malloc(C.component{})
	ret.ptr = unsafe.Pointer(conn)
	return ret
}

//export Close
func Close(comp *C.component) *C.char {
	db := cgo.CastPtr[sql.DB](comp.ptr)
	err := db.Close()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Ping
func Ping(conn *C.component) *C.char {
	db := cgo.CastPtr[sql.DB](conn.ptr)
	err := db.Ping()
	if err != nil {
		return C.CString(err.Error())
	}
	return nil
}

//export Query
func Query(conn *C.component, query *C.char) *C.char {
	db := cgo.CastPtr[sql.DB](conn.ptr)
	stmt, err := db.Prepare(C.GoString(query))
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil
	}
	count := len(columns)
	tableData := make([]map[string]any, 0)
	values := make([]any, count)
	valuePtrs := make([]any, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]any)
		for i, col := range columns {
			var v any
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return nil
	}
	return C.CString(string(jsonData))
}

//export Exec
func Exec(conn *C.component, query *C.char) *C.char {
	db := cgo.CastPtr[sql.DB](conn.ptr)
	// defer C.free(unsafe.Pointer(query))
	res, err := db.Exec(C.GoString(query))
	if err != nil {
		return C.CString(err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return C.CString(err.Error())
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return C.CString(err.Error())
	}
	if rowCnt != 0 {
		return C.CString(fmt.Sprintf("Row count (%d) and/or lastId (%d) are wrong.", lastId, rowCnt))
	}
	return nil
}

func main() {}
