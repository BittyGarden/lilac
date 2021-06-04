package types

import (
    "fmt"
    "strings"
)

//RangeSet is a range set, such as {[1,1],[5,100],[108,203],[300,400]}
type RangeSet struct {
    rangeSet []*Range
}

//Range between [start,end]
type Range struct {
    start uint64
    end   uint64
}

//NewRangeSet create a range set
func NewRangeSet() *RangeSet {
    return &RangeSet{
        rangeSet: make([]*Range, 0),
    }
}

//NewRangeSet create a range set
func NewRange(start, end uint64) *Range {
    return &Range{start: start, end: end}
}

func (rs *RangeSet) Add(r *Range) {
    if r == nil {
        return
    }

    if r.start > r.end {
        return
    }

    if r.start < 0 || r.end < 0 {
        return
    }

    if len(rs.rangeSet) == 0 {
        rs.rangeSet = append(rs.rangeSet, &Range{start: r.start, end: r.end})
        return
    }

    insertIndex := -1
    for index, rr := range rs.rangeSet {
        if r.start < rr.start && insertIndex == -1 {
            insertIndex = index
        }
        if haveSameIndex(r, rr) {
            rr.start = Min(r.start, rr.start)
            rr.end = Max(r.end, rr.end)
            nextIndex := index + 1
            if nextIndex < len(rs.rangeSet) && rr.end+1 >= rs.rangeSet[nextIndex].start {
                temp := rs.rangeSet[nextIndex]
                if index+2 < len(rs.rangeSet) {
                    rs.rangeSet = append(rs.rangeSet[0:nextIndex], rs.rangeSet[index+2:]...)
                } else {
                    rs.rangeSet = rs.rangeSet[0:nextIndex]
                }
                rs.Add(temp)
            }
            return
        }
    }

    if insertIndex == -1 {
        insertIndex = len(rs.rangeSet)
    }

    insert := &Range{
        start: r.start,
        end:   r.end,
    }
    rs.rangeSet = append(rs.rangeSet, insert)
    copy(rs.rangeSet[insertIndex+1:], rs.rangeSet[insertIndex:])
    rs.rangeSet[insertIndex] = insert
}

//RemoveBefore remove range before index
func (rs *RangeSet) RemoveBefore(removeRangeEnd uint64) {
    removeIndex := -1
    for index, v := range rs.rangeSet {
        if v.end >= removeRangeEnd {
            break
        }
        removeIndex = index
    }
    if removeIndex == -1 {
        return
    }
    removeIndex = removeIndex + 1
    if removeIndex >= len(rs.rangeSet) {
        rs.rangeSet = make([]*Range, 0)
        return
    }
    rs.rangeSet = rs.rangeSet[removeIndex:]
}

func (rs *RangeSet) String() string {
    builder := strings.Builder{}
    for _, r := range rs.rangeSet {
        builder.WriteString("[")
        builder.WriteString(fmt.Sprintf("%d", r.start))
        builder.WriteString(string(","))
        builder.WriteString(fmt.Sprintf("%d", r.end))
        builder.WriteString("] ")
    }
    return builder.String()
}

func (rs *RangeSet) MissingCount() int {
    if len(rs.rangeSet) == 0 {
        return 0
    }
    var start uint64 = 1
    missingCount := 0
    for _, r := range rs.rangeSet {
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

func (rs *RangeSet) Contains(index uint64) bool {
    for _, r := range rs.rangeSet {
        if index >= r.start && index <= r.end {
            return true
        }
    }
    return false
}

func (rs *RangeSet) index(index uint64) int {
    for i, r := range rs.rangeSet {
        if index >= r.start && index <= r.end {
            return i
        }
    }
    return -1
}

func haveSameIndex(a *Range, b *Range) bool {
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
