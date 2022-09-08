package firewall

import (
	"fmt"
	"gitee.com/bytesworld/tomato/internal/logger"
	_inet "gitee.com/bytesworld/tomato/pkg/utils/inet"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
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
		device       string = "eno1"
		err          error
		handle       *pcap.Handle
	)
	handle, err = pcap.OpenLive(device, pc.snaplen, pc.promisc, pc.timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	var filter  = "port 8883"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 8883 packets.")
	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	fmt.Println(packet)
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}

	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}

	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}

	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		fmt.Printf("%s\n", applicationLayer.Payload())

		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}