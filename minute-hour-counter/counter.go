package main

import (
	"fmt"
	"time"
)

type Counter struct {
	minuteList *DataPointList
	hourList   *DataPointList
}

func NewCounter() *Counter {
	return &Counter{
		minuteList: NewDataPointList(),
		hourList:   NewDataPointList(),
	}
}

// adds count (in bytes)
func (c *Counter) Add(count int) {
	dp := NewDataPoint(count, time.Now())
	c.minuteList.Push(dp)
	c.hourList.Push(dp)
}

// returns counts received during one minute (60 seconds) from current
func (c *Counter) MinuteCount() int {
	now := time.Now()
	return c.count(c.minuteList, 60*time.Second, now)
}

// returns counts received during one hour (60 minutes) from current
func (c *Counter) HourCount() int {
	now := time.Now()
	return c.count(c.hourList, 60*time.Minute, now)
}

func (c *Counter) count(dl *DataPointList, window time.Duration, now time.Time) int {
	for {
		currentDp := dl.Tail
		if now.Sub(currentDp.Timestamp) <= window {
			break
		}

		dl.Pop()
	}
	return dl.TotalCount
}

func main() {
	counter := NewCounter()
	counter.Add(10)
	counter.Add(20)
	fmt.Println(counter.MinuteCount())
	fmt.Println(counter.HourCount())
}
