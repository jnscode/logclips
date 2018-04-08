package main

import (
	"bufio"
	"flag"
	"fmt"
	"logclips/clips"
	"logclips/fileop"
	"os"
	"path/filepath"
	"time"
)

func procError(title string, err error) {
	if err != nil {
		fmt.Println(title, ",", err.Error())
		os.Exit(1)
	}
}

func testTime() {
	tm1, e := time.Parse("2006-01-02 15:04:05", "2018-03-31 16:31:08")
	if e != nil {
		fmt.Println("error", e.Error())
	} else {
		fmt.Println(tm1.Format("2006/01/02 03-04-05"))
	}

	tm2, e := time.Parse("01-02 15:04:05", "03-31 16:31:08")
	if e != nil {
		fmt.Println("error", e.Error())
	} else {
		fmt.Println(tm2.Format("01/02 03-04-05"))
	}

	//tm2.Add = 2018
	if tm2.Before(tm1) {
		fmt.Println("tm2 big")
	} else if tm2.After(tm1) {
		fmt.Println("tm1 big")
	} else {
		fmt.Println("same")
	}
}

func main() {
	flag.Parse()

	/*
		if flag.NArg() < 2 {
			fmt.Println("====== Usage ======")
			fmt.Println("Param1: log folder")
			fmt.Println("Param2: begin time")
			fmt.Println("Param3: end time")
			fmt.Println()
			return
		}
	*/

	for {
		var dir string
		var tm string
		var tmb time.Time
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("Please input log dir...")
			input, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			dir = string(input)
			exists, err := fileop.PathExists(dir)
			if !exists {
				fmt.Errorf("Error, dir not exist", err)
				continue
			}

			break
		}

		for {
			fmt.Println("Please input a begin time, like 2006-01-02 15:04:05...")

			input, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			tm = string(input)
			tmb, err = clips.Str2Time(tm)
			if err != nil {
				fmt.Println("Invalid time", err.Error())
				continue
			}

			break
		}

		dirDst := dir + " after " + tmb.Format("01-02 15.04.05")
		fmt.Println("\nCopy logs after", tm, "from", filepath.Base(dir), "to dir", filepath.Base(dirDst))

		clips.ClipLog(dir, dirDst, tmb)
		fmt.Println("Process end")
		fmt.Println("")
	}
}
