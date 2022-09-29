package firewall

import (
	"fmt"
	"gitee.com/bytesworld/tomato/internal/logger"
	_inet "gitee.com/bytesworld/tomato/pkg/utils/inet"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

type PacketCapture struct {
	snaplen int32
	promisc bool
	timeout time.Duration
}

func NewPcap() *PacketCapture {
	// 配置和之前python抓取数据一致
	return &PacketCapture{
		snaplen: 65535,
		promisc: false,
		timeout: 5 * time.Millisecond,
	}
}

func (pc *PacketCapture) GetDevs() (ifs []_inet.Nic, err error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		logger.Logger.Error("Can't get inet")
		logger.Logger.Error(err)
	}
	for _, device := range devices {
		var nic _inet.Nic
		nic.Name = device.Name
		for _, address := range device.Addresses {
			var inet _inet.InetAddr
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

func (pc *PacketCapture) GetPackets() {
	var (
		device string = "en0"
		err    error
		handle *pcap.Handle
	)
	handle, err = pcap.OpenLive(device, pc.snaplen, pc.promisc, pc.timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	//var filter  = "port 8883"
	//err = handle.SetBPFFilter(filter)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Only capturing TCP port 8883 packets.")
	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, ok := ethernetLayer.(*layers.Ethernet)
		if ok {
			if ethernetPacket.EthernetType == layers.EthernetTypeIPv6 {
				fmt.Println("暂不支持ipv6")
			} else if ethernetPacket.EthernetType == layers.EthernetTypeIPv4 {
				fmt.Println("ipv4进入")
				ipLayer := packet.Layer(layers.LayerTypeIPv4)
				ip, _ := ipLayer.(*layers.IPv4)
				SrcIp := ip.SrcIP
				DstIp := ip.DstIP
				_protocol := ip.Protocol
				if _protocol == layers.IPProtocolUDP {
					udpLayer := packet.Layer(layers.LayerTypeUDP)
					udp, _ := udpLayer.(*layers.UDP)
					SrcPort := udp.SrcPort
					DestPort := udp.DstPort
					fmt.Printf("Src %v:%v udp->Dst %v:%v\n", SrcIp, SrcPort, DstIp, DestPort)
				} else if _protocol == layers.IPProtocolTCP {
					tcpLayer := packet.Layer(layers.LayerTypeTCP)
					tcp, _ := tcpLayer.(*layers.TCP)
					SrcPort := tcp.SrcPort
					DestPort := tcp.DstPort
					fmt.Printf("Src %v:%v tcp->Dst %v:%v\n", SrcIp, SrcPort, DstIp, DestPort)
				} else if _protocol == layers.IPProtocolICMPv4 {
					IcmpLayer := packet.Layer(layers.LayerTypeICMPv4)
					_, _ = IcmpLayer.(*layers.ICMPv4)
					fmt.Printf("Src %v icmp->Dst :%v\n", SrcIp, DstIp)
				} else {
					fmt.Println("无需支持的protocol")
				}

			}
		}
	}
}
