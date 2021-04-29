package pid

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositiveDelta(t *testing.T) {
	controller := New(0.25, 0.25, 0.25, 10)
	assert.NotNil(t, controller)
	controller.SetTarget(120)
	delta := controller.Update(10.0)
	assert.Greater(t, delta, 0.0, "Delta is positive")
}

func TestNegativeDelta(t *testing.T) {
	controller := New(0.25, 0.25, 0.25, 10)
	controller.SetTarget(120.0)
	delta := controller.Update(125.0)
	assert.Less(t, delta, 0.0, "Delta should be a negative number")
}

func TestNumberConvergesToTarget(t *testing.T) {

	controller := New(0.25, 0.25, 0.25, 10)
	controller.SetTarget(120.0)
	measure := 0.0
	for i := 0; i < 100; i++ {
		delta := controller.Update(measure)
		measure += delta
	}
	measureDiff := math.Abs(120.0 - measure)
	assert.LessOrEqual(t, measureDiff, 0.1, "Error should be minimal")
}
