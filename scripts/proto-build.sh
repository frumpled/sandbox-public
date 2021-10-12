#!/bin/sh

protoc --twirp_out=. --go_out=. twirp/proto/*.proto