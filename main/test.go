package main

import (
	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {
	leds.Test()
	leds.Error()
	leds.Test()
	coms.Tempc()

}
