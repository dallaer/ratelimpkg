package ratelimiter

import (
	"time"
)

type limit struct {
	limit     int // Maximum number of concurrent tasks.
	limit_min int // Maximum number of tasks per minute.
}

var lim limit

// Limitations initialization.
func init() {
	lim.limit = 2
	lim.limit_min = 10
}

// Initialization for custom values of limits.
func Initialization(limit, limit_min int) {
	lim.limit = limit
	lim.limit_min = limit_min
}

func minus(i *int) {
	*i--
}

// Renewal for variable that is responsible for the limit per minute
func renewal(x *int) {
	for {
		time.Sleep(60 * time.Second)
		*x = 0
	}
}

// A function that receives a channel and reads all tasks from the channel
func Ratelimiter(ch chan func()) {
	for {
		select {
		case arg := <-ch:
			go launch(arg) // If there are several tasks in the channel, then  they will be run in different goroutines
		}
	}
}

// When this function receiving a task, it runs it in parallel, depending on the limits
func launch(f func()) {
	i := 0
	x := 0
	go renewal(&x) // start a minute timer to renewal
	for {          // infinite loop to  update tasks, instead of finishing
		if i < lim.limit && x < lim.limit_min { // check for exceeding limits
			i++               // concurrent task counter
			x++               // task counter per minute
			go func(i *int) { // run the task in parallel
				defer minus(i) // when the task finishes the counter will decrease and a new task will run
				f()
			}(&i)
		}
	}
}
