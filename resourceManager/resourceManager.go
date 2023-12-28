package main

import (
	"fmt"
	"sync"
)

type ResourceManager struct {
	IsResourceAvailable bool
}

var instance *ResourceManager
var once sync.Once

func GetResourceManagerInstance() *ResourceManager {
	once.Do(func() {
		instance = &ResourceManager{IsResourceAvailable: true}

	})

	return instance
}

// AcquireResource acuires the shared resource
func (rm *ResourceManager) AcquireResource() bool {
	if rm.IsResourceAvailable {
		rm.IsResourceAvailable = false
		return true
	}

	return false
}

func (rm *ResourceManager) ReleaseResource() {
	rm.IsResourceAvailable = true
}

func main() {
	resourceMgr := GetResourceManagerInstance()
	success := resourceMgr.AcquireResource()
	if success {
		fmt.Println("Resource acquired!")
		resourceMgr.ReleaseResource()
		fmt.Println("Resource released!")
	} else {
		fmt.Println("Failed to acquire resource")
	}
}
