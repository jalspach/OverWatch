package main

import (
	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {
	leds.Error()
	leds.Test()
	coms.Tempc()
	leds.Reset()

}
