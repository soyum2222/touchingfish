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
		t.Fail()
	}

	p.mark(begin, end)

	fmt.Println(p.idle(500))


}

func TestMark(t *testing.T) {

	p := BytePool{}
	p.New(1024)

	p.mark(0, 64)

	begin, end := p.idle(64)

	fmt.Println(begin, end)

}
