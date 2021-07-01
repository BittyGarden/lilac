package set_tool

import (
    "fmt"
    "strings"
)

//IntervalSet is an interval set, such as {[1,1],[5,100],[108,203],[300,400]}
type IntervalSet struct {
    intervalSet []*Interval
}

//Interval between [start,end]
type Interval struct {
    start uint64
    end   uint64
}

//NewIntervalSet create an interval set
func NewIntervalSet() *IntervalSet {
    return &IntervalSet{
        intervalSet: make([]*Interval, 0),
    }
}

//NewIntervalSet create an interval set
func NewInterval(start, end uint64) *Interval {
    return &Interval{start: start, end: end}
}

func (rs *IntervalSet) Add(r *Interval) {
    if r == nil {
        return
    }

    if r.start > r.end {
        return
    }

    if r.start < 0 || r.end < 0 {
        return
    }

    if len(rs.intervalSet) == 0 {
        rs.intervalSet = append(rs.intervalSet, &Interval{start: r.start, end: r.end})
        return
    }

    insertIndex := -1
    for index, rr := range rs.intervalSet {
        if r.start < rr.start && insertIndex == -1 {
            insertIndex = index
        }
        if haveSameIndex(r, rr) {
            rr.start = Min(r.start, rr.start)
            rr.end = Max(r.end, rr.end)
            nextIndex := index + 1
            if nextIndex < len(rs.intervalSet) && rr.end+1 >= rs.intervalSet[nextIndex].start {
                temp := rs.intervalSet[nextIndex]
                if index+2 < len(rs.intervalSet) {
                    rs.intervalSet = append(rs.intervalSet[0:nextIndex], rs.intervalSet[index+2:]...)
                } else {
                    rs.intervalSet = rs.intervalSet[0:nextIndex]
                }
                rs.Add(temp)
            }
            return
        }
    }

    if insertIndex == -1 {
        insertIndex = len(rs.intervalSet)
    }

    insert := &Interval{
        start: r.start,
        end:   r.end,
    }
    rs.intervalSet = append(rs.intervalSet, insert)
    copy(rs.intervalSet[insertIndex+1:], rs.intervalSet[insertIndex:])
    rs.intervalSet[insertIndex] = insert
}

//RemoveBefore remove interval before index
func (rs *IntervalSet) RemoveBefore(removeIntervalEnd uint64) {
    removeIndex := -1
    for index, v := range rs.intervalSet {
        if v.end >= removeIntervalEnd {
            break
        }
        removeIndex = index
    }
    if removeIndex == -1 {
        return
    }
    removeIndex = removeIndex + 1
    if removeIndex >= len(rs.intervalSet) {
        rs.intervalSet = make([]*Interval, 0)
        return
    }
    rs.intervalSet = rs.intervalSet[removeIndex:]
}

func (rs *IntervalSet) String() string {
    builder := strings.Builder{}
    for _, r := range rs.intervalSet {
        builder.WriteString("[")
        builder.WriteString(fmt.Sprintf("%d", r.start))
        builder.WriteString(string(","))
        builder.WriteString(fmt.Sprintf("%d", r.end))
        builder.WriteString("] ")
    }
    return builder.String()
}

func (rs *IntervalSet) MissingCount() int {
    if len(rs.intervalSet) == 0 {
        return 0
    }
    var start uint64 = 1
    missingCount := 0
    for _, r := range rs.intervalSet {
        for {
            if start < r.start {
                start++
                missingCount++
            }
            if start == r.start {
                start = r.end + 1
                break
            }
        }
    }
    return missingCount
}

func (rs *IntervalSet) Contains(index uint64) bool {
    for _, r := range rs.intervalSet {
        if index >= r.start && index <= r.end {
            return true
        }
    }
    return false
}

func (rs *IntervalSet) index(index uint64) int {
    for i, r := range rs.intervalSet {
        if index >= r.start && index <= r.end {
            return i
        }
    }
    return -1
}

func haveSameIndex(a *Interval, b *Interval) bool {
    return (a.start+1 >= b.start && a.start <= b.end+1) || (a.end+1 >= b.start && a.end <= b.end+1)
}

func Min(a, b uint64) uint64 {
    if a < b {
        return a
    }
    return b
}

func Max(a, b uint64) uint64 {
    if a > b {
        return a
    }
    return b
}
