package utils

import (
    "sync"
    "time"
)

type cacheNode struct {
    createTime time.Time
    value      interface{} //cache value
    next       *cacheNode
    prev       *cacheNode
}

//FixedDurationCache cache data with fixed duration from [now - duration,now]
type FixedDurationCache struct {
    duration time.Duration //data cache time
    mutex    sync.Mutex
    start    *cacheNode
    end      *cacheNode
    keyFunc  func(v interface{}) string // get key from value
}

func NewFixedDurationCache(duration time.Duration, keyFunc func(v interface{}) string) *FixedDurationCache {
    return &FixedDurationCache{
        duration: duration,
        keyFunc:  keyFunc,
    }
}

//Add add data to cache
func (fdc *FixedDurationCache) Add(v interface{}) {
    fdc.mutex.Lock()
    defer fdc.mutex.Unlock()

    node := newCacheNode(v)

    //new node
    if fdc.end == nil {
        fdc.start = node
        fdc.end = node
        return
    }

    //add to end
    node.prev = fdc.end
    fdc.end.next = node
    fdc.end = node

    //after add data to cache we can assert that last data is not expire
    node = fdc.start
    for {
        if fdc.exipreTimeLine().Before(node.createTime) {
            fdc.start = node
            prev := node.prev
            //cut link list from node
            if prev != nil {
                node.prev.next = nil
                node.prev = nil
            }
            break
        }
        node = node.next
    }
}

//Exists if in cache return ture, else return false
func (fdc *FixedDurationCache) ExistsByKey(key string) bool {
    for node := fdc.end; node != nil; node = node.prev {
        if node.createTime.Before(fdc.exipreTimeLine()) {
            //not search from expire cache data
            return false
        }
        if key == fdc.keyFunc(node.value) {
            return true
        }
    }
    return false
}

//Exists if v in cache return ture, else return false
func (fdc *FixedDurationCache) Exists(v interface{}) bool {
    return fdc.ExistsByKey(fdc.keyFunc(v))
}

//GetByKey return cache data mapped by key
func (fdc *FixedDurationCache) GetByKey(key string) interface{} {
    for node := fdc.end; node != nil; node = node.prev {
        if node.createTime.Before(fdc.exipreTimeLine()) {
            //not search from expire cache data
            return nil
        }
        if key == fdc.keyFunc(node.value) {
            return node.value
        }
    }
    return nil
}

//Diff return cache data who's key not exists in key list
func (fdc *FixedDurationCache) Diff(keyList []string) []interface{} {
    if keyList == nil || len(keyList) == 0 {
        return fdc.GetAll()
    }

    keyMap := make(map[string]struct{})
    for _, key := range keyList {
        keyMap[key] = struct{}{}
    }

    result := make([]interface{}, 0)
    for node := fdc.end; node != nil; node = node.prev {
        if node.createTime.Before(fdc.exipreTimeLine()) {
            break
        }
        nodeKey := fdc.keyFunc(node.value)
        if _, ok := keyMap[nodeKey]; !ok {
            result = append(result, node.value)
        }
    }
    return result
}

//GetAll get all value in cache
func (fdc *FixedDurationCache) GetAll() []interface{} {
    result := make([]interface{}, 0)
    for node := fdc.end; node != nil; node = node.prev {
        if node.createTime.Before(fdc.exipreTimeLine()) {
            break
        }
        result = append(result, node.value)
    }
    return result
}

//expireTimeLine return time line, before the line data has expire, after the line data not expire
func (fdc *FixedDurationCache) exipreTimeLine() time.Time {
    return time.Now().Add(fdc.duration * -1)
}

func newCacheNode(v interface{}) *cacheNode {
    return &cacheNode{
        createTime: time.Now(),
        value:      v,
        next:       nil,
        prev:       nil,
    }
}
