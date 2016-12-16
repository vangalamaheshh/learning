package main

import (
	"fmt"
	"runtime"
	"time"
	"strconv"
	"sync"
)

func showNumber( i int, g *sync.WaitGroup ) {
	timestamp := strconv.FormatInt( time.Now().UnixNano(), 10 )
	fmt.Println( i, timestamp )
	time.Sleep( time.Millisecond * 10 )
	g.Done()
}

func main() {
	runtime.GOMAXPROCS( 2 )
	goGroup := new( sync.WaitGroup )
	num := 10
	for i := 0; i <= num; i++ {
		go showNumber( i, goGroup )
	}

	fmt.Println( "Goodbye!" )
	goGroup.Add( 10 )
	goGroup.Wait()
}
