package main

import (
	"time"
	"github.com/boltdb/bolt"
)

// boltdb is a simple key-value store completely written in Go so you can just go get it,
// no further installation is needed, it will store the data in a file
// it can only store byte slices in so called buckets though so we need to do some converting sometimes
// you can also nest buckets into each other so like a user_id key has a bucket as value where the users email, name etc. is stored
func main() {
	// there is a timeout because only one process can write in the database at a time,
	// so it is not good to leave it open forever if there is an error
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
}