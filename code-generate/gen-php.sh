#!/usr/bin/env bash

protoDir="../protos"
outDir="../languages/php/lightweight"

protoc --proto_path=${protoDir} \
  --php_out=${outDir} \
  --grpc_out=${outDir} \
  --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin \
  ${protoDir}/*.proto

