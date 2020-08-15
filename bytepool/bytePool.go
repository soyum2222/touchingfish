package bytePool

import (
	"errors"
	"math"
	"unsafe"
)

type BytePool struct {
	b []byte
	m []uint64
}

func (b *BytePool) New(size int64) {

	b.b = make([]byte, size)

	b.m = make([]uint64, int(math.Ceil(float64(size)/float64(64))))

}

func (b *BytePool) Malloc(size int64) ([]byte, error) {

	begin, end := b.idle(size)

	if begin == 0 && end == 0 {
		return nil, errors.New("no suitable continuous space")
	}

	if end-begin != size {
		return nil, errors.New("bad idle")
	}

	b.mark(begin, end)

	return b.b[begin:end], nil

}

func (b *BytePool) Free(s []byte) error {

	addr1 := (uintptr)(unsafe.Pointer(&b.b[0]))

	addr2 := (uintptr)(unsafe.Pointer(&s[0]))

	if addr1 > addr2 {
		return errors.New("not mine")
	}

	offset := int(addr2 - addr1)

	if offset > len(b.b) {
		return errors.New("not mine")
	}

	end := offset + len(s)

	bit := offset % 64

	for i := int(math.Ceil(float64(offset) / 64)); i < len(b.m); i++ {

		for {

			if bit%64 == 0 {

				if end-bit >= 64 {
					b.m[i] = 0
					bit += 64
					break
				}

				b.m[i] = b.m[i] &^ ((1 << ((end - bit) + 1)) - 1)
				return nil

			}

			b.m[i] = b.m[i] & (^(1 << (bit % 64)))

			bit++

			if bit >= end {
				return nil
			}

			if bit%64 == 0 {
				break
			}
		}
	}

	return nil

}

func (b *BytePool) idle(size int64) (int64, int64) {
	var bit int64

	for i := 0; i < len(b.m); i++ {

		if b.m[i] == (1<<64)-1 {
			bit += 64
			continue
		}

		for n := 0; n < 64; n++ {

			if b.m[i]&(1<<(bit%64)) == 0 {

				begin := bit
				end := bit

				for j := i; j < len(b.m); j++ {

					// 64bit batch and

					// var end modulo to 64 == 0 ,bit pointer direction the b.m slice boundary
					if end%64 == 0 {

						remain := (size + begin) - end
						if remain >= 64 && b.m[j]&((1<<64)-1) == 0 {
							end += 64
							continue
						}

						if remain < 64 {

							if b.m[j]&((1<<(remain+1))-1) == 0 {
								end += remain
								return begin, end
							}
							bit = end
							i = j
							break

						}

					}

					// bit pointer not direction b.m slice boundary,then continuous offset pointer
				loop:

					if end/int64(j+1) == 64 {
						continue
					}

					if b.m[j]&(1<<(end%64)) == 0 {

						end++

						if end-begin >= size {
							return begin, end
						}

						goto loop
					} else {
						bit = end
						i = j
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

	var bit = begin - (index * 64)

	size := end - begin

	for i := index; i < int64(len(b.m)); i++ {

	loop:

		if bit == 0 {

			if size > 64 {
				b.m[i] = (1 << 64) - 1
				size -= 64
				continue
			}

			b.m[i] = (1 << (size)) - 1
			return
		}

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
