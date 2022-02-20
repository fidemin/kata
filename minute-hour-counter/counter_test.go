package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTime struct {
	mock.Mock
}

func NewMockTime(t time.Time) *mockTime {
	testTime := new(mockTime)
	testTime.On("Now").Return(t)
	return testTime
}

func (m *mockTime) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)

}

func TestCounterWithRealTime(t *testing.T) {
	assert := assert.New(t)

	counter := NewCounter()
	bytes1 := 3
	bytes2 := 4
	bytes3 := 6

	counter.Add(bytes1)
	counter.Add(bytes2)
	counter.Add(bytes3)

	sum := bytes1 + bytes2 + bytes3
	assert.Equal(sum, counter.MinuteCount())
	assert.Equal(sum, counter.HourCount())
}

func TestCounterHourCount(t *testing.T) {
	assert := assert.New(t)

	mockCounter := NewCounter()

	utcLoc := time.UTC

	bytes1 := 1
	bytes2 := 2
	bytes3 := 3

	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 0, 0, 0, utcLoc))
	mockCounter.Add(bytes1)

	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 10, 0, 0, utcLoc))
	mockCounter.Add(bytes2)

	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 20, 0, 0, utcLoc))
	mockCounter.Add(bytes3)

	// change current time to count all bytes
	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 1, 0, 0, 0, utcLoc))
	assert.Equal(bytes1+bytes2+bytes3, mockCounter.HourCount())

	// change current time to remove bytes1, bytes2
	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 1, 20, 0, 0, utcLoc))
	assert.Equal(bytes3, mockCounter.HourCount())

	// change current time to remove all
	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 1, 30, 0, 0, utcLoc))
	assert.Equal(0, mockCounter.HourCount())
}

func TestCounterMinuteCount(t *testing.T) {
	assert := assert.New(t)

	mockCounter := NewCounter()

	utcLoc := time.UTC

	bytes1 := 1
	bytes2 := 2
	bytes3 := 3

	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 0, 0, 0, utcLoc))
	mockCounter.Add(bytes1)

	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 0, 10, 0, utcLoc))
	mockCounter.Add(bytes2)

	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 0, 20, 0, utcLoc))
	mockCounter.Add(bytes3)

	// change current time to count all bytes
	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 1, 0, 0, utcLoc))
	assert.Equal(bytes1+bytes2+bytes3, mockCounter.MinuteCount())

	// change current time to remove bytes1, bytes2
	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 1, 20, 0, utcLoc))
	assert.Equal(bytes3, mockCounter.MinuteCount())

	// change current time to remove all
	mockCounter.realTime = NewMockTime(time.Date(2022, 01, 1, 0, 1, 30, 0, utcLoc))
	assert.Equal(0, mockCounter.MinuteCount())
}
