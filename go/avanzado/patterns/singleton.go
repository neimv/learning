package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (d Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton for database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creating done")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating connection to database")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB already created")
	}

	return db
}

func main() {
	fmt.Println("Hello")
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}

	wg.Wait()
}
