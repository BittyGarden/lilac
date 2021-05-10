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

    splitArray = SplitArray(2, data)
    assert.Equal(t, 50, len(splitArray))

    splitArray = SplitArray(3, data)
    assert.Equal(t, 34, len(splitArray))

    splitArray = SplitArray(99, data)
    assert.Equal(t, 2, len(splitArray))

    splitArray = SplitArray(100, data)
    assert.Equal(t, 1, len(splitArray))

    splitArray = SplitArray(101, data)
    assert.Equal(t, 1, len(splitArray))
}
