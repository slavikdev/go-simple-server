# Go HTTP Server Example
HTTP server which on each request responds with a counter of the total number of requests that it has received during the previous 60 seconds (moving window).
The server continues to return correct numbers after restart, by persisting data to a file.

[![Build Status](https://travis-ci.org/slavikdev/go-simple-server.svg?branch=master)](https://travis-ci.org/slavikdev/go-simple-server)
