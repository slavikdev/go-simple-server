/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T03:43:45+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T00:18:51+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// StateFile is a file which contains application state.
type StateFile struct{}

const stateFileName = "/tmp/go-simple-server-state.json"

// NewStateFile initializes new instance of the state file.
func NewStateFile() *StateFile {
	return &StateFile{}
}

// Save saves request log to the state file.
func (stateFile *StateFile) Save(state *State) {
	json, err := json.Marshal(state)
	if err != nil {
		log.Fatalf("Invalid JSON state object: %+v. Error: %s", state, err.Error())
	} else {
		err = ioutil.WriteFile(stateFileName, json, 0644)
		if err != nil {
			log.Fatalf("Couldn't save state: %s", err.Error())
		}
	}
}

// Load loads state from file.
func (stateFile *StateFile) Load() *State {
	state := &State{}
	data, err := ioutil.ReadFile(stateFileName)
	if err != nil {
		log.Printf("State file isn't available. Working on defaults. Reason: %s", err.Error())
	} else {
		err = json.Unmarshal(data, state)
		if err != nil {
			log.Printf("State file couldn't be parsed. Working on defaults. Reason: %s", err.Error())
		}
	}
	return state
}
