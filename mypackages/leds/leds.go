package leds

import (
	"fmt"
	"time"

	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/stianeikeland/go-rpio"
)

//init LEDS - opens the channel and resets LEDS to low.
func Init() int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()
	yellow := rpio.Pin(27)
	yellow.Output()
	red := rpio.Pin(22)
	red.Output()

	red.Low()
	yellow.Low()
	green.Low()
	return 0
}

/*
 Set the status of the LED's based on the the provided Hex value. If an invalid number is passed will return 0X_BEEF if everything worked it will return 0x_FFFF.
 Uses binary to set the status. 1 bit for each LED. Starting with the LSB its Red, Green, Yellow.
 Add 8 (b1000) for flash however, currtnly is just toggles that bit from what it was. ToDo: build flash, cover other combos.
*/
func Setstatus(hex int) int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()
	yellow := rpio.Pin(27)
	yellow.Output()
	red := rpio.Pin(22)
	red.Output()

	// Set the status of the leds
	switch hex {
	case 0x_0:
		//all off
		red.Low()
		yellow.Low()
		green.Low()
	case 0x_1:
		//Red solid
		red.High()
		yellow.Low()
		green.Low()
	case 0x_2:
		//Yellow solid
		red.Low()
		yellow.High()
		green.Low()
	case 0x_3:
		//Red and Yellow solid
		red.High()
		yellow.High()
		green.Low()
	case 0x_4:
		//Green solid
		red.Low()
		yellow.Low()
		green.High()

	case 0x_5:
		//Red and Green solid
		red.High()
		yellow.Low()
		green.High()
	case 0x_6:
		//Yellow and Green solid
		red.Low()
		yellow.High()
		green.High()
	case 0x_7:
		//Red, Green and Yellow solid
		red.High()
		yellow.High()
		green.High()
	case 0x_8:
		//Not Valid (all off but still flashing)
	case 0x_9:
		//Red flashing
		yellow.Low()
		green.Low()
		red.High()
		time.Sleep(time.Millisecond * 250)
		red.Low()
		time.Sleep(time.Millisecond * 250)
		red.High()
		time.Sleep(time.Millisecond * 250)
		red.Low()

	case 0x_A:
		//Yellow flashing

		red.Low()
		green.Low()
		yellow.High()
		time.Sleep(time.Millisecond * 250)
		yellow.Low()
		time.Sleep(time.Millisecond * 250)
		yellow.High()
		time.Sleep(time.Millisecond * 250)
		yellow.Low()

	case 0x_B:
		//Red and Yellow flashing
		green.Low()
		yellow.High()
		red.High()
		time.Sleep(time.Millisecond * 250)
		yellow.Low()
		red.Low()
		time.Sleep(time.Millisecond * 250)
		yellow.High()
		red.High()
		time.Sleep(time.Millisecond * 250)
		yellow.Low()
		red.Low()

	case 0x_C:
		//Green flashing
		red.Low()
		yellow.Low()
		green.High()
		time.Sleep(time.Millisecond * 250)
		green.Low()
		time.Sleep(time.Millisecond * 250)
		green.High()
		time.Sleep(time.Millisecond * 250)
		green.Low()
	case 0x_D:
		//Red Green flashing
		yellow.Low()
		red.High()
		green.High()
		time.Sleep(time.Millisecond * 250)
		red.Low()
		green.Low()
		time.Sleep(time.Millisecond * 250)
		red.High()
		green.High()
		time.Sleep(time.Millisecond * 250)
		green.Low()
		red.Low()

	case 0x_E:
		//Yellow and Green flashing

		red.Low()
		yellow.High()
		green.High()
		time.Sleep(time.Millisecond * 250)
		yellow.Low()
		green.Low()
		time.Sleep(time.Millisecond * 250)
		yellow.High()
		green.High()
		time.Sleep(time.Millisecond * 250)
		green.Low()
		yellow.Low()

	case 0x_F:
		//Red, Yellow and Green flashing
		red.High()
		yellow.High()
		green.High()
		time.Sleep(time.Millisecond * 250)
		red.Low()
		yellow.Low()
		green.Low()
		time.Sleep(time.Millisecond * 250)
		red.High()
		yellow.High()
		green.High()
		time.Sleep(time.Millisecond * 250)
		red.Low()
		green.Low()
		yellow.Low()
	default:
		return 0x_BEEF
	}
	return 0x_FFFF
}

func Setthinking() int {
	//build these to call setstatus with a specific hex value instead of setting LED's manually
	// set alert, set bootup, etc...
	return 0
}

//Sweep the LEDS from Red to Green using the dwell time (in milliseconds)
func SweepR2G(dwell int) int {
	Setstatus(0x0)
	for x := 0; x < 10; x++ {
		Setstatus(0x1)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
		Setstatus(0x2)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
		Setstatus(0x4)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
	}
	Setstatus(0x0)
	return 0
}

//Sweep the LEDS from Green to Red using the dwell time (in milliseconds)
func SweepG2R(dwell int) int {
	Setstatus(0x0)
	for x := 0; x < 10; x++ {
		Setstatus(0x4)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
		Setstatus(0x2)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
		Setstatus(0x1)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
	}
	Setstatus(0x0)
	return 0
}

// Displays an error combo of flashing leds using the dwell time (in milliseconds)
func Error(dwell int) int {
	Setstatus(0x0)
	for x := 0; x < 10; x++ {
		Setstatus(0x2)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
		Setstatus(0x5)
		time.Sleep(time.Duration(dwell) * time.Millisecond)
	}
	Setstatus(0x0)
	return 0
}

// Runs through all of the statusus in order to test LED fuctionality
func Test() int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()
	yellow := rpio.Pin(27)
	yellow.Output()
	red := rpio.Pin(22)
	red.Output()

	for x := 0; x < 17; x++ {
		Setstatus(x)
		time.Sleep(time.Second * 1)

	}
	return 0
}

//Display Temprature on LED's
func DisplayTemp() int {
	coms.Listb20s()
	var c float64 = coms.Tempc()
	fmt.Printf("%f\n", c)

	f := c/(.556) + 32
	fmt.Printf("%f\n", f)
	switch {
	case f > 100:
		Setstatus(0x_1)
	case f > 90:
		Setstatus(0x_3)
	case f > 84:
		Setstatus(0x_6)
	case f > 75:
		Setstatus(0x_4)
	case f > 68:
		Setstatus(0x_6)
	case f > 60:
		Setstatus(0x_3)
	default:
		Setstatus(0x_1)
	}
	return 0
}

// returns status of the LED's in bianry
// want to know whether an LED is on, off or flashing
// binary (LSB) Red, Yellow Green, Blink
// 0 all off, 1 Red solid, 2 yellow solid, 3 red and yellow solid, 4 green solid, 5 red and green solid, 6 green and yellow solid, 7 red yellow and green solid
// 8 NA (All off and flashing), 9 Red Flashing, A yellow flashing, B red and yellow flashing, C green flashing, D Green and Red flashing, E green and yellow flashing, green yellow and red flashing

func Checkstatus() int {
	//returns the status of the LED's
	return 0
}
