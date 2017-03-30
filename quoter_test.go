package sqlexp

import (
	"database/sql/driver"
	"testing"

	ms "github.com/denisenkom/go-mssqldb"
)

func TestMSSQL(t *testing.T) {
	var driver driver.Driver = &ms.MssqlDriver{}
	q := FromDriver(driver)
	if q == nil {
		t.Fatal("failed to get driver")
	}

	qs := q.Value("It's")
	wanted := "'It''s'"
	if qs != wanted {
		t.Error("quote value failed: got %s wanted %s", qs, wanted)
	}
}
