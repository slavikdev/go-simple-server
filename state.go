/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T23:02:37+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T00:29:33+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import "time"

// State incapsulates app properties which can be saved/restored as state.
type State struct {
	Log []time.Time `json:"log"`
}

// NewState creates new instance of the state.
func NewState(rqlog *RequestLog) *State {
	state := &State{}
	state.Log = rqlog.MinuteHits()
	return state
}
