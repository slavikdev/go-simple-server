/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T20:15:17+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T00:27:56+03:00
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

// NewRequestLogFromState creates new instance
// of the RequestLog from specified state.
func NewRequestLogFromState(state *State) *RequestLog {
	rqlog := NewRequestLog()
	rqlog.log = state.Log
	return rqlog
}

// Hit adds current time to the log.
func (requestLog *RequestLog) Hit() {
	requestLog.touchLog(func() {
		requestLog.log = append(requestLog.log, time.Now())
	})
}

// MinuteHits returns hits of the last minute. Also removes old hits.
func (requestLog *RequestLog) MinuteHits() (hits []time.Time) {
	requestLog.touchLog(func() {
		// Remove old records.
		for idx, timestamp := range requestLog.log {
			if time.Since(timestamp).Seconds() <= 60 {
				requestLog.log = requestLog.log[idx:]
				break
			}
		}
		hits = requestLog.log
	})
	return
}

// MinuteHitsTotal returns the amount of hits occured in the past 60 sec.
func (requestLog *RequestLog) MinuteHitsTotal() int {
	return len(requestLog.MinuteHits())
}

// touchLog locks state using mutex and performs specific operation on log.
func (requestLog *RequestLog) touchLog(touch func()) {
	requestLog.mutex.Lock()
	touch()
	requestLog.mutex.Unlock()
}
