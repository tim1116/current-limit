package common

import (
	"testing"
	"time"
)

func TestShowTime(t *testing.T) {
	tt := time.Now()
	ShowTime(tt)
}
