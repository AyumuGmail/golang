#!/bin/bash

go build ex03.go
./ex03 test1 test2
./ex03


echo "Bench nomarl vs strings.Join"
go test -bench=.
