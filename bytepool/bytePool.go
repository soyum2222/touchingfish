package bytePool

import (
	"errors"
	"math"
)

type BytePool struct {
	b []byte
	m []int64
}

func (b *BytePool) New(size int64) {

	b.b = make([]byte, size)

	if size%64 == 0 {
		b.m = make([]int64, (size / 64))
	} else {
		b.m = make([]int64, (size/64)+1)
	}

}

func (b *BytePool) Malloc(size int64) ([]byte, error) {

	begin, end := b.idle(size)

	if begin == 0 && end == 0 {
		return nil, errors.New("no suitable continuous space")
	}

	b.mark(begin, end)

	return b.b[begin:end], nil

}

//func (b *BytePool) Free(s []byte) error {
//
//	addr1 := (*int64)(unsafe.Pointer(&b.b[0]))
//	addr2 := (*int64)(unsafe.Pointer(&s[0]))
//
//	if *addr1 > *addr2 {
//		return errors.New("not mine")
//	}
//
//	offset := *addr2 - *addr1
//
//}

func (b *BytePool) idle(size int64) (int64, int64) {
	var bit int64

	for i := 0; i < len(b.m); i++ {

		for n := 0; n < 64; n++ {

			if b.m[i]&(1<<(bit%64)) == 0 {

				begin := bit
				end := bit

				for j := i; j < len(b.m); j++ {

				loop:

					if math.Ceil(float64(end)/float64(i+1)) > 64 {
						continue
					}

					if b.m[j]&1<<end%int64(i+1) == 0 {
						end++

						if end-begin >= size {
							return begin, end
						}

						goto loop
					} else {
						break
					}

				}

			} else {
				bit++
			}

		}
	}

	return 0, 0
}

func (b *BytePool) mark(begin, end int64) {

	index := begin / 64

	bit := index % 64

	size := end - begin

	for i := index; i < int64(len(b.m)); i++ {

	loop:

		b.m[i] = b.m[i] | (1 << bit)

		bit++
		size--
		if size == 0 {
			break
		}
		if bit == 64 {
			bit = bit % 64
			continue
		}

		goto loop
	}

}
