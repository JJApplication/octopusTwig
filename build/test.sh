#!/usr/bin/env bash

go test -bench=.  -benchmem -benchtime=5s

go test -v .