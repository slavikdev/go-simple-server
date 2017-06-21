/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-22T01:04:48+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T01:10:11+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"testing"
)

func TestNewState(t *testing.T) {
	rqlog := NewRequestLog()
	const hitsTotal = 5
	for i := 0; i < hitsTotal; i++ {
		rqlog.Hit()
	}
	state := NewState(rqlog)
	stateLogLen := len(state.Log)
	assertHits(t, hitsTotal, stateLogLen)
}
