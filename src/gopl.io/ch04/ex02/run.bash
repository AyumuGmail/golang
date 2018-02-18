#!/bin/bash

echo "-t 256"
go run ex02.go -t 256 <<EOF
kataoka
EOF
echo "-t 512"
go run ex02.go -t 512 <<EOF
kataoka
EOF
echo "-t 384"
go run ex02.go -t 384 <<EOF
kataoka
EOF
echo "-t 16"
go run ex02.go -t 16 <<EOF
kataoka
EOF


