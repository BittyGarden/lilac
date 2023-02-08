package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	model := "modelA"
	GetTimer(model).Start("A1")

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			GetTimer(model).StartSub("A1", "A1"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			GetTimer(model).EndSub("A1", "A1"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	GetTimer(model).StartSub("A1", "A2")
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			GetTimer(model).StartSub("A2", "A2"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			GetTimer(model).EndSub("A2", "A2"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	GetTimer(model).EndSub("A1", "A2")

	GetTimer(model).StartSub("A2", "A3")
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			GetTimer(model).StartSub("A3", "A3"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			GetTimer(model).EndSub("A3", "A3"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	GetTimer(model).EndSub("A2", "A3")

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			GetTimer(model).StartSub("A20", "A20"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			GetTimer(model).EndSub("A20", "A20"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	GetTimer(model).End("A1")
	GetTimer(model).Exit()

	fmt.Println(GetTimer(model).String())
}
