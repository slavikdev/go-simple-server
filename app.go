/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-21T20:10:12+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T02:39:37+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

func main() {
	stateFile := NewStateFile()
	rqlog := NewRequestLogFromState(stateFile.Load())
	server := NewServer(rqlog)
	defer stateFile.Save(NewState(rqlog))
	server.Start()
}
