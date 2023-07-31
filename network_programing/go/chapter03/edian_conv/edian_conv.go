package main

import (
	"encoding/binary"
	"fmt"
)

// Go언어는 호스트 바이트 순서를 기본적으로 리틀 엔디안을 사용
// Go에서는 에디안 변환을 위해선 크기에 맞는 바이트 슬라이스를 따로 준비해야 함

func main() {
	// 호스트 바이트 순서 확인
	val := uint16(0x1234)

	if binary.LittleEndian.Uint16([]byte{byte(val), byte(val >> 8)}) == val {
		fmt.Println("Little Endian")
	} else if binary.BigEndian.Uint16([]byte{byte(val >> 8), byte(val)}) == val {
		fmt.Println("Big Endian")
	} else {
		fmt.Println("Unknown")
	}

	hostPort := uint16(0x1234)
	netPort := hostToNetworkUint16(hostPort)
	hostAddr := uint32(0x12345678)
	netAddr := hostToNetworkUint32(hostAddr)

	fmt.Printf("Host ordered port: %#x\n", hostPort)
	fmt.Printf("Network ordered port: %#x\n", netPort)
	fmt.Printf("Host ordered address: %#x\n", hostAddr)
	fmt.Printf("Network ordered address: %#x\n", netAddr)
}

func hostToNetworkUint16(hostPort uint16) uint16 {
	portBuffer := make([]byte, 2)
	binary.LittleEndian.PutUint16(portBuffer, hostPort) // 호스트 바이트 순서(리틀 에디안)로 바이트 슬라이스에 저장
	return binary.BigEndian.Uint16(portBuffer)          // 저장된 바이트 슬라이스에서 네트워크 바이트 순서(빅 에디안)로 리턴
}

func hostToNetworkUint32(hostAddr uint32) uint32 {
	addrBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(addrBuffer, hostAddr)
	return binary.BigEndian.Uint32(addrBuffer)
}
