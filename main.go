package main

import (
	"fmt"
)

type ReadBuffer struct {
	buff []byte
	num  int
}

func (ra *ReadBuffer) Read() []byte {
	fmt.Println("Metodo invocado")

	return ra.buff
}

func (ra *ReadBuffer) buffering() {
	fmt.Println("Metodo invocado")
}

func ReadAll(r Reader) []byte {
	mabuf := r.Read()
	return mabuf
}

// Interface
type Reader interface {
	Read() []byte
	//buffering()
}

func main() {
	buf := NewReader("Hola mundo yo soy anon que tal")
	mabuf := ReadAll(buf)

	fmt.Printf("%s", mabuf)
}

func NewReader(pattern string) *ReadBuffer {
	patToByte := []byte(pattern)
	return &ReadBuffer{buff: patToByte, num: len(patToByte)}
}
