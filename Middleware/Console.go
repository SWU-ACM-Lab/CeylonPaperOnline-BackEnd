package Middleware

import "fmt"

type systemConsole struct {
	logEnable bool
	index int
}

var Console systemConsole

func (s* systemConsole) Log (info error, description string) {
	s.index += 1
	if s.logEnable {
		fmt.Printf("[%d]\tCeylonSystem inner error in step %s :\n\t%s", s.index, description, info)
	}
}

func (s* systemConsole) SetStatus (status bool) {
	s.logEnable = status
}