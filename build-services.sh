#!/bin/sh

protoc --go_out=plugins=grpc:service --go_opt=paths=source_relative services.proto
