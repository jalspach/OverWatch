package util

import (
	"log"
	"net"

	"github.com/garethpaul/purpleair-go"

	//"os"
	"fmt"
	"time"
)

// Get preferred outbound ip of this machine by reaching out to the world
func GetOutsideIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// useful links:
// https://stackoverflow.com/questions/27410764/dial-with-a-specific-address-interface-golang
// https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6

//Gets first local ipv4 address it finds on given interface
func GetInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return "", fmt.Errorf(fmt.Sprintf("interface %s doesn't have an ipv4 address\n", interfaceName))
	}
	return ipv4Addr.String(), nil
}

//checks to see if a port "picks up"
func PortCheckSimpleTxt(host, port string) string {
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		//fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
		return "not responding on port"
	} else {
		//fmt.Printf("%s %s %s\n", host, "responding on port:", port)
		return "responding on port"
	}
}

func PortCheckSimpleBool(host, port string) string {
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		//fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
		return "0"
	} else {
		//fmt.Printf("%s %s %s\n", host, "responding on port:", port)
		return "1"
	}
}

func CheckAQI(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	for i := 0; i < len(s.Results); i++ {
		fmt.Println("Air Quality: " + s.Results[i].PM25Value)
	}
	return "all good"
}

func CheckAQI1(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].PM25Value
}

func CurrentDTS() string {
	//currenttime := time.Now()
	//fmt.Println(currenttime.Format())
	fmt.Println(time.Now())
	//fmt.Println(time.Now().Date())
	return time.Now().Format("2006-01-02 15:04:05 PST")
}
