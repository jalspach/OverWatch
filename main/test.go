package main

import (
	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {
	leds.Test()
	leds.Error(3, 3)
	leds.Test()
	coms.Tempc()

}
