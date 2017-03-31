// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlexp

import (
	"database/sql/driver"
)

const (
	DialectPostgres = "postgres"
	DialectTSQL     = "tsql"
	DialectMySQL    = "mysql"
	DialectSQLite   = "sqlite"
	DialectOracle   = "oracle"
)

// DriverNamer returns the name of the database and the SQL dialect it
// uses.
type DriverNamer interface {
	// Name of the database management system.
	//
	// Examples:
	//    "posgresql-9.6"
	//    "sqlserver-10.54.32"
	//    "cockroachdb-1.0"
	Name() string

	// Dialect of SQL used in the database.
	Dialect() string
}

// NamerFromDriver returns the DriverNamer from the Driver if
// it is implemented.
func NamerFromDriver(d driver.Driver) DriverNamer {
	dn, _ := d.(DriverNamer)
	return dn
}
