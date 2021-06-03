package types

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "math/rand"
    "testing"
)

func TestRangeSet_add(t *testing.T) {
    set := NewRangeSet()

    set.Add(&Range{start: 1, end: 1})
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(1), set.rangeSet[0].end)

    set.Add(&Range{start: 3, end: 3})
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(1), set.rangeSet[0].end)

    assert.Equal(t, uint64(3), set.rangeSet[1].start)
    assert.Equal(t, uint64(3), set.rangeSet[1].end)

    set.Add(&Range{start: 2, end: 2})
    assert.Equal(t, 1, len(set.rangeSet))
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(3), set.rangeSet[0].end)

    set.Add(&Range{start: 5, end: 5})
    assert.Equal(t, 2, len(set.rangeSet))
    assert.Equal(t, uint64(5), set.rangeSet[1].start)
    assert.Equal(t, uint64(5), set.rangeSet[1].end)

    set.Add(&Range{start: 10, end: 10})
    assert.Equal(t, 3, len(set.rangeSet))
    assert.Equal(t, uint64(10), set.rangeSet[2].start)
    assert.Equal(t, uint64(10), set.rangeSet[2].end)

    set.Add(&Range{start: 4, end: 10})
    assert.Equal(t, 1, len(set.rangeSet))
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(10), set.rangeSet[0].end)

    set.Add(&Range{start: 1, end: 1})
    assert.Equal(t, 1, len(set.rangeSet))
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(10), set.rangeSet[0].end)

    set.Add(&Range{start: 10, end: 12})
    assert.Equal(t, 1, len(set.rangeSet))
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(12), set.rangeSet[0].end)

    set.Add(&Range{start: 13, end: 14})
    assert.Equal(t, 1, len(set.rangeSet))
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(14), set.rangeSet[0].end)

    set.Add(&Range{start: 100, end: 120})
    assert.Equal(t, 2, len(set.rangeSet))
    assert.Equal(t, uint64(100), set.rangeSet[1].start)
    assert.Equal(t, uint64(120), set.rangeSet[1].end)

    set.Add(&Range{start: 130, end: 140})
    assert.Equal(t, 3, len(set.rangeSet))
    assert.Equal(t, uint64(130), set.rangeSet[2].start)
    assert.Equal(t, uint64(140), set.rangeSet[2].end)

    set.Add(&Range{start: 122, end: 123})
    assert.Equal(t, 4, len(set.rangeSet))
    assert.Equal(t, uint64(122), set.rangeSet[2].start)
    assert.Equal(t, uint64(123), set.rangeSet[2].end)

    set.Add(&Range{start: 15, end: 99})
    set.Add(&Range{start: 120, end: 140})
    assert.Equal(t, 1, len(set.rangeSet))
    assert.Equal(t, uint64(1), set.rangeSet[0].start)
    assert.Equal(t, uint64(140), set.rangeSet[0].end)
}

func TestA(t *testing.T) {
    set := NewRangeSet()
    set.Add(NewRange(7208, 7208))
    set.Add(NewRange(5448, 5448))
    set.Add(NewRange(7762, 7762))
    set.Add(NewRange(4610, 4610))

    assert.Equal(t, uint64(4610), set.rangeSet[0].start)
}

func TestRangeSet(t *testing.T) {
    set := NewRangeSet()
    for {
        value := rand.Intn(10000)
        set.Add(NewRange(uint64(value), uint64(value)))
        var r *Range
        for _, v := range set.rangeSet {
            if r == nil {
                r = v
                continue
            }
            b := r.end < v.start
            if !b {
                fmt.Println(set.String())
            }
            assert.True(t, b)
            r = v
        }
        if set.rangeSet[0].start == uint64(0) && set.rangeSet[0].end == uint64(9999) {
            fmt.Println(set.String())
            break
        }
    }
}

func TestHaveSameIndex(t *testing.T) {
    a := &Range{start: 1, end: 1}
    b := &Range{start: 1, end: 10}
    assert.True(t, haveSameIndex(a, b))
}

func TestRemoveBefore(t *testing.T) {
    set := NewRangeSet()
    set.Add(NewRange(1, 1))
    set.Add(NewRange(3, 3))
    set.Add(NewRange(5, 5))
    set.Add(NewRange(7, 7))

    set.RemoveBefore(0)
    assert.Equal(t, 4, len(set.rangeSet))

    set.RemoveBefore(2)
    assert.Equal(t, 3, len(set.rangeSet))

    set.RemoveBefore(7)
    assert.Equal(t, 1, len(set.rangeSet))

    set.RemoveBefore(8)
    assert.Equal(t, 0, len(set.rangeSet))
}

func TestRemoveBeforeEmpty(t *testing.T) {
    set := NewRangeSet()

    set.RemoveBefore(0)
    assert.Equal(t, 0, len(set.rangeSet))
}

func TestContains(t *testing.T) {
    set := NewRangeSet()

    set.Add(NewRange(1, 1))

    set.Add(NewRange(3, 3))

    set.Add(NewRange(9, 100))
    set.Add(NewRange(100, 200))

    set.Add(NewRange(10000, 99999))

    assert.False(t, set.Contains(0))
    assert.True(t, set.Contains(1))
    assert.False(t, set.Contains(2))
    assert.True(t, set.Contains(3))
    assert.False(t, set.Contains(4))

    assert.False(t, set.Contains(7))
    assert.False(t, set.Contains(8))
    assert.True(t, set.Contains(9))
    assert.True(t, set.Contains(10))
    assert.True(t, set.Contains(99))
    assert.True(t, set.Contains(100))
    assert.True(t, set.Contains(199))
    assert.True(t, set.Contains(200))
    assert.False(t, set.Contains(201))
    assert.False(t, set.Contains(202))

    assert.False(t, set.Contains(8888))
    assert.False(t, set.Contains(9999))
    assert.True(t, set.Contains(10000))
    assert.True(t, set.Contains(88888))
    assert.True(t, set.Contains(99999))

    assert.False(t, set.Contains(100000))
    assert.False(t, set.Contains(100001))
    assert.False(t, set.Contains(500001))
}
