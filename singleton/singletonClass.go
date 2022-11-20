package singleton

import(
	"fmt"
	"sync"
)

lock := &sync.Mutex{}

type instance struct {
}

var classInstance *instance

func getSingleInstance() *instance {
	lock.Lock()
	defer lock.Unlock()
	if classInstance==nil {
		classInstance = &instance{}
	}
	return classInstance
}
