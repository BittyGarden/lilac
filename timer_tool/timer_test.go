package util

import (
    "fmt"
    "math/rand"
    "strconv"
    "sync"
    "testing"
    "time"
)

func TestExampleA(t *testing.T) {
    timer := NewTimer("Task1")
    timer.Start("codeA")
    time.Sleep(time.Millisecond * 10)
    timer.End("codeA")
    fmt.Println(timer.String())
}

func TestExampleB(t *testing.T) {
    timer := NewTimer("Task2")
    timer.Start("codeA")
    time.Sleep(time.Millisecond * 10)
    timer.StartSub("codeA", "codeA_1")
    time.Sleep(time.Millisecond * 10)
    timer.EndSub("codeA", "codeA_1")
    timer.StartSub("codeA", "codeA2")
    time.Sleep(time.Millisecond * 10)
    timer.EndSub("codeA", "codeA2")
    time.Sleep(time.Millisecond * 10)
    timer.End("codeA")

    fmt.Println(timer.String())
}

func TestUnuUsed(t *testing.T) {
    timer := NewTimer("Task2")
    timer.Start("codeA")
    time.Sleep(time.Millisecond * 10)
    timer.StartSub("codeA", "codeA_1")
    time.Sleep(time.Millisecond * 10)
    timer.EndSub("codeA", "codeA_1")
    timer.StartSub("codeA", "codeA2")
    time.Sleep(time.Millisecond * 10)
    timer.EndSub("codeA", "codeA2")
    time.Sleep(time.Millisecond * 10)
    timer.End("codeA")

    timer.StartSub("aa", "ss")
    timer.End("cc")

    fmt.Println(timer.String())
}

func TestTimer(t *testing.T) {
    timer := NewTimer("A")
    timer.Start("A1")

    waitGroup := sync.WaitGroup{}
    for i := 0; i < 3; i++ {
        waitGroup.Add(1)
        go func(index int) {
            timer.StartSub("A1", "A1"+strconv.Itoa(index))
            time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
            timer.EndSub("A1", "A1"+strconv.Itoa(index))
            waitGroup.Done()
        }(i)
    }
    waitGroup.Wait()

    timer.StartSub("A1", "A2")
    for i := 0; i < 3; i++ {
        waitGroup.Add(1)
        go func(index int) {
            timer.StartSub("A2", "A2"+strconv.Itoa(index))
            time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
            timer.EndSub("A2", "A2"+strconv.Itoa(index))
            waitGroup.Done()
        }(i)
    }
    waitGroup.Wait()
    timer.EndSub("A1", "A2")

    timer.StartSub("A2", "A3")
    for i := 0; i < 3; i++ {
        waitGroup.Add(1)
        go func(index int) {
            timer.StartSub("A3", "A3"+strconv.Itoa(index))
            time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
            timer.EndSub("A3", "A3"+strconv.Itoa(index))
            waitGroup.Done()
        }(i)
    }
    waitGroup.Wait()

    timer.EndSub("A2", "A3")

    for i := 0; i < 3; i++ {
        waitGroup.Add(1)
        go func(index int) {
            timer.StartSub("A20", "A20"+strconv.Itoa(index))
            time.Sleep(time.Millisecond * time.Duration(rand.Int()%300))
            timer.EndSub("A20", "A20"+strconv.Itoa(index))
            waitGroup.Done()
        }(i)
    }
    waitGroup.Wait()

    timer.End("A1")
    timer.Exit()

    fmt.Println(timer.String())
}
