package main

import (
	"fmt"
	"runtime"

	"github.com/bennygunawan/go-buher-main/cmd"
)

// @title Bukit Hermon Go Apps
// @version 1.0
// @description API for LMS Module of Bukit Hermon Go Apps
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("starting app")
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
