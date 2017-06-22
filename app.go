/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T20:10:12+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T03:06:07+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import "os"
import "os/signal"
import "syscall"

func main() {
	stateFile := NewStateFile()
	rqlog := NewRequestLogFromState(stateFile.Load())
	server := NewServer(rqlog)
	signalsChannel := make(chan os.Signal, 1)
	signal.Notify(signalsChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalsChannel
		stateFile.Save(NewState(rqlog))
		os.Exit(0)
	}()
	server.Start()
}
