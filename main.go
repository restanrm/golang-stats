package stats

import (
	"fmt"
	"sync"
) // Debugging
//import "strconv"
import "time"

//
// Our data structure holding key/value pairs
//
var data map[string]int = make(map[string]int)
var m sync.RWMutex

/**
* Increment a stat.
 */
func IncrStat(key string) {
	AddStat(key, 1)
}

/**
* Increment a key by a specified value.
 */
func AddStat(key string, value int) {
	m.Lock()
	defer m.Unlock()
	data[key] += value
}

/**
* Decrement a stat.
 */
func DecrStat(key string) {
	SubStat(key, 1)
}

/**
* Decrement a key by a specified value.
 */
func SubStat(key string, value int) {
	m.Lock()
	defer m.Unlock()
	data[key] -= value
}

/**
* Grab the value of a specific key.
 */
func Stat(key string) int {
	m.RLock()
	defer m.RUnlock()
	return (data[key])
}

/**
* Get stats for all keys.
 */
func StatAll() map[string]int {
	return (data)
}

/**
* Fire our callback every specified interval, presumably to print out
* our stats (or dump them to a database or whatever).
*
* @param {float64} interval How many seconds between runs
* @param {func} cb The function to call.
*
 */
func StatDumpFunc(interval float64, cb func(data map[string]int)) {

	seconds_string := fmt.Sprintf("%f", interval)
	duration, _ := time.ParseDuration(seconds_string + "s")

	for {
		time.Sleep(duration)
		m.RLock()
		cb(data)
		m.RUnlock()
	}

} // End of StatsDumpFunc()

/**
* Dump stats periodically.  This is a wrapper for StatsDumpFunc() with
* a built in callback to print to stdout.
*
* @param {float64} interval How many seconds between runs
*
 */
func StatDump(interval float64) {

	StatDumpFunc(interval, func(data map[string]int) {
		fmt.Println("StatsDump():", data)
	})

} // End of StatsDump()
