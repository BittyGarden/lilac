package array_tool

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestSplitArray(t *testing.T) {

    data := make([]int, 100)

    for i := 0; i < 100; i++ {
        data[i] = i
    }

    splitArray := SplitArray(1, data)
    assert.Equal(t, 100, len(splitArray))
    assert.Equal(t, 1, len(splitArray[0]))
    assert.Equal(t, 1, len(splitArray[99]))

    splitArray = SplitArray(2, data)
    assert.Equal(t, 50, len(splitArray))
    assert.Equal(t, 2, len(splitArray[0]))
    assert.Equal(t, 2, len(splitArray[49]))

    splitArray = SplitArray(3, data)
    assert.Equal(t, 34, len(splitArray))
    assert.Equal(t, 3, len(splitArray[0]))
    assert.Equal(t, 1, len(splitArray[33]))

    splitArray = SplitArray(99, data)
    assert.Equal(t, 2, len(splitArray))
    assert.Equal(t, 99, len(splitArray[0]))
    assert.Equal(t, 1, len(splitArray[1]))

    splitArray = SplitArray(100, data)
    assert.Equal(t, 1, len(splitArray))
    assert.Equal(t, 100, len(splitArray[0]))

    splitArray = SplitArray(101, data)
    assert.Equal(t, 1, len(splitArray))
    assert.Equal(t, 100, len(splitArray[0]))
}
