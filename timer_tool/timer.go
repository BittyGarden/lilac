package util

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"
)

const blank = "    "

// NewTimer create a timer
func NewTimer(name string) *Timer {
	return &Timer{
		model:     name,
		done:      false,
		root:      newTimerNode("Timer_" + name),
		unUsedTag: make([]string, 0),
	}
}

/*
Timer record duration for code execute since start to end
and format result by Timer.String()

such as:

	timer := NewTimer("Task1")
	timer.Start("codeA")
	time.Sleep(time.Second)
	timer.End("codeA")

	fmt.Println(timer.String())

	OutPut:
	Timer_Task1: 10ms
	    codeA: 10ms

or:

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

	OutPut:
	Timer_Task2: 44ms
	    codeA: 44ms
	        codeA_1: 11ms
	        codeA2: 11ms
*/
type Timer struct {
	mutex     sync.Mutex
	model     string
	done      bool
	root      *TimerNode
	unUsedTag []string
}

func (t *Timer) Start(tag string) {
	if tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.addChild(newTimerNode(tag))
}

func (t *Timer) End(tag string) {
	if tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.root.setEndTime(tag) {
		t.unUsedTag = append(t.unUsedTag, "_"+tag)
	}
}

func (t *Timer) StartSub(parentTag, tag string) {
	if parentTag == "" || tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	newNode := newTimerNode(tag)
	if !t.root.addChildNode(parentTag, newNode) {
		t.unUsedTag = append(t.unUsedTag, parentTag+"_"+tag)
	}
}

func (t *Timer) EndSub(parentTag, tag string) {
	if parentTag == "" || tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.root.setChildEndTime(parentTag, tag) {
		t.unUsedTag = append(t.unUsedTag, parentTag+"_"+tag)
	}
}

func (t *Timer) ReSet() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.done = false
	now := time.Now()
	t.root.start = &now
	t.root.end = nil
	t.root.children = make([]*TimerNode, 0)
	t.unUsedTag = make([]string, 0)
}

func (t *Timer) Exit() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.done {
		now := time.Now()
		t.root.end = &now
	}
}

func (t *Timer) String() string {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.done {
		t.done = true
		now := time.Now()
		t.root.end = &now
	}

	result := t.root.String()
	if len(t.unUsedTag) != 0 {
		result = result + "\nun used tag: " + strings.Join(t.unUsedTag, ",")
	}
	return result
}

type TimerNode struct {
	tag      string       //标签
	start    *time.Time   //开始时间
	end      *time.Time   //结束时间
	children []*TimerNode //子节点
}

func (t *TimerNode) addChild(child *TimerNode) {
	if t.children == nil {
		t.children = make([]*TimerNode, 0)
	}
	t.children = append(t.children, child)
}

func (t *TimerNode) addChildNode(parentTag string, childNode *TimerNode) bool {
	if t.children == nil {
		t.children = make([]*TimerNode, 0)
	}
	for _, child := range t.children {
		if addChildNode(child, parentTag, childNode) {
			return true
		}
	}
	return false
}

func (t *TimerNode) setChildEndTime(parentTag, tag string) bool {
	if t.children == nil {
		t.children = make([]*TimerNode, 0)
	}
	for _, child := range t.children {
		if setChildEndTime(child, parentTag, tag) {
			return true
		}
	}
	return false
}

func addChildNode(parentNode *TimerNode, parentTag string, childNode *TimerNode) bool {
	if parentNode.tag == parentTag {
		parentNode.children = append(parentNode.children, childNode)
		return true
	}

	for _, child := range parentNode.children {
		if addChildNode(child, parentTag, childNode) {
			return true
		}
	}
	return false
}

func setChildEndTime(parentNode *TimerNode, parentTag, tag string) bool {
	if parentNode.tag == parentTag {
		for _, child := range parentNode.children {
			if child.tag == tag {
				now := time.Now()
				child.end = &now
				return true
			}
		}
	}

	for _, child := range parentNode.children {
		if setChildEndTime(child, parentTag, tag) {
			return true
		}
	}
	return false
}

func (t *TimerNode) setEndTime(tag string) bool {
	for _, child := range t.children {
		if child.tag == tag {
			now := time.Now()
			child.end = &now
			return true
		}
	}
	return false
}

func (t *TimerNode) String() string {
	buf := &bytes.Buffer{}
	walkNode(t, buf, 0)
	return buf.String()
}

func walkNode(node *TimerNode, buffer *bytes.Buffer, deep int) {
	blank := getPrefixBlank(deep)
	if node.start == nil || node.end == nil {
		buffer.WriteString(blank)
		buffer.WriteString(node.tag)
		buffer.WriteString(": ")
		buffer.WriteString(" time is nil")
		buffer.WriteString("\n")
	} else {
		buffer.WriteString(blank)
		buffer.WriteString(node.tag)
		buffer.WriteString(": ")
		buffer.WriteString(strconv.Itoa(int(node.end.Sub(*node.start).Milliseconds())))
		buffer.WriteString("ms")
		buffer.WriteString("\n")
	}

	for _, child := range node.children {
		walkNode(child, buffer, deep+1)
	}
}

func newTimerNode(tag string) *TimerNode {
	now := time.Now()
	return &TimerNode{
		tag:      tag,
		start:    &now,
		children: make([]*TimerNode, 0),
	}
}

func getPrefixBlank(deep int) string {
	result := ""
	for i := 0; i < deep; i++ {
		result = result + blank
	}
	return result
}
