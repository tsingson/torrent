package storage

import (
	"bytes"
	"context"
	"io"
	"testing"
)

type reader struct {
	b *bytes.Buffer
}

func (r reader) Read(p []byte) (n int, err error) {
	return r.b.Read(p)
}

func TestAlloc(t *testing.T) {
	bufPool := NewBufferPool()

	buff, err := bufPool.Get(context.Background(), 2097152)

	if err != nil {
		t.Fatal(err)
	}

	err = buff.Close()

	if err != nil {
		t.Fatal(err)
	}

	buff, err = bufPool.Get(context.Background(), 2097152)

	if err != nil {
		t.Fatal(err)
	}

	input := [2097152]byte{}
	buff.Write(input[:])

	err = buff.Close()

	if err != nil {
		t.Fatal(err)
	}

	source := reader{bytes.NewBuffer(nil)}

	source.b.Write(input[:])

	buff, err = bufPool.Get(context.Background(), 2097152)

	if err != nil {
		t.Fatal(err)
	}

	cap := buff.Cap()

	io.Copy(buff, source)

	if cap != buff.Cap() {
		t.Fatal("Buffer resized in copy")
	}

	err = buff.Close()

	if err != nil {
		t.Fatal(err)
	}
}

func TestLimitedAlloc(t *testing.T) {
	bufPool := NewLimitedBufferPool(NewBufferPool(), 50_000_000)

	buff, err := bufPool.Get(context.Background(), 2097152)

	if err != nil {
		t.Fatal(err)
	}

	err = buff.Close()

	if err != nil {
		t.Fatal(err)
	}

	buff, err = bufPool.Get(context.Background(), 2097152)

	if err != nil {
		t.Fatal(err)
	}

	input := [2097152]byte{}
	buff.Write(input[:])

	err = buff.Close()

	if err != nil {
		t.Fatal(err)
	}

	source := reader{bytes.NewBuffer(nil)}

	source.b.Write(input[:])

	buff, err = bufPool.Get(context.Background(), 2097152)

	if err != nil {
		t.Fatal(err)
	}

	cap := buff.Cap()

	io.Copy(buff, source)

	if cap != buff.Cap() {
		t.Fatal("Buffer resized in copy")
	}

	err = buff.Close()

	if err != nil {
		t.Fatal(err)
	}
}
