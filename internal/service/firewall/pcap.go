package firewall

import (
	"gitee.com/bytesworld/tomato/internal/logger"
	_inet "gitee.com/bytesworld/tomato/pkg/utils/inet"
	"github.com/google/gopacket/pcap"
	"net"
)

type PacketCapture struct {
}

type InetAddr struct {
	IP        net.IP
	Netmask   net.IPMask
	Broadcast net.IP
	Version   string
}

type Nic struct {
	Name      string
	Addresses []InetAddr
}

func NewPcap() *PacketCapture {
	return &PacketCapture{}
}

func (pc *PacketCapture) GetDevs() (ifs []Nic, err error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		logger.Logger.Error("Can't get inet")
		logger.Logger.Error(err)
	}
	for _, device := range devices {
		var nic Nic
		nic.Name = device.Name
		for _, address := range device.Addresses {
			var inet InetAddr
			if address.IP.IsLoopback() {
				break
			} else {
				inet.IP = address.IP
				inet.Broadcast = address.Broadaddr
				inet.Netmask = address.Netmask
				switch _inet.Ip4or6(inet.IP.String()) {
				case _inet.IPv4:
					inet.Version = _inet.IPv4
				case _inet.IPv6:
					inet.Version = _inet.IPv6
				default:
					break
				}

				nic.Addresses = append(nic.Addresses, inet)
			}
		}
		if nic.Addresses != nil {
			ifs = append(ifs, nic)
		}
	}
	//b,err:=json.Marshal((ifs))
	//fmt.Println(string(b))
	return
}
