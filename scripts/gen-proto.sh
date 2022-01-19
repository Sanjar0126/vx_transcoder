#!/bin/bash
CURRENT_DIR=$(pwd)
echo $CURRENT_DIR
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done

for x in $(find ${CURRENT_DIR}/genproto/* -type d); do
  ls ${x}/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
done
