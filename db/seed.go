package db

// SeedData will add some data to the database
func SeedData(db *DB) error {
	return db.WithTX(func(tx *TX) error {
		var i int64
		for i = 0; i < 10; i++ {
			if _, err := tx.InsertTimestamp(i * 100000000); err != nil {
				return err
			}
		}
		return nil
	})
}
