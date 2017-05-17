// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlexp

import (
	"context"
	"database/sql/driver"
	"testing"

	ms "github.com/denisenkom/go-mssqldb"
)

func TestMSSQL(t *testing.T) {
	ctx := context.Background()
	var driver driver.Driver = &ms.MssqlDriver{}
	q, err := QuoterFromDriver(driver, ctx)
	if err != nil {
		t.Fatal("failed to get driver", err)
	}

	qs := q.Value("It's")
	wanted := "'It''s'"
	if qs != wanted {
		t.Errorf("quote value failed: got %s wanted %s", qs, wanted)
	}
}
