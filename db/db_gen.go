package db

import "github.com/willowtreeapps/rootx"

func (db *DB) AllTimestamps(instances interface{}) error {
	return rootx.SelectAll(db, "all.sql", instances)
}

func (tx *TX) AllTimestamps(instances interface{}) error {
	return rootx.SelectAll(tx, "all.sql", instances)
}

func (tx *TX) InsertTimestamp(seconds int64) (int64, error) {
	return rootx.Insert(tx, "insert.sql", seconds)
}
