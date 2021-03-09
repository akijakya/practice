package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

/*
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
*/

// this is typecastinh from string to byte slice
var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key int
	Value string
}

// this will not run automatically on program start because starts with a capital I instead of a lowercase i
func Init(dbPath string) error {
	// the error needs to be declared separately, because in the next line, if the err is not something that
	// is already declared, the 'db' variable will be a local variable as well, instead of using the previously
	// declared 'db' variable outside the scope of this function
	var err error
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask to create task
func CreateTask (task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// integer to byte slice converter
func itob (v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// byte slice to integer converter
func btoi (b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}