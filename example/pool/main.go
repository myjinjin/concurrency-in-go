package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// create pool of bytes.Buffers which can be reused.
var bufPool = sync.Pool{
	New: func() any {
		fmt.Println("allocating new bytes.Buffer.")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	var b = bufPool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
	bufPool.Put(b) // 반환 -> 반환을 해야 나중에 재활용할 수 있다
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}

/*
 go run .
allocating new bytes.Buffer.
20:34:15 : debug-string1
20:34:15 : debug-string2
*/
