package internal

import (
	"log"
	"unicode/utf8"
	"unsafe"
)

type Buffer struct {
	address *Buffer
	buf     []byte
}

// Bufferがコピーされているかチェックする
func (b *Buffer) copyCheck() {
	if b.address == nil {
		b.address = (*Buffer)(NoEscape(unsafe.Pointer(b)))
	} else if b.address != b {
		panic("illegal use of non-zero Buffer copied by value")
	}
}

//go:nosplit
//go:nocheckptr
func NoEscape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *Buffer) ToString() string {
	if len(b.buf) == 0 {
        return ""
    }
	return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
}

func (b *Buffer) AllocMem(n int) {
	if n < 0 {
		log.Fatal("negative number is invalid")
	}
	if cap(b.buf)-len(b.buf) < n {
		newCap := len(b.buf) + n  // 既存のデータ長 + 必要な追加容量
		var tempBuf []byte = make([]byte, 0, newCap)
		copy(tempBuf, b.buf)
		b.buf = tempBuf
	}
}

func (b *Buffer) BufLen() int {
	return len(b.buf)
}

func (b *Buffer) BufCap() int {
	return cap(b.buf)
}

func (b *Buffer) AccumuRune(r rune) error {
	b.copyCheck()
	b.buf = utf8.AppendRune(b.buf, r)
	return nil
}

func (b *Buffer) AccumuBytes(p []byte) error {
	b.copyCheck()
	b.buf = append(b.buf, p...)
	return nil
}

func (b *Buffer) BufReset() {
	b.address = nil
	b.buf = nil
}

func (b *Buffer) LeaveCap() {
	if b.buf != nil {
		b.buf = b.buf[:0]
	}
	if b.address != nil {
		b.address = nil
	}
}
