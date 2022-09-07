package inet
const (
	IPv4 = "ipv4"
	IPv6 = "ipv6"
)
func Ip4or6(s string) string {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return IPv4
		case ':':
			return IPv6
		}
	}
	return "unknown"

}
