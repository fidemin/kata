package counter

import "time"

type DataPointList struct {
	Head       *DataPoint
	Tail       *DataPoint
	TotalCount int
}

func NewDataPointList() *DataPointList {
	return &DataPointList{}
}

func (dl *DataPointList) Push(dp *DataPoint) {
	dl.TotalCount += dp.Count
	if dl.Head == nil {
		// for empty DataPointList
		dl.Head = dp
		dl.Tail = dp
		return
	}

	headDp := dl.Head

	dp.Before = headDp
	headDp.Next = dp
	dl.Head = dp
}

func (dl *DataPointList) Pop() *DataPoint {
	if dl.Head == nil && dl.Tail == nil {
		return nil
	}

	tailDp := dl.Tail
	nextDp := dl.Tail.Next

	tailDp.Next = nil
	dl.TotalCount -= tailDp.Count

	if nextDp == nil {
		dl.Head = nil
		dl.Tail = nil
		return tailDp
	}

	dl.Tail = nextDp
	nextDp.Before = nil

	return tailDp
}

type DataPoint struct {
	Next      *DataPoint
	Before    *DataPoint
	Timestamp time.Time
	Count     int
}

func NewDataPoint(count int, timestamp time.Time) *DataPoint {
	return &DataPoint{
		Count:     count,
		Timestamp: timestamp,
	}
}
