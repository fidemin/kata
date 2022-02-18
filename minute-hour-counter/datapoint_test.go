package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDataPoint(t *testing.T) {
	assert := assert.New(t)

	timestamp := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})
	count := 10
	dp := NewDataPoint(count, timestamp)

	assert.Equal(count, dp.Count, "count not equal")
	assert.Equal(timestamp, dp.Timestamp, "timestamp not equal")
}

func TestDataPointListPush(t *testing.T) {
	assert := assert.New(t)

	// distiguish data point with count value
	count1 := 1
	timestamp1 := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})

	count2 := 2
	timestamp2 := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})

	count3 := 3
	timestamp3 := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})

	dp1 := NewDataPoint(count1, timestamp1)
	dp2 := NewDataPoint(count2, timestamp2)
	dp3 := NewDataPoint(count3, timestamp3)

	dpList := NewDataPointList()
	dpList.Push(dp1)
	assert.Equal(count1, dpList.TotalCount)
	dpList.Push(dp2)
	assert.Equal(count1+count2, dpList.TotalCount)
	dpList.Push(dp3)
	assert.Equal(count1+count2+count3, dpList.TotalCount)

	assert.Equal(3, dpList.Head.Count)
	assert.Equal(1, dpList.Tail.Count)
	assert.Equal(1, dp2.Before.Count)
	assert.Equal(3, dp2.Next.Count)

}

func TestDataPointListPop(t *testing.T) {
	assert := assert.New(t)

	// distiguish data point with count value
	count1 := 1
	timestamp1 := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})

	count2 := 2
	timestamp2 := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})

	count3 := 3
	timestamp3 := time.Date(2010, 10, 1, 0, 0, 0, 0, &time.Location{})

	dp1 := NewDataPoint(count1, timestamp1)
	dp2 := NewDataPoint(count2, timestamp2)
	dp3 := NewDataPoint(count3, timestamp3)

	dpList := NewDataPointList()
	dpList.Push(dp1)
	dpList.Push(dp2)
	dpList.Push(dp3)

	pop1 := dpList.Pop()
	assert.Nil(pop1.Before)
	assert.Nil(pop1.Next)
	assert.Equal(count1, pop1.Count)
	assert.Equal(count2+count3, dpList.TotalCount)

	pop2 := dpList.Pop()
	assert.Nil(pop2.Before)
	assert.Nil(pop2.Next)
	assert.Equal(count2, pop2.Count)
	assert.Equal(count3, dpList.TotalCount)

	pop3 := dpList.Pop()
	assert.Nil(pop3.Before)
	assert.Nil(pop3.Next)
	assert.Equal(count3, pop3.Count)
	assert.Equal(0, dpList.TotalCount)

	assert.Nil(dpList.Head)
	assert.Nil(dpList.Tail)

}
