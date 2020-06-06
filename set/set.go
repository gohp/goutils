package set

import (
	"fmt"
	"strings"
	"sync"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/5
 */

type Set struct {
	lock *sync.RWMutex
	m    map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		lock: new(sync.RWMutex),
		m:    make(map[interface{}]struct{}),
	}
}

// Add
func (s *Set) Add(val interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.m[val]
	if ok {
		return false
	}

	s.m[val] = struct{}{}
	return true
}

func (s *Set) Equal(ss *Set) bool {
	if s.Len() != ss.Len() {
		return false
	}
	for val := range s.m {
		if !ss.Contains(val) {
			return false
		}
	}
	return true
}

// Del
func (s *Set) Del(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.m, val)
}

// Flush 清空
func (s *Set) Flush() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m = make(map[interface{}]struct{})
}

// Contains
func (s *Set) Contains(val ...interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, i := range val {
		_, found := s.m[i]
		if !found {
			return false
		}
	}
	return true
}

// Len
func (s *Set) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.m)
}

// Union 并集
func (s *Set) Union(set *Set) *Set {
	ret := NewSet()

	s.lock.RLock()
	defer s.lock.RUnlock()
	set.lock.RLock()
	defer set.lock.RUnlock()

	for val := range s.m {
		ret.Add(val)
	}

	for val := range set.m {
		ret.Add(val)
	}
	return ret
}

// Intersect 交集
func (s *Set) Intersect(set *Set) *Set {
	ret := NewSet()

	s.lock.RLock()
	defer s.lock.RUnlock()
	set.lock.RLock()
	defer set.lock.RUnlock()

	for val := range s.m {
		if set.Contains(val) {
			ret.Add(val)
		}
	}
	return ret
}

// Difference 差集
func (s *Set) Difference(set *Set) *Set {
	ret := NewSet()

	s.lock.RLock()
	defer s.lock.RUnlock()
	set.lock.RLock()
	defer set.lock.RUnlock()

	for val := range s.m {
		if !set.Contains(val) {
			ret.Add(val)
		}
	}
	return ret
}

// String
func (s *Set) ToString() string {
	items := make([]string, 0, s.Len())
	s.lock.RLock()
	defer s.lock.RUnlock()

	for val := range s.m {
		items = append(items, fmt.Sprintf("%v", val))
	}
	return fmt.Sprintf("Set{ %s }", strings.Join(items, ", "))
}

// String
func (s *Set) ToSlice() []interface{} {
	items := make([]interface{}, 0, s.Len())
	s.lock.RLock()
	defer s.lock.RUnlock()

	for val := range s.m {
		items = append(items, val)
	}
	return items
}
