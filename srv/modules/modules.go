package routes

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var DataDir string

func ReadLine(day string) lineModel {
	lines := ReadMonthFile(day[0:4] + day[5:7])
	for _, l := range lines {
		if day == l.Day {
			return l
		}
	}
	return lineModel{}
}

func WriteLine(day string, memo string) {
	month := day[0:4] + day[5:7]
	lines := ReadMonthFile(month)
	flag := false
	for i, l := range lines {
		if day == l.Day {
			lines[i].Memo = memo
			flag = true
		}
		lines[i].Memo = strings.TrimSpace(lines[i].Memo)
	}
	if flag {
		lines = append(lines, lineModel{day, memo})
	}
	writeMonthFile(month, lines)
	return
}

func ReadMonthFile(month string) []lineModel {
	p := month + ".txt"
	p = filepath.Join(DataDir, p)

	_, err := os.Stat(p)
	if err != nil {
		return []lineModel{}
	}

	fp, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 5000)
	var lines []lineModel
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		s := string(line)
		day := s[0:10]
		s = s[10:]

		d := lineModel{day, s}
		lines = append(lines, d)
	}

	return lines
}

func writeMonthFile(month string, lines []lineModel) {
	sort.Slice(lines, func(i, j int) bool { return lines[i].Day > lines[j].Day })

	p := month + ".txt"
	p = filepath.Join(DataDir, p)

	dataFile := ""
	for _, l := range lines {
		if dataFile != "" {
			dataFile += "\n"
		}
		dataFile += l.Day + " " + l.Memo
	}
	ioutil.WriteFile(p, []byte(dataFile), os.ModePerm)
}
