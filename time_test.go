package utils

import "testing"

func TestToTime(t *testing.T) {
	t.Log(ToTime("yyyy-MM-dd", "2017-09-14").Unix())
}
