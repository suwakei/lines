package internal

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"
)

func check(t *testing.T, b *Buffer, want string) {
	t.Helper()
	got := b.ToString()
	if got != want {
		t.Errorf("String: got %#q; want %#q", got, want)
		return
	}
	if n := b.BufLen(); n != len(got) {
		t.Errorf("Len: got %d; but len(String()) is %d", n, len(got))
	}
	if n := b.BufCap(); n < len(got) {
		t.Errorf("Cap: got %d; but len(String()) is %d", n, len(got))
	}
}

func TestBufferBufReset(t *testing.T) {
	var b Buffer
	check(t, &b, "")
	b.AccumuRune('a')
	s := b.ToString()
	check(t, &b, "a")
	b.BufReset()
	check(t, &b, "")

	// Ensure that writing after Reset doesn't alter
	// previously returned strings.
	b.AccumuRune('b')
	check(t, &b, "b")
	if want := "a"; s != want {
		t.Errorf("previous String result changed after Reset: got %q; want %q", s, want)
	}
}

func TestBufferAllocMem(t *testing.T) {
	for _, growLen := range []int{0, 100, 1000, 10000, 100000} {
		p := bytes.Repeat([]byte{'a'}, growLen)
		allocs := testing.AllocsPerRun(100, func() {
			var b Buffer
			b.AllocMem(growLen) // should be only alloc, when growLen > 0
			if b.BufCap() < growLen {
				t.Fatalf("growLen=%d: Cap() is lower than growLen", growLen)
			}
			b.AccumuBytes(p)
			if b.ToString() != string(p) {
				t.Fatalf("growLen=%d: bad data written after Grow", growLen)
			}
		})
		wantAllocs := 1
		if growLen == 0 {
			wantAllocs = 0
		}
		if g, w := int(allocs), wantAllocs; g != w {
			t.Errorf("growLen=%d: got %d allocs during Write; want %v", growLen, g, w)
		}
	}
	// when growLen < 0, should panic
	var a Buffer
	n := -1
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("a.Grow(%d) should panic()", n)
		}
	}()
	a.AllocMem(n)
}

func TestBufferAccumuMultiBytesChar(t *testing.T) {
	const s0 = "hello 世界"
	for _, tt := range []struct {
		name string
		fn   func(b *Buffer) (error)
		want string
	}{
		{
			"AccumuBytes",
			func(b *Buffer) (error) { return b.AccumuBytes([]byte(s0)) },
			s0,
		},
		{
			"AccumuRune",
			func(b *Buffer) (error) { return b.AccumuRune('a') },
			"a",
		},
		{
			"AccumuRuneWide",
			func(b *Buffer) (error) { return b.AccumuRune('世') },
			"世",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var b Buffer
			err := tt.fn(&b)
			if err != nil {
				t.Fatalf("first call: got %s", err)
			}
			check(t, &b, tt.want)

			err = tt.fn(&b)
			if err != nil {
				t.Fatalf("second call: got %s", err)
			}
			check(t, &b, tt.want+tt.want)
		})
	}
}

func TestBufferAllocs(t *testing.T) {
	// Buffer allocated to heap by escaping
	n := testing.AllocsPerRun(10000, func() {
		var b Buffer
		arr := []rune{'a', 'b', 'c', 'd', 'e'}
		b.AllocMem(5)
		for _, ar := range arr {
			b.AccumuRune(rune(ar))
		}
		_ = b.ToString()
	})
	if n != 1 {
		t.Errorf("Builder allocs = %v; want 1", n)
	}
}

func TestBufferCopyPanic(t *testing.T) {
	tests := []struct {
		name      string
		fn        func()
		wantPanic bool
	}{
		{
			name:      "BufLen",
			wantPanic: false,
			fn: func() {
				var a Buffer
				a.AccumuRune('x')
				b := a
				b.BufLen()
			},
		},
		{
			name:      "BufCap",
			wantPanic: false,
			fn: func() {
				var a Buffer
				a.AccumuRune('x')
				b := a
				b.BufCap()
			},
		},
		{
			name:      "BufReset",
			wantPanic: false,
			fn: func() {
				var a Buffer
				a.AccumuRune('x')
				b := a
				b.BufReset()
				b.AccumuRune('y')
			},
		},
		{
			name:      "AccumuBytes",
			wantPanic: true,
			fn: func() {
				var a Buffer
				a.AccumuBytes([]byte("x"))
				b := a
				b.AccumuBytes([]byte("y"))
			},
		},
		{
			name:      "AccumuRune",
			wantPanic: true,
			fn: func() {
				var a Buffer
				a.AccumuRune('x')
				b := a
				b.AccumuRune('y')
			},
		},
		{
			name:      "AllocMem",
			wantPanic: true,
			fn: func() {
				var a Buffer
				a.AllocMem(1)
				b := a
				b.AllocMem(2)
			},
		},
	}
	for _, tt := range tests {
		didPanic := make(chan bool)
		go func() {
			defer func() { didPanic <- recover() != nil }()
			tt.fn()
		}()
		if got := <-didPanic; got != tt.wantPanic {
			t.Errorf("%s: panicked = %v; want %v", tt.name, got, tt.wantPanic)
		}
	}
}

func TestBufferAccumuInvalidRune(t *testing.T) {
	// Invalid runes, including negative ones, should be written as
	// utf8.RuneError.
	for _, r := range []rune{-1, utf8.MaxRune + 1} {
		var b Buffer
		b.AccumuRune(r)
		check(t, &b, "\uFFFD")
	}
}

var someBytes = []byte("some bytes asdfghjkl")

var sinkS string

func benchmarkBuffer(b *testing.B, f func(b *testing.B, numWrite int, grow bool)) {
	b.Run("1Write_NoGrow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 1, false)
	})
	b.Run("3Write_NoGrow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, false)
	})
	b.Run("3Write_Grow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, true)
	})
}

func BenchmarkBuildString_Buffer(b *testing.B) {
	benchmarkBuffer(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf Buffer
			if grow {
				buf.AllocMem(len(someBytes) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.AccumuBytes(someBytes)
			}
			sinkS = buf.ToString()
		}
	})
}

func BenchmarkBufferString_ByteBuffer(b *testing.B) {
	benchmarkBuffer(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			if grow {
				buf.Grow(len(someBytes) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.Write(someBytes)
			}
			sinkS = buf.String()
		}
	})
}

func TestBufferAllocMemSizeclasses(t *testing.T) {
	ss := strings.Repeat("a", 19)
	allocs := testing.AllocsPerRun(100, func() {
		var b Buffer
		b.AllocMem(18)
		for _, s := range ss {
			b.AccumuRune(s)
		}
		_ = b.ToString()
	})
	if allocs > 1 {
		t.Fatalf("unexpected amount of allocations: %v, want: 1", allocs)
	}
}
