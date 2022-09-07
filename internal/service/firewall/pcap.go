package firewall

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

type PacketCapture struct {

}

type Inet struct {
	Name string
	Ip string
	Version string
}

func NewPcap() *PacketCapture {
	return &PacketCapture{}
}

func (pc *PacketCapture) GetDevs() (inet Inet) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v",devices)
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
	inet = Inet{
		Name: "eth0",
		Ip : "192.168.1.1",
		Version: "ipv4",
	}
	return
}

