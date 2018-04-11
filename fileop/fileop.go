package fileop

import (
	"io/ioutil"
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetDirFiles(dir string, files *[]string) error {
	fileInfos, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			subDir := filepath.Join(dir, info.Name())
			GetDirFiles(subDir, files)
			continue
		}

		//if path.Ext(info.Name()) == ".log" {
			s := filepath.Join(dir, info.Name())
			*files = append(*files, s)
		//}
	}

	return nil
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		handler(line)
	}

	return nil
}

func SaveFile(file string, lineList []string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	for _, line := range lineList {
		_, e := f.WriteString(line + "\r\n")
		if e != nil {
			return e
		}
	}

	return nil
}
