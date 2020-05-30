package safemap

import "sync"

type SafeMap struct {
	lock *sync.RWMutex
	m    map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		m:    make(map[interface{}]interface{}),
	}
}

func (s *SafeMap) Get(key interface{}) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if val, ok := s.m[key]; ok {
		return val
	}
	return nil
}

func (s *SafeMap) Set(key interface{}, value interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if val, ok := s.m[key]; !ok {
		s.m[key] = value
	} else if val != value {
		s.m[key] = value
	} else {
		return false
	}
	return true
}

func (s *SafeMap) Delete(key interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.m, key)
}

func (s *SafeMap) Exist(key interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.m[key]
	return ok
}

func (s *SafeMap) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.m)
}

func (s *SafeMap) Items() map[interface{}]interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	r := make(map[interface{}]interface{})
	for k, v := range s.m {
		r[k] = v
	}
	return r
}

func (s *SafeMap) Keys() []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	r := make([]interface{}, 0)
	for k, _ := range s.m {
		r = append(r, k)
	}
	return r
}

func (s *SafeMap) Values() []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	r := make([]interface{}, 0)
	for _, v := range s.m {
		r = append(r, v)
	}
	return r
}
