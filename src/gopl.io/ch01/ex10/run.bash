#!/bin/bash

go build ex10.go
rm -R tmp/*
./ex10 https://golang.org http://gopl.io https://godoc.org http://archive.oreilly.com/oreillyschool/courses/advancedjavascript/Advanced%20JavaScript%20Essentials%20v1.pdf
./ex10 https://golang.org http://gopl.io https://godoc.org http://archive.oreilly.com/oreillyschool/courses/advancedjavascript/Advanced%20JavaScript%20Essentials%20v1.pdf


echo "ls outputfiles"
ls -l tmp/
echo "diff outputfiles"
diff tmp/0.out tmp/0.out.add
diff tmp/1.out tmp/1.out.add
diff tmp/2.out tmp/2.out.add
diff tmp/3.out tmp/3.out.add
