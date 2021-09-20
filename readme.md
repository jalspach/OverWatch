# Network and cabinet monitor #
## Basic Plan ##
This was designed to work with my hardware project which is a 2 layer board tht plugs into a rasp pi b+ (26 pin header). This unit is placed in a cabinet, attached to the network and powered up.

Its designed to watch the network and report out (via leds and MQTT) if there are problems. VERY SIMPLE Red, Yellow Green (plus flashing) status.
Also has a built in temprature sensor to report out more status on the cabinet it is in.

All of this gets dumped to MQTT for remote monitoring and tranding. This device can also be a target for simople conectivity tests (the NIC will not support accutrate speed tests on any modern network).

*Hopefully*...if I have read the pages and studied right...this will be a *proper* module (for a project that really doesnt need one but whatever...I am learning here). :-)

## Followed these pages: ##

* Modules 
    * https://golangbyexample.com/package-vs-module-golang/
    * Need links
* MQTT
    * https://www.cloudmqtt.com/docs/go.html
    * https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
    * https://www.eclipse.org/paho/index.php?page=clients/golang/index.php
    * http://mqtt-explorer.com/
* 1wire
    * Add
    * Add
* Utils
    * https://gist.github.com/schwarzeni/f25031a3123f895ff3785970921e962c#file-util-go
    * https://pkg.go.dev/github.com/sparrc/go-ping#readme-installation
    * https://github.com/garethpaul/purpleair-go/blob/master/results.go
* Go routines / packages
    * https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/
    * https://go.dev/blog/context
    * https://pkg.go.dev/context


## ToDo: ##

* -[x] Module to talk to LED's
* -[x] Module to talk to Temp Sensor - Random temp spikes messed up graphing so I now check 3 times and keep the middle one.
* -[x] Module to upload data to MQTT
* -[x] Get internal IP address
* -[] Get external address
* -[] Run network test(s) <- smple version. Need to use the origional output but figure out how to shrink it into one variable. Good test would be to report on packet loss.
* -[] set the LED's as apropriate (temp or network)	
* -[] *Maybe do this using goroutines if thats not too much extra work* <- EZ PEASY so far 
* -[] build run types
* -[] add MQTT Tombstone per example in Hive binder I have
* -[] Pull specifics from env or local file (allows it to run from work or home easily)
* -[] pull  clienname (topic) in code
* -[] add PurpleAir package

Based on run type (below) do the above as apropirate
	
## Run types ##
Name | Description or Function
------|-----------------------
oneoff | test leds, report temprature while doing nextork test? Threads?
continious | test leds, report temp while doing network test, loop over these two
service | skip LED test
*other* | *thoughts*
