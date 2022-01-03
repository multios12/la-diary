package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/multios12/la-diary/routes"
)

var proxyMode bool

// コンテキスト
func main() {
	flag.StringVar(&routes.DataDir, "dir", "./data", "data directory")
	flag.BoolVar(&proxyMode, "ReverseProxyMode", false, "reverse proxy mode")
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	routes.DataDir, _ = filepath.Abs(routes.DataDir)
	_, err := os.Stat(routes.DataDir)
	if err != nil {
		os.Mkdir(routes.DataDir, 0777)
	}

	router := routes.Initial(proxyMode)
	router.Run(*port)
}
