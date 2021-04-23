package random_tool

import (
    "errors"
    "math/rand"
)

//RandomNumber generate specified count random number that in range [0, max),
func RandomNumber(max int, count int) ([]int, error) {
    if max <= 0 || count <= 0 {
        return nil, errors.New("random number max or count should greater than 0")
    }
    if max < count {
        return nil, errors.New("random number condition: max >= count")
    }
    return rand.Perm(max)[0:count], nil
}
