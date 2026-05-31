package main

import (
	"fmt"
	"sync"
)

type Buffer struct {
	data []byte
}

// ok so over here we are declaring a pool that is thread safe by default , it will be having two methods get and put which can be safely called by multiple go routines
var bufferPool = sync.Pool{
	//the new function is a factory menthod , The pool only excecutes its when someone asks for an item using pool.get() and when pool is totally empty , if pool already has an available object inside it completely skips this function and handle over the recycled object instead
	New: func() interface{} {
		fmt.Println("Allocating new buffer")
		return &Buffer{
			data: make([]byte, 0, 1024),
		}
	},
}

func main() {

	buf := bufferPool.Get().(*Buffer)

	// use it
	buf.data = append(buf.data, []byte("Hello world")...)
	fmt.Println(string(buf.data))

	// reset before putting back
	buf.data = buf.data[:0]

	// put back into pool
	bufferPool.Put(buf)

	// Get again (may reuse)
	buf2 := bufferPool.Get().(*Buffer)
	fmt.Println("Reused buffer length", len(buf2.data))

}
