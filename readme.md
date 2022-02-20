# Network and cabinet monitor #
## Basic Plan ##
This was designed to work with my hardware project which is a 2 layer board that plugs into a rasp pi b+ (26 pin header). This unit is placed in a cabinet, attached to the network and powered up.

Its designed to watch the network and report out (via LEDs and MQTT) if there are problems. VERY SIMPLE Red, Yellow Green (plus flashing) status.
Also has a built in temperature sensor to report out more status on the cabinet it is in.

All of this gets dumped to MQTT for remote monitoring and trending. This device can also be a target for simple connectivity tests (the NIC will not support accurate speed tests on any modern network).

*Hopefully*...if I have read the pages and studied right...this will be a *proper* module (for a project that really doesn't need one but whatever...I am learning here). :-)

## Followed these pages: ##

* Modules 
    * https://golangbyexample.com/package-vs-module-golang/
    * https://www.golang-book.com/books/intro/11
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
    
* Go routines / packages
    * https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/
    * https://go.dev/blog/context
    * https://pkg.go.dev/context
    * https://gobyexample.com/
* Purple Air
    * https://github.com/garethpaul/purpleair-go/blob/master/results.go
    * https://www2.purpleair.com/community/faq#hc-json-object-fields
    * https://docs.google.com/document/d/15ijz94dXJ-YAZLi9iZ_RaBwrZ4KtYeCy08goGBwnbCU/edit
* SpeedTest
    * https://github.com/showwin/speedtest-go

## ToDo: ##

* -[x] Module to talk to LED's
* -[x] Module to talk to Temp Sensor - Random temp spikes messed up graphing so I now check 3 times and keep the middle one.
* -[x] Module to upload data to MQTT
* -[x] Get internal IP address
* -[] Get external address
* -[] Run network test(s) <- simple version. Need to use the original output but figure out how to shrink it into one variable. Good test would be to report on packet loss and or average ms rt time
* -[] set the LED's as appropriate (temp or network)	
* -[] *Maybe do this using goroutines if thats not too much extra work* <- EZ PEASY so far  - How to stop a goroutine that is running, after a foreground process is done?
* -[] build run types
* -[] add MQTT Tombstone per example in Hive binder I have
* -[] Pull specifics from env or local file (allows it to run from work or home easily)
* -[x] pull  clientname (topic) in code
* -[x] add PurpleAir package

Based on run type (below) do the above as appropriate
	
## Run types ##
Name | Description or Function
------|-----------------------
oneoff | test leds, report temprature while doing nextork test? Threads?
continious | test leds, report temp while doing network test, loop over these two
service | skip LED test
*other* | *thoughts*
