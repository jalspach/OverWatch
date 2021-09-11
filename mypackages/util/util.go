package util

import (
	"log"
	"net"
	//"os"
	"fmt"
	"errors"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetlocalIP(host string)  {
addrs, _ := net.LookupIP(host)
for _, addr := range addrs {
    if ipv4 := addr.To4(); ipv4 != nil {
        fmt.Println("IPv4: ", ipv4)
    }   
}
}

// useful links:
// https://stackoverflow.com/questions/27410764/dial-with-a-specific-address-interface-golang
// https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6
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
        return "", errors.New(fmt.Sprintf("interface %s don't have an ipv4 address\n", interfaceName))
    }
    return ipv4Addr.String(), nil
}
func Portcheck (host, port string ) {
timeout := time.Duration(1 * time.Second)
_, err := net.DialTimeout("tcp", host+":"+port, timeout)
if err != nil {
	fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
} else {
	fmt.Printf("%s %s %s\n", host, "responding on port:", port)
}
}
