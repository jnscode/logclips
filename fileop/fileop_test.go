package fileop

import (
	"fmt"
    "testing"
)

func TestReadLine(t *testing.T){
	var lines []string
	ReadLine("d:/1.txt", func(line string) {
		lines = append(lines, line)
	})

	fmt.Println(lines)
}