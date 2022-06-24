package fake

import (
	"math/rand"
	"net"
)

func IPAddress() net.IPAddr {
	return net.IPAddr{
		IP: net.IPv4(byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255))),
	}
}
