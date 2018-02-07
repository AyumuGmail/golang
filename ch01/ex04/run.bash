#!/bin/bash

go build ex04.go


echo "expected STDIN"
./ex04 <<EOF
test
test
EOF


echo "expected No duplicatedLine file"
./ex04 <<EOF
test
test2
EOF

