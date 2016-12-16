package main

import (
	"fmt"
	"time"
	"runtime"
	"io/ioutil"
	"sync"
)

type Job struct {
	i int
	max int
	text string
}

func process( j *Job, g *sync.WaitGroup ) {
	filename := j.text + ".txt"
	filecontents := ""
	for j.i < j.max {
		fmt.Println( j.text )
		time.Sleep( 1 * time.Millisecond )
		filecontents += j.text
		j.i++
	}
	
	err := ioutil.WriteFile( filename, []byte( filecontents ), 0644 )
	
	if( err != nil ) {
		panic( "something went wrong" )
	}
	defer g.Done()
}

func main() {

	runtime.GOMAXPROCS( 2 )
	goGroup := new( sync.WaitGroup )
	hello := new( Job )
	world := new( Job )

	hello.i = 0
	hello.max = 2
	hello.text = "hello"

	world.i = 0
	world.max = 5
	world.text = "world"

	go process( hello, goGroup )
	go process( world, goGroup )

	goGroup.Add( 2 )
	goGroup.Wait()
}
