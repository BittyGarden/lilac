package random_tool

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestRandomNumber(t *testing.T) {
    randomNumber, err := RandomNumber(0, 0)
    assert.Nil(t, randomNumber)
    assert.NotNil(t, err)

    randomNumber, err = RandomNumber(1, 2)
    assert.Nil(t, randomNumber)
    assert.NotNil(t, err)

    randomNumber, err = RandomNumber(1, 1)
    assert.Nil(t, err)
    assert.Equal(t, 1, len(randomNumber))
    assert.Equal(t, 0, randomNumber[0])

    randomNumber, err = RandomNumber(100, 1)
    assert.Nil(t, err)
    assert.Equal(t, 1, len(randomNumber))

    randomNumber, err = RandomNumber(100, 99)
    assert.Nil(t, err)
    assert.Equal(t, 99, len(randomNumber))

    randomNumber, err = RandomNumber(100, 100)
    assert.Nil(t, err)
    assert.Equal(t, 100, len(randomNumber))

    numberSet := make(map[int]struct{})
    for _, v := range randomNumber {
        _, ok := numberSet[v]
        assert.False(t, ok)
        numberSet[v] = struct{}{}
    }
}
