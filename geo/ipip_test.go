package geo

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

var ipip, _ = newIpipDB("./ipip.txtx")

func TestIpipDB(t *testing.T) {
	r, e := ipip.FindIP("221.228.97.157")
	fmt.Println("rst>>>", r, e)

	r, e = ipip.FindIP("012.122.111.187")
	fmt.Println("rst>>>", r, e)
	r, e = ipip.FindIP("012.172.161.187")
	fmt.Println("rst>>>", r, e)

	r, e = ipip.FindIP("122.13.132.175")
	fmt.Println("rst>>>", r, e)
}

func BenchmarkIpipDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ipip.FindIP(int2ip(uint32(i)).String())
		//	fmt.Println("rst>>>", r, e)
	}

}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
