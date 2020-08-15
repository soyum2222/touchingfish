package bytePool

import (
	"fmt"
	"testing"
)

func TestIdle(t *testing.T) {

	p := BytePool{}

	p.New(1024)

	begin, end := p.idle(500)

	if begin != 0 || end != 500 {
		fmt.Println(begin, end)
		t.Fail()
	}

}

//func TestMark(t *testing.T) {
//
//	p := BytePool{}
//	p.New(1024)
//
//	p.mark(0, 500)
//
//	begin, end := p.idle(100)
//
//	if begin != 500 || end != 600 {
//		fmt.Println(begin, end)
//		t.Fail()
//	}
//
//	fmt.Println(begin, end)
//}

func TestFree(t *testing.T) {
	p := BytePool{}
	p.New(1024)

	b, err := p.Malloc(500)
	if err != nil {
		panic(err)
	}

	for k := range b {
		b[k] = uint8(k)
	}

	if p.Free(b) != nil {
		panic(err)
	}

	b, err = p.Malloc(500)
	if err != nil {
		panic(err)
	}

	for k := range b {
		if b[k] != uint8(k) {
			t.Fail()
		}
	}

}

func BenchmarkBytePool(b *testing.B) {
	pool := BytePool{}
	pool.New(1 << 20)

	for i := 0; i < b.N; i++ {

		b, err := pool.Malloc(1 << 10)
		if err != nil {
			panic(err)
		}

		donothing(b)
		//
		//c, err := pool.Malloc(1 << 10)
		//if err != nil {
		//	panic(err)
		//}

		//donothing(c)

		if err = pool.Free(b); err != nil {
			panic(err)
		}

		//if err = pool.Free(c); err != nil {
		//	panic(err)
		//}
	}

}

func BenchmarkNewByte(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b := make([]byte, 1<<10)
		donothing(b)
	}

}

func donothing(b []byte) {
	for k := range b {
		b[k] = uint8(k)
	}
}
