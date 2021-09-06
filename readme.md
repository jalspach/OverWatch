#Network and cabinet monitor#
##Basic Plan##
This was designed to work with my hardware project which is a 2 layer board tht plugs into a rasp pi b+ (26 pin header). This unit is placed in a cabinet, attached to the network and powered up.

Its designed to watch the network and report out (via leds and MQTT) if there are problems. VERY SIMPLE Red, Yellow Green (plus flashing) status.
Also has a built in temprature sensor to report out more status on the cabinet it is in.

ALl of this gets dumped to MQTT for remote monitoring and tranding. This device can also be a target for simople conectivity tests (the NIC will not support accutrate speed tests on any modern network).

*Hopefully*...if I have read the pages and studied right...this will be a *proper* module (for a project that really doesnt need one but whatever...I am learning here). :-)

##ToDo:##
-[x] module to talk to LED's
-[x] module to talk to Temp Sensor
-[] Upload data to MQTT (IP Address and temp)
-[] Run network test(s)
-[] Upload data to MQTT (test results)
-[] set the LED's as apropriate (temp or network)
	
-[] *Maybe do this using threads if thats not too much extra work*

Based on run type (below) do the above as apropirate
	
##Run types##
Name|Description or Function
-----------------------------
oneoff|test leds, report temprature while doing nextork test? Threads?
continious|test leds, report temp while doing network test, loop over these two
service|skip LED test
*other*|*thoughts*
