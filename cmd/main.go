package main

import (
	"vtp-apis/external/delivery/rest"
	"vtp-apis/packages/ginh"
	"vtp-apis/packages/shutdown"
)

// @title rest apis vtp
// @version 1.0.0
// @description List APIs of VTP
// @in header
// @name rest apis
func main() {
	defer shutdown.SigtermHandler().Wait()
	registerErrorsMaps()
	s := initServer()
	s.Run()
}

func registerErrorsMaps() {
	ginh.RegisterDefaultErrorsMap(rest.ErrsMapDefault)
}
