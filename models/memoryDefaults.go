package models

import "UASM/neander"

const (
	OneValue uint16 = 255 - iota
	ZeroValue
	MinusOneValue
	AlternateOneValue
	AcCache0Addr
	AcCache1Addr
	SiAddr
)

func GetBuiltinConstants() []uint16 {
	return []uint16{
		SiAddr,
		AcCache1Addr,
		AcCache0Addr,
		AlternateOneValue,
		MinusOneValue,
		ZeroValue,
		OneValue,
	}
}

func GetBuiltinMulFunc(start uint16, arg uint16) []uint16 {
	p1 := []uint16{
		neander.STA, AcCache0Addr,
		neander.LDA, arg,
		neander.STA, SiAddr,
		neander.LDA, ZeroValue,
		neander.STA, AcCache1Addr,
		neander.LDA, MinusOneValue,
		neander.STA, AlternateOneValue,
		neander.LDA, SiAddr,
		neander.NOT,
		neander.ADD, OneValue,
		neander.JN, start + 33, // go @start
		neander.JZ, start + 33, // go @start
		neander.LDA, OneValue,
		neander.STA, AlternateOneValue,
		neander.NOT,
		neander.ADD, OneValue,
	}
	p2 := []uint16{
		// create @start
		neander.JZ, uint16(47), // go @fim,
		neander.LDA, AcCache1Addr,
		neander.ADD, AcCache0Addr,
		neander.STA, AcCache1Addr,
		neander.LDA, SiAddr,
		neander.ADD, AlternateOneValue,
		neander.STA, SiAddr,
		neander.JMP, uint16(int(start) + len(p1) + 1), // go @start,
	}
	p3 := []uint16{
		// create @fim,
		neander.LDA, AcCache1Addr,
	}
	fin := append(p1, p2...)
	return append(fin, p3...)
}

type heapStruct struct {
	content []uint16
	last    int8
}

//var heap = heapStruct{
//    start
//    content: []uint16{},
//    last:    0,
//}
//
//func AlocateNum(value uint16) int8 {
//    if heap == nil {
//        heap = []uint16{}
//    }
//    heap.content[heap.last] = value
//    heap.last++
//    return heap.last - 1
//}
