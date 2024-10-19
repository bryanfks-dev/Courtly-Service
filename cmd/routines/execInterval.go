package routines

import "time"

// exeInterval executes a function at a given interval time
//
// timeInterval: the interval at which the function should be executed
// execF: the function to be executed
//
// Retuns void
func execInterval(timeInterval time.Duration, execF func()) {
	for {
		time.Sleep(timeInterval)

		execF()
	}
}
