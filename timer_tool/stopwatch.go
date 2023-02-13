package util

import (
	"bytes"
	"strconv"
	"sync"
	"time"
)

const fourBlank = "    "

func NewStopWatcher(name string) *StopWatcher {
	return &StopWatcher{
		model: name,
		done:  false,
		root:  newTimerNode("StopWatcher_" + name),
	}
}

type StopWatcher struct {
	mutex sync.Mutex
	model string
	done  bool
	root  *TimerNode
}

func (t *StopWatcher) Start(tag string) {
	if tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.addChild(newTimerNode(tag))
}

func (t *StopWatcher) End(tag string) {
	if tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.setEndTime(tag)
}

func (t *StopWatcher) Exit() {
	if !t.done {
		now := time.Now()
		t.root.end = &now
	}
}

func (t *StopWatcher) String() string {
	if !t.done {
		t.done = true
		now := time.Now()
		t.root.end = &now
	}
	return t.root.String()
}

func (t *StopWatcher) StartSub(parentTag, tag string) {
	if parentTag == "" || tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.addChildNode(parentTag, newTimerNode(tag))

}

func (t *StopWatcher) EndSub(parentTag, tag string) {
	if parentTag == "" || tag == "" {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.setChildEndTime(parentTag, tag)
}

type TimerNode struct {
	tag      string     //标签
	start    *time.Time //开始时间
	end      *time.Time //结束时间
	children []*TimerNode
}

func (t *TimerNode) addChild(child *TimerNode) {
	if t.children == nil {
		t.children = make([]*TimerNode, 0)
	}
	t.children = append(t.children, child)
}

func (t *TimerNode) addChildNode(parentTag string, childNode *TimerNode) {
	if t.children == nil {
		t.children = make([]*TimerNode, 0)
	}
	for _, child := range t.children {
		if addChildNode(child, parentTag, childNode) {
			return
		}
	}
	panic("addChildNode " + parentTag + "-" + childNode.tag)
}

func (t *TimerNode) setChildEndTime(parentTag, tag string) {
	if t.children == nil {
		t.children = make([]*TimerNode, 0)
	}
	for _, child := range t.children {
		if setChildEndTime(child, parentTag, tag) {
			return
		}
	}
	panic("setChildEndTime")
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

func (t *TimerNode) setEndTime(tag string) {
	for _, child := range t.children {
		if child.tag == tag {
			now := time.Now()
			child.end = &now
		}
	}
}

func (t *TimerNode) String() string {
	buf := &bytes.Buffer{}
	WalkNode(t, buf, 0)
	return buf.String()
}

func WalkNode(node *TimerNode, buffer *bytes.Buffer, deep int) {
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
		WalkNode(child, buffer, deep+1)
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
		result = result + fourBlank
	}
	return result
}
