package unit_test

import (
	"github.com/ice-cream-heaven/utils/unit"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Log(unit.DurationMinuteSecond(time.Second * 10))
}
