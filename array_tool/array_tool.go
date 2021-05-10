package array_tool

import "errors"

// SplitArray split array to some segment by split size
func SplitArray(splitSize int, array []int) [][]int {
    if splitSize < 0 {
        panic(errors.New("split size should grater than zero"))
    }

    arraySize := len(array)

    if array == nil || arraySize == 0 {
        return make([][]int, 0)
    }

    size := arraySize/splitSize + 1

    if splitSize > arraySize {
        size = 1
    }

    if arraySize%splitSize == 0 {
        size = arraySize / splitSize
    }

    result := make([][]int, size)
    for i := 0; i < size; i++ {
        start := i * splitSize
        end := start + splitSize
        if end > arraySize {
            end = arraySize
        }
        temp := make([]int, end-start)
        copy(temp, array[start:end])
        result[i] = temp
    }
    return result
}
