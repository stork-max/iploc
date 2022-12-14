package dict

import (
    "net"
)

type IP struct {
    net.IP
}

func (ip *IP) ToPath() [8]byte{
    return [8]byte{
        ip.IP[0] >> 4,
        ip.IP[0] << 4 >> 4,
        ip.IP[1] >> 4,
        ip.IP[1] << 4 >> 4,
        ip.IP[2] >> 4,
        ip.IP[2] << 4 >> 4,
        ip.IP[3] >> 4,
        ip.IP[3] << 4 >> 4,
    }
}

func (ip *IP) ToPath4() [4]byte{
    return [4]byte{
        ip.IP[0],
        ip.IP[1],
        ip.IP[2],
        ip.IP[3],
    }
}

func (ip *IP) ToUint32() uint32{
    sum := uint32(0)
    sum += uint32(ip.IP[0]) << 24
    sum += uint32(ip.IP[1]) << 16
    sum += uint32(ip.IP[2]) << 8
    sum += uint32(ip.IP[3])
    return sum
}

func (ip *IP) Len() int {
    return len(ip.IP)
}

func NewUint32IP(l uint32) (IP) {
    return IP{ []byte{
            byte(l>>24),
            byte(l>>16),
            byte(l>>8),
            byte(l),
        } }
}

func NewStringIP(s string) (IP) {
    ip := IP{}
    ip.IP = net.ParseIP(s).To4()
    return ip
}

func NewBytesIP(b [8]byte) (IP) {
    return IP{ []byte{
        b[0] << 4 | b[1],
        b[2] << 4 | b[3],
        b[4] << 4 | b[5],
        b[6] << 4 | b[7],
    } }
}

