# ratelimpkg
# Overview
Ratelimpkg is used to run tasks in parallel.
# Constants
This section is empty.
# Variables
This section is empty.
# Functions
func Initialization(limit int, limit_min int)

func Ratelimiter(ch chan func())
# Initialization
Receive two values, the first is a limit on the number of parallel tasks, the second is on the number of tasks per minute. Default values are 2 and 10
# Ratelimiter
Receive a func() channel. All tasks in the channel will be launched, the limits are calculated for each separately. Running tasks that change the values of global variables in parallel is not safe
    
