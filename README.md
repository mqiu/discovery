# Stream Reader

## Description

This is a http service that will allow query over streamID value.
Currently only HTTP GET method is accepted.

Based on the input of stream ID, it will return following case:

* Input of invalid id, then 4xx error will be returned

* Input of valid id, but data parser is wrong, then 5xx error will be returned

* Input of valid id, and no error processing data, then json data of the stream record will be returned.

## How to Run

Please go to the `/cmd` directory, and run `go run main.go`

_Note:_ Currently no module initiation like `go mod init` as it really depending on the
project owner's decision on how to handle dependency. Currently it only users the standard library and a gorilla mux

## Storage layer

For simplicity, the data is no db layer. The `Storage` layer is implemented by a simple file reader which will read the whole records into the memory.

I manually input the file `data.json` in `/cmd/` directory.
Any data can be updated into the file and it will be considered as a data storage.

## Unit Testing

Didn't implemet unit testing at the moment due to time constraint. Should always en-forcing the unit testing for all of the exported
functions later.
