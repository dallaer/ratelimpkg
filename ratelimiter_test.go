package ratelimiter

import (
	"math/rand"
	"testing"
	"time"
)

// Initialization tests on random values
func TestInitialization(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	var x, y int

	for i := 0; i <= 1000; i++ {
		x = rand.Intn(1000)
		y = rand.Intn(1000)
		Initialization(x, y)
		if lim.limit != x || lim.limit_min != y {
			t.Errorf("Incorrect result, Expect %d %d, got %d %d", x, y, lim.limit, lim.limit_min)
		}
	}
}

// A function that will help to count the number of goroutine launches, in special cases
//(when running same functions, when working with a package, security is not guaranteed)
func test(x *int) {
	*x++
}

// Test for limit values when threads should not run
func TestRatelimiter(t *testing.T) {
	Initialization(0, 0)
	x := 0
	go launch(func() { test(&x) })
	time.Sleep(1 * time.Second)
	if x != 0 {
		t.Errorf("Incorrect result, Expect %d, got %d", 0, x)
	}
	Initialization(0, 1)
	x = 0
	go launch(func() { test(&x) })
	time.Sleep(1 * time.Second)
	if x != 0 {
		t.Errorf("Incorrect result, Expect %d, got %d", 0, x)
	}
	Initialization(1, 0)
	x = 0
	go launch(func() { test(&x) })
	time.Sleep(1 * time.Second)
	if x != 0 {
		t.Errorf("Incorrect result, Expect %d, got %d", 0, x)
	}
}

// Limit per minute test
func TestRatelimiter2(t *testing.T) {
	for n := 0; n <= 10; n++ {
		Initialization(1, n)
		x := 0
		go launch(func() { test(&x) })
		time.Sleep(1 * time.Second)
		if x != n {
			t.Errorf("Incorrect result, Expect %d, got %d", n, x)
		}
	}
}

// Test for channels with different length
func TestRatelimiter3(t *testing.T) {
	Initialization(1, 1)
	for n := 1; n < 11; n++ {
		x := 0
		ch := make(chan func(), n)
		for i := 0; i < n; i++ {

			ch <- func() { test(&x) }
		}
		go Ratelimiter(ch)
		time.Sleep(1 * time.Second)
		if x != n {
			t.Errorf("Incorrect result, Expect %d, got %d", n, x)
		}
	}
}
