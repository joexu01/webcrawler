package engine

//import (
//	"sync"
//)
//
//type JsonStore struct {
//	source   map[string]string
//	JsonPath string
//	once     sync.Once
//	sync.RWMutex
//}
//
//func (t *JsonStore) Set(key string, g interface{}) bool {
//	t.Lock()
//	defer t.Unlock()
//	t.source[key] = g
//	return true
//}
//
//func (t *JsonStore) Get(id string) interface{} {
//	t.RLock()
//	defer t.RUnlock()
//
//	v, ok := t.source[id]
//	if !ok {
//		return nil
//	}
//	return v
//}
