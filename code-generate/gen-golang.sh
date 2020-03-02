#!/usr/bin/env bash

protoDir="../protos"
outDir="../languages/golang/lightweight"
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}

