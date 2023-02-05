package main

import (
	"fmt"
	"sync"
)

type Resource struct {
	ID int
}

type ResourcePool struct {
	mu        sync.Mutex
	resources chan *Resource
}

func NewResourcePool(numResources int) *ResourcePool {
	rp := &ResourcePool{
		resources: make(chan *Resource, numResources),
	}
	for i := 0; i < numResources; i++ {
		rp.resources <- &Resource{ID: i}
	}
	return rp
}

func (rp *ResourcePool) Get() *Resource {
	return <-rp.resources
}

func (rp *ResourcePool) Put(r *Resource) {
	rp.resources <- r
}

func main() {
	rp := NewResourcePool(5)
	res := rp.Get()
	fmt.Println("Got resource:", res.ID)
	rp.Put(res)
	fmt.Println("Returned resource:", res.ID)
}

/*
go run resourcepoolwithchannel.go
Got resource: 0
Returned resource: 0



 */
