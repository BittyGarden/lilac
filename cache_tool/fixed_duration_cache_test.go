package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

type ptest struct {
	id   string
	name string
	age  int
}

func TestNewFixDurationCache(t *testing.T) {
	cache := NewFixedDurationCache(time.Second*100, func(v interface{}) string {
		return v.(ptest).id
	})

	assert.NotNil(t, cache)
}

func TestExistsAndExpire(t *testing.T) {
	duration := time.Millisecond * 100
	cache := NewFixedDurationCache(duration, func(v interface{}) string {
		return v.(ptest).id
	})
	assert.NotNil(t, cache)

	v1 := ptest{id: "1", name: "1", age: 1}
	cache.Add(v1)
	v2 := ptest{id: "2", name: "2", age: 2}
	cache.Add(v2)

	assert.True(t, cache.Exists(v1))
	assert.True(t, cache.Exists(v2))
	assert.True(t, cache.ExistsByKey("1"))
	assert.True(t, cache.ExistsByKey("2"))

	all := cache.GetAll()
	assert.Equal(t, 2, len(all))

	time.Sleep(duration)
	assert.False(t, cache.Exists(v1))
	assert.False(t, cache.Exists(v2))
	assert.False(t, cache.ExistsByKey("1"))
	assert.False(t, cache.ExistsByKey("2"))

	v3 := ptest{id: "3", name: "3", age: 3}
	cache.Add(v3)

	assert.Equal(t, "3", cache.start.value.(ptest).id)
}

func TestDiff(t *testing.T) {
	duration := time.Millisecond * 100
	cache := NewFixedDurationCache(duration, func(v interface{}) string {
		return v.(ptest).id
	})
	assert.NotNil(t, cache)

	v1 := ptest{id: "1", name: "1", age: 1}
	cache.Add(v1)
	v2 := ptest{id: "2", name: "2", age: 2}
	cache.Add(v2)
	v3 := ptest{id: "3", name: "3", age: 3}
	cache.Add(v3)
	v4 := ptest{id: "4", name: "4", age: 4}
	cache.Add(v4)

	all := cache.GetAll()
	assert.Equal(t, 4, len(all))

	diff := cache.Diff([]string{"2", "4"})
	assert.Equal(t, 2, len(diff))
}

func TestParallel(t *testing.T) {
	goroutine := 10
	addCount := 100000

	duration := time.Millisecond * 1000

	cache := NewFixedDurationCache(duration, func(v interface{}) string {
		return v.(ptest).id
	})

	start := time.Now().UnixNano()
	waitGroup := sync.WaitGroup{}
	for i := 0; i < goroutine; i++ {
		waitGroup.Add(1)
		go func(index int) {
			for c := 0; c < addCount; c++ {
				cache.Add(ptest{id: getId(index, c), name: "_", age: index})
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000000)

	start = time.Now().UnixNano()
	p := ptest{id: getId(goroutine-1, addCount-1), name: "_", age: 0}
	end = time.Now().UnixNano()
	fmt.Println((end - start) / 1000000)
	assert.True(t, cache.Exists(p))

	all := cache.GetAll()
	assert.LessOrEqual(t, 0, len(all))

	time.Sleep(duration)

	all = cache.GetAll()
	assert.Equal(t, 0, len(all))
}

func getId(prefix int, index int) string {
	return fmt.Sprintf("%d_%d", prefix, index)
}
