package clips

import (
	"os"
	"fmt"
	"logclips/fileop"
	"path/filepath"
	"time"
)

func GetLogTime(log string) (bool, string) {
	var str string
	cnt := len(log)
	if cnt < 19 {
		return false, str
	}

	if log[0] != '2' {
		str = "2018-" + log[:14]
	} else {
		str = log[:19]
	}

	if str[4] != '-' || str[7] != '-' || str[13] != ':' || str[16] != ':' {
		return false, str
	}

	return true, str
}

func Str2Time(str string) (time.Time, error) {
	tm, e := time.Parse("2006-01-02 15:04:05", str)
	return tm, e

	//	tm, e = time.Parse("01-02 15:04:05", str)
	//	if e == nil {
	//		return tm, e
	//	}
}

func ClipLog(dirSrc, dirDst string, tmb time.Time) bool {
	var files []string
	fileop.GetDirFiles(dirSrc, &files)

	dirLen := len(dirSrc)

	for i, f := range files {
		files[i] = f[dirLen:]
	}

	cnt := len(files)

	for i, f := range files {
		fmt.Println("process", i+1, "/", cnt)

		var lines []string
		err := fileop.ReadLine(filepath.Join(dirSrc, f), func(line string) {
			lines = append(lines, line)
		})

		if err != nil {
			fmt.Println("read file failed,", err.Error())
			continue
		}

		var lines2 []string
		for i, line := range lines {
			hasTime, tmStr := GetLogTime(line)
			if !hasTime {
				continue
			}

			tm, e := Str2Time(tmStr)
			if e != nil {
				continue
			}

			if tm.After(tmb) {
				lines2 = lines[i:]
				break
			}
		}

		if len(lines2) == 0 {
			//fmt.Println("No log to save")
			continue
		}

		file := filepath.Join(dirDst, f)
		os.MkdirAll(filepath.Dir(file), os.ModePerm)
		//fmt.Println("Save log to", f)
		
		e := fileop.SaveFile(file, lines2)
		if e != nil {
			fmt.Println(e.Error())
		}
	}

	return true
}
