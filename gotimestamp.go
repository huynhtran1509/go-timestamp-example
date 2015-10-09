package timestamp

import (
	"encoding/json"

	"github.com/willowtreeapps/timestamp/db"
)

// Timestamp holds seconds since epoch
type Timestamp struct {
	SecondsSinceEpoch int64 `db:"seconds_since_epoch" json:"seconds_since_epoch"`
}

type dbWrapper struct {
	*db.DB
}

var globalDB *dbWrapper

// SetUp initializes the database
func SetUp(path string) error {
	innerdb, err := db.NewDB(path)
	if err != nil {
		return err
	}
	globalDB = &dbWrapper{innerdb}
	return nil
}

// Bootstrap sets up the new database
func Bootstrap() error {
	if err := globalDB.SetUp(); err != nil {
		return err
	}
	return globalDB.SeedData()
}

// AllJSON returns all the timestamps in the database as JSON
func AllJSON() ([]byte, error) {
	var timestamps []Timestamp
	if err := globalDB.AllTimestamps(&timestamps); err != nil {
		return nil, err
	}
	return json.Marshal(timestamps)
}

// Insert creates a timestamp
func Insert(seconds int64) error {
	return globalDB.WithTX(func(tx *db.TX) error {
		_, err := tx.InsertTimestamp(seconds)
		return err
	})
}
