package singleton

import (
	"fmt"
	"sync"
)

type singleton struct{}

var instance *singleton
var lock = &sync.Mutex{}

func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance")
			instance = &singleton{}
		} else {
			fmt.Println("Instance already created, inner condition")
		}
	} else {
		fmt.Println("Instance already created inside loop, outer condition")
	}

	return instance
}

// When two goroutines reach the first if instance == nil check simultaneously, both will pass the condition and acquire the lock
// The second check inside the lock ensures that only the first goroutine creates the instance, while the second one skips it
