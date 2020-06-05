package set

import "sync"

/**
* @Author: Jam Wong
* @Date: 2020/6/5
 */

type Set struct {
	lock *sync.RWMutex
	m map[interface{}]bool
}

func New() *Set {
	return &Set{
		lock: new(sync.RWMutex),
		m: make(map[interface{}]bool),
	}
}

func (s *Set) Add(val interface{})  {
	panic("1")
}

func (s *Set) Del(val interface{})  {
	panic("1")
}

