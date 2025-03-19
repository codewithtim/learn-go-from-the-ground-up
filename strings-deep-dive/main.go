package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {

	s := "hello, 世界"
	runes := []rune(s)
	fmt.Println(len(s))     // 13 (bytes)
	fmt.Println(len(runes)) // 9 (characters)

	// Now you can access characters by index
	fmt.Println(string(runes[7])) // 世

	// Convert to rune slice for random access:
	fmt.Println(string(runes[7]))
	// Use the utf8 package for manual decoding
	r, size := utf8.DecodeRuneInString(s[7:])
	fmt.Printf("%c %d\n", r, size)

	// Use the range loop for sequential access:
	for i, r := range s {
		// i is the byte index, r is the rune value
		fmt.Printf("bye position: %2d, rune: %c, unicode: %U, bytes: % X\n", i, r, r, string(r))
	}

	s1 := "hello, world"
	s2 := s1[0:5]
	p1 := unsafe.StringData(s1)
	p2 := unsafe.StringData(s2)

	fmt.Printf("s1 data pointer: %p\n", p1)
	fmt.Printf("s2 data pointer: %p\n", p2)

	// String literals share the same memory if identical
	hello1 := "hello"
	hello2 := "hello"

	// Get pointers to the underlying byte sequences
	p1 = unsafe.StringData(hello1)
	p2 = unsafe.StringData(hello2)

	fmt.Printf("hello1 data pointer: %p\n", p1)
	fmt.Printf("hello2 data pointer: %p\n", p2)
	// (Same address! The data is shared)

	// Now create a new string through concatenation
	helloWorld := hello1 + " world"
	p3 := unsafe.StringData(helloWorld)

	fmt.Printf("helloWorld data pointer: %p\n", p3)
	// (Different address - new memory allocation)

	// We can also examine what happens with a substring
	hello3 := helloWorld[0:5] // Should be "hello" again
	p4 := unsafe.StringData(hello3)

	fmt.Printf("hello3 (substring) data pointer: %p\n", p4)
	fmt.Printf("helloWorld data pointer: %p\n", p3)
	// Output:
	// (Same starting address! The substring shares memory with helloWorld)

	// We can also check the length
	fmt.Printf("helloWorld length: %d\n", len(helloWorld)) // 11
	fmt.Printf("hello3 length: %d\n", len(hello3))         // 5
}
