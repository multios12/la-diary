package main

import (
	"embed"
	"flag"
	"os"
	"path/filepath"

	modules "github.com/multios12/la-diary/modules"
)

//go:embed static/*
var local embed.FS

// コンテキスト
func main() {
	flag.StringVar(&modules.DataDir, "dir", "./data", "data directory")
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	modules.DataDir, _ = filepath.Abs(modules.DataDir)
	_, err := os.Stat(modules.DataDir)
	if err != nil {
		os.Mkdir(modules.DataDir, 0777)
	}

	router := modules.Initial(local)
	router.Run(*port)
}
