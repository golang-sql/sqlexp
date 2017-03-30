// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlexp

import (
	"database/sql/driver"
	"reflect"
	"strings"
)

// Quoter returns safe and valid SQL strings to use when building a SQL text.
type Quoter interface {
	// ID quotes identifiers such as schema, table, or column names.
	// ID does not operate on multipart identifiers such as "public.Table",
	// it only operates on single identifiers such as "public" and "Table".
	ID(name string) string

	// Value quotes database values such as string or []byte types as strings
	// that are suitable and safe to embed in SQL text. The returned value
	// of a string will include all surrounding quotes.
	//
	// If a value type is not supported it must panic.
	Value(v interface{}) string
}

// DriverQuoter returns a Quoter interface and is suitable for extending
// the driver.Driver type.
type DriverQuoter interface {
	Quoter() Quoter
}

// FromDriver takes a database driver, often obtained through a sql.DB.Driver
// call or from using it directly to get the quoter interface.
//
// Currently MssqlDriver is hard-coded to also return a valided Quoter.
func FromDriver(d driver.Driver) Quoter {
	if q, is := d.(DriverQuoter); is {
		return q.Quoter()
	}
	dv := reflect.ValueOf(d)
	switch dv.Type().String() {
	default:
		return nil
	case "*mssql.MssqlDriver":
		return sqlServerQuoter{}
	}
}

type sqlServerQuoter struct{}

func (sqlServerQuoter) ID(name string) string {
	return "[" + strings.Replace(name, "]", "]]", -1) + "]"
}
func (sqlServerQuoter) Value(v interface{}) string {
	switch v := v.(type) {
	default:
		panic("unsupported value")
	case string:
		return "'" + strings.Replace(v, "'", "''", -1) + "'"
	}
}
