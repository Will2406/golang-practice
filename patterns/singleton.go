package main

import (
	"fmt"
	"sync"
)

type Database struct {
}

func (database Database) createSingletonConection() {
	fmt.Println("Creacion de la base de datos")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("La base de datos esta siendo creado")
		db = &Database{}
		db.createSingletonConection()
	} else {
		fmt.Println("Ya existe una instancia de la base de datos")
	}
	return db
}

/*func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}*/
