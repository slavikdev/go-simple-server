/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T20:40:26+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T02:36:24+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"log"
	"testing"
	"time"
)

func TestHit(t *testing.T) {
	rqlog := NewRequestLog()
	if rqlog.MinuteHitsTotal() > 0 {
		t.Error("Log must be empty at the beginning")
	}
	const hitsTotal = 10
	for i := 0; i < hitsTotal; i++ {
		rqlog.Hit()
	}
	hitsCurrent := rqlog.MinuteHitsTotal()
	assertHits(t, hitsTotal, hitsCurrent)
}

func TestHitsCleanup(t *testing.T) {
	rqlog := NewRequestLog()
	const hitsTotal = 6
	for i := 0; i < hitsTotal; i++ {
		log.Println("HIT")
		rqlog.Hit()
		time.Sleep(10 * time.Second)
	}
	for i := hitsTotal - 1; i > 0; i-- {
		currentTotal := rqlog.MinuteHitsTotal()
		log.Printf("Minute hits total %d", currentTotal)
		assertHits(t, i, currentTotal)
		time.Sleep(10 * time.Second)
	}
}

func TestNewRequestLogFromState(t *testing.T) {
	state := CreateFakeState()
	rqlog := NewRequestLogFromState(state)
	assertHits(t, len(state.Log), rqlog.MinuteHitsTotal())
}

func assertHits(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("Expected minute hits total to equal %d but it was %d.", expected, actual)
	}
}
