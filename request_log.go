/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T20:15:17+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-21T22:50:09+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"sync"
	"time"
)

// RequestLog stores request timestamps for the last 60 seconds.
type RequestLog struct {
	// Collection of timestamps
	log []time.Time

	// Mutex for synchronization between goroutines
	mutex sync.Mutex
}

// NewRequestLog creates new instance of the RequestLog.
func NewRequestLog() *RequestLog {
	return &RequestLog{}
}

// Hit adds current time to the log.
func (requestLog *RequestLog) Hit() {
	requestLog.touchLog(func() {
		requestLog.log = append(requestLog.log, time.Now())
	})
}

// MinuteHitsTotal returns the amount of hits occured in the past 60 sec.
func (requestLog *RequestLog) MinuteHitsTotal() (total int) {
	requestLog.touchLog(func() {
		// Remove old records.
		for idx, timestamp := range requestLog.log {
			if time.Since(timestamp).Seconds() <= 60 {
				requestLog.log = requestLog.log[idx:]
				break
			}
		}
		total = len(requestLog.log)
	})
	return
}

// touchLog locks state using mutex and performs specific operation on log.
func (requestLog *RequestLog) touchLog(touch func()) {
	requestLog.mutex.Lock()
	touch()
	requestLog.mutex.Unlock()
}
