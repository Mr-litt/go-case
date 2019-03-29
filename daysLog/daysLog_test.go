package daysLog

import (
	"testing"
	"log"
)

func TestDaysLog(t *testing.T) {
	daysLog := NewDaysLog()

	err := daysLog.Info("info1", "test", "TestDaysLog")
	if err != nil {
		log.Println(err)
	}

	daysLog.Info("info2", "test", "TestDaysLog")
	daysLog.Error("error1", "test", "TestDaysLog")
}
