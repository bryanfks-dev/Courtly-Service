package routines

import "time"

// Run is the entry point of the routines package
//
// Returns void
func Run() {
	// Run task in every 24h intervals
	go execInterval(24*time.Hour, func() {
		// Run the delete blacklist token routine
		go runClearBlacklistedToken()
	})
}
