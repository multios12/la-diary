package main

import (
	"embed"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

//go:embed static/*
var static embed.FS
var months []string
var dataDir string

type lineModel struct {
	Day  string // 日付(yyyy-mm-dd形式)
	Memo string // 本文
}
type listModel struct {
	WritedMonths []string // 記載された年月(yyyy-mm形式)
	Lines        []lineModel
}

// コンテキスト
func main() {
	flag.StringVar(&dataDir, "dir", "./data", "data directory")
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	dataDir, _ = filepath.Abs(dataDir)
	if _, err := os.Stat(dataDir); err != nil {
		os.Mkdir(dataDir, 0777)
	}

	months = getWritedMonths()

	router := routerInitial(static)
	router.Run(*port)
}

func getWritedMonths() []string {
	l, err := ioutil.ReadDir(dataDir)
	var months []string
	if err != nil {
		return months
	}

	r := regexp.MustCompile(`\d\d\d\d\d\d\.txt`)

	for _, f := range l {
		if !f.IsDir() && r.MatchString(f.Name()) {
			v := f.Name()[:4] + "-" + f.Name()[4:6]
			months = append(months, v)
		}
	}

	if len(months) > 0 {
		sort.Slice(months, func(i int, j int) bool { return months[i] > months[j] })
	} else {
		months = append(months, time.Now().Format("200601"))
	}

	return months
}
