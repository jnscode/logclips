package clips

import (
    "testing"
)

func TestGetLogTime(t *testing.T){
	var log string
	log = "02-26 09:20:22"
	b, s := GetLogTime(log)
	if b {
		t.Error(log)
	}

	log = "02-26 09:20:22 339[INFO]: [] [3600:4832] "
	b, s = GetLogTime(log)
	if !b {
		t.Error(log)
	}else if s != "2018-02-26 09:20:22"{
		t.Error(s)
	}

	log  = "2018-02-26 09:20:22 339[INFO]: [] [3600:4832] "
	b, s = GetLogTime(log)
	if !b {
		t.Error(log)
	}else if s != "2018-02-26 09:20:22"{
		t.Error(s)
	}
}