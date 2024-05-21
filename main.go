package main

import (
	"github.com/intyouss/Traceability/cmd"
)

// @title Traceability
// @version 0.0.1
// @description A web-based short video platform
func main() {
	defer cmd.Clean()
	cmd.Start()
}
