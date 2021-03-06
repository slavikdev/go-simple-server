/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-22T00:37:49+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T01:10:47+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func CreateFakeState() *State {
	log := []time.Time{time.Now(), time.Now(), time.Now()}
	return &State{Log: log}
}

func TestLoadNoFile(t *testing.T) {
	os.Remove(stateFileName)
	stateFile := NewStateFile()
	state := stateFile.Load()
	assertDefaultState(t, state)
}

func TestLoadBadFile(t *testing.T) {
	os.Remove(stateFileName)
	ioutil.WriteFile(stateFileName, []byte("bad-json"), 0644)
	stateFile := NewStateFile()
	state := stateFile.Load()
	assertDefaultState(t, state)
}

func TestSave(t *testing.T) {
	stateFile := NewStateFile()
	state := CreateFakeState()
	stateFile.Save(state)
	loadedState := stateFile.Load()
	loadedLen := len(loadedState.Log)
	originalLen := len(state.Log)
	assertHits(t, originalLen, loadedLen)
	for idx, originalTime := range state.Log {
		if originalTime.Sub(loadedState.Log[idx]).Seconds() > 1.0 {
			t.Errorf("Expected timestamp %v to match %v.", originalTime, loadedState.Log[idx])
		}
	}
}

func assertDefaultState(t *testing.T, state *State) {
	logLen := len(state.Log)
	if logLen != 0 {
		t.Errorf("Expected default log length to be 0, but it was %d.", logLen)
	}
}
