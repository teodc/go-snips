package main

import (
	"fmt"
	"math/rand"
)

type Protocol string // type alias (alternate name for an existing type without creating a new distinct type)
const (
	TCP   Protocol = "TCP"
	UDP   Protocol = "UDP"
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
)

type PacketStatus int

const (
	Accepted PacketStatus = iota // automatically generates incrementing values starting from 0
	Rejected
	Suspicious
)

func (ps PacketStatus) String() string {
	// [...] is a slice literal used to create a fixed-length array.
	// It infers the length of the array from the number of elements passed to it.
	return [...]string{"Accepted", "Rejected", "Suspicious"}[ps]
}

type Packet struct {
	ID       int
	Protocol Protocol
	SrcIP    string
	DstIP    string
	Size     int
}

type PacketAnalyzer struct {
	TotalPackets      int
	AcceptedPackets   int
	RejectedPackets   int
	SuspiciousPackets int
}

func (pa *PacketAnalyzer) Analyze(p *Packet) PacketStatus {
	pa.TotalPackets++

	if p.Protocol == HTTP {
		pa.SuspiciousPackets++
		return Suspicious
	}

	if p.Size > 1500 {
		pa.RejectedPackets++
		return Rejected
	}

	pa.AcceptedPackets++

	return Accepted
}

func generatePacket() Packet {
	protocols := []Protocol{TCP, UDP, HTTP, HTTPS}
	return Packet{
		ID:       rand.Intn(10000),
		Protocol: protocols[rand.Intn(len(protocols))],
		SrcIP:    fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		DstIP:    fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		Size:     rand.Intn(3000),
	}
}

func main() {
	analyzer := PacketAnalyzer{}

	for i := 0; i < 100; i++ {
		p := generatePacket()
		status := analyzer.Analyze(&p)
		// %s will format the status as a string,
		// and Status implements the Stringer interface with the String() method.
		fmt.Printf("Packet %s: %#v\n", status, p)
	}

	fmt.Printf("Analysys results: %#v\n", analyzer)
}
