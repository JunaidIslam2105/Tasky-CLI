package db

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"time"
)

var taskBucket = []byte("tasks") // Name of the bucket that stores the current tasks

var db *bolt.DB // Pointer to the bolt.DB

type Task struct {
	// Struct that defines a Task
	Value string
	Key   int
}

func Init(dbPath string) error {
	var err error
	//Creates a new database at 'homepath/tasks.db' and stores the pointer in 'db'
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		//Creates a new bucket with bucketname 'taskBucket'
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		//Uses the inbuilt NextSequence() to get next id and store the task in the DB
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task)) // Puts the key along with the corresponding task in the bucket
	})

	if err != nil {
		return -1, err
	}

	return 0, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			//Iterates over all the tasks and stores in 'tasks'
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(Key int) error {
	// Deletes a task from the database with a particular key
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		b.Delete(itob(Key))

		return nil
	})
}

func itob(v int) []byte {
	//Helper function to convert integer to byte slice
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v)) // BigEndian is used as a convention
	return b
}

func btoi(b []byte) int {

	// Helper function to convert byte slice to integer
	return int(binary.BigEndian.Uint64(b))
}
