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
	stopWatcher := NewStopWatcher("A")
	stopWatcher.Start("A1")

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			stopWatcher.StartSub("A1", "A1"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			stopWatcher.EndSub("A1", "A1"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	stopWatcher.StartSub("A1", "A2")
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			stopWatcher.StartSub("A2", "A2"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			stopWatcher.EndSub("A2", "A2"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	stopWatcher.EndSub("A1", "A2")

	stopWatcher.StartSub("A2", "A3")
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			stopWatcher.StartSub("A3", "A3"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			stopWatcher.EndSub("A3", "A3"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	stopWatcher.EndSub("A2", "A3")

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func(index int) {
			stopWatcher.StartSub("A20", "A20"+strconv.Itoa(index))
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
			stopWatcher.EndSub("A20", "A20"+strconv.Itoa(index))
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

	stopWatcher.End("A1")
	stopWatcher.Exit()

	fmt.Println(stopWatcher.String())
}
