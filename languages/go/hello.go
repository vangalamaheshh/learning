package main

import ( 
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i int
	max int
	text string
}

func process( j *Job, g *sync.WaitGroup ) {
	for j.i < j.max {
		fmt.Println( j.text )
		time.Sleep( 1 * time.Millisecond )
		j.i++
	}
	g.Done()
}

func main() {
	hello := new( Job )
	world := new( Job )
	goGroup := new( sync.WaitGroup )
	goGroup.Add(2)
	
	hello.i = 0
	hello.max = 2
	hello.text = "hello"

	world.i = 0
	world.max = 5
	world.text = "world"

	go process( hello, goGroup )
	go process( world, goGroup )
	
	goGroup.Wait()
}
