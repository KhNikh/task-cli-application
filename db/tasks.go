package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("task")
var db *bolt.DB

type Task struct {
	key   int
	value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()

		// Put method takes bytes as input so we have to convert id64 to bytes
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	} else {
		return id, nil
	}
}

func itob(id int) []byte {
	//use binary package to convert int to binary
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return b

}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
