package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

type Packet struct {
	Version int64
	ID int64
	Val int64

	//If the length type ID is 0, then the next 15 bits are a number that represents the total length in bits
	//of the sub-packets contained by this packet.
	//If the length type ID is 1, then the next 11 bits are a number that represents the number
	//of sub-packets immediately contained by this packet.
	LengthTypeID int64
	SubLength int64
	SubNumber int64
	SubPackets []*Packet
}

func (p *Packet) SumAllVersions() int64 {
	sum := p.Version
	for _, sub := range p.SubPackets {
		sum += sub.SumAllVersions()
	}
	return sum
}

func (p *Packet) Calculate() int64 {
	switch p.ID {
	case 0:
		sum := int64(0)
		for _, sub := range p.SubPackets {
			sum += sub.Calculate()
		}
		return sum
	case 1:
		sum := int64(1)
		for _, sub := range p.SubPackets {
			sum *= sub.Calculate()
		}
		return sum
	case 2:
		min := int64(math.MaxInt64)
		for _, sub := range p.SubPackets {
			val := sub.Calculate()
			if val < min {
				min = val
			}
		}
		return min
	case 3:
		max := int64(0)
		for _, sub := range p.SubPackets {
			val := sub.Calculate()
			if val > max {
				max = val
			}
		}
		return max
	case 4:
		return p.Val
	case 5:
		if p.SubPackets[0].Calculate() > p.SubPackets[1].Calculate() {
			return 1
		} else {
			return 0
		}
	case 6:
		if p.SubPackets[0].Calculate() < p.SubPackets[1].Calculate() {
			return 1
		} else {
			return 0
		}
	case 7:
		if p.SubPackets[0].Calculate() == p.SubPackets[1].Calculate() {
			return 1
		} else {
			return 0
		}
	default:
		panic(fmt.Errorf("unknown type id %d", p.ID))
	}
}

func parse(filename string) (string, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	binary, err := hexToBin(string(raw))
	if err != nil {
		return "", err
	}
	return binary,nil
}

func binToInt(input string) int64 {
	val, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse binary %s to int64: %w", input, err))
	}
	return val
}

func parsePacket(input string) (*Packet, int64) {
	parseI := int64(0)

	packet := &Packet{}
	packet.Version = binToInt(input[parseI : parseI+3])
	parseI +=3
	packet.ID = binToInt(input[parseI : parseI+3])
	parseI +=3
	fmt.Printf("\n found packet version: %d, id: %d", packet.Version, packet.ID)

	// parse literal value
	if packet.ID == 4 {
		rawVal := ""
		for {
			nextFive := input[parseI:parseI+5]
			parseI +=5
			rawVal += nextFive[1:]
			if string(nextFive[0]) == "0" {
				break
			}
		}
		packet.Val = binToInt(rawVal)
		fmt.Printf(" packet val %d", packet.Val)
	} else {
		packet.LengthTypeID = binToInt(input[parseI : parseI+1])
		parseI++
		fmt.Printf("\n packet length type id %d", packet.LengthTypeID)
		if packet.LengthTypeID == 0 {
			packet.SubLength = binToInt(input[parseI: parseI+15])
			fmt.Printf("\n packet sublength len: %d", packet.SubLength)
			parseI += 15
			subInput := input[parseI:parseI+packet.SubLength]
			parseI += packet.SubLength
			subI := int64(0)
			for subI < packet.SubLength {
				subPacket, i := parsePacket(subInput[subI:])
				packet.SubPackets = append(packet.SubPackets, subPacket)
				subI += i
			}
		} else {
			packet.SubNumber = binToInt(input[parseI:parseI+11])
			parseI += 11
			fmt.Printf("\n packet sub number: %d", packet.SubNumber)
			for s := int64(0); s < packet.SubNumber; s++ {
				subPacket, i := parsePacket(input[parseI:])
				packet.SubPackets = append(packet.SubPackets, subPacket)
				parseI += i
			}
		}
	}

	return packet, parseI
}


func run() error {
	binaryInput, err := parse("day16/input.txt")
	if err != nil {
		return err
	}
	fmt.Printf("\ninput: %s", binaryInput)

	packet, _ := parsePacket(binaryInput)
	//	fmt.Printf("\n versions sum: %d", packet.SumAllVersions())
	fmt.Printf("\n value: %d", packet.Calculate())
	return nil
}

func hexToBin(hex string) (string, error) {
	result := ""
	parseI := 0

	for parseI < len(hex) {
		nextI := parseI + 4
		if nextI >= len(hex) {
			nextI = len(hex)
		}
		parsing := hex[parseI:nextI]
		val, err := strconv.ParseInt(parsing, 16, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse %s: %w", parsing, err)
		}

		diff := nextI - parseI
		template := fmt.Sprintf("%%0%db",diff*4)
		parsed := fmt.Sprintf(template, val)
		result += parsed
		parseI = nextI
	}

	return result, nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
