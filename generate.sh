#!/bin/bash

PROTO_DIR="borealis-prototypes"
REPO_PATH="github.com/aurora-is-near/borealis-prototypes-go"
GEN_PATH="generated"

proto_files=$(find "$PROTO_DIR" -name "*.proto")

declare -a go_opts
for proto_file in $proto_files; do
  proto_path="${proto_file#"$PROTO_DIR"/}"
  dir_path=$(dirname "$proto_path")
  last_dir=$(basename "$dir_path")
  package_name="pb_${last_dir}"
  go_package_path="${REPO_PATH}/${GEN_PATH}/${dir_path}"
  go_opts+=("--go_opt=M${proto_path}=${go_package_path};${package_name}")
  go_opts+=("--go-vtproto_opt=M${proto_path}=${go_package_path};${package_name}")
done

protoc \
	--proto_path="$PROTO_DIR" \
	--go_out=paths=source_relative:$GEN_PATH/. \
	--go-vtproto_out=paths=source_relative:$GEN_PATH/. \
    "${go_opts[@]}" \
    $proto_files
