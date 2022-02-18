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

	assert.Equal(dp.Count, count, "count not equal")
	assert.Equal(dp.Timestamp, timestamp, "timestamp not equal")
}
