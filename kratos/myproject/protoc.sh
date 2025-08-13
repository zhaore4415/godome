#!/usr/bin/env bash

export goModName="$(grep '^module' go.mod | awk '{print $2}')"
export protoImport="$(dirname $(dirname $(pwd)))"
export protoIn="$(dirname $(dirname $(pwd)))/proto"


# Read proto directories from environment variable
IFS=' ' read -r -a proto_dirs <<< "${PROTO_DIRS}"
export protoDirs=()
for dir in "${proto_dirs[@]}"; do
    protoDirs+=("$protoIn/$dir")
done

echo "✅ Found proto directories:"
echo "  - Import path: $protoImport"
echo "  - Source directory: $protoIn"
echo "  - Go Mod name: $goModName"

echo " Starting protoc compilation..."
for dir in "${protoDirs[@]}"; do
    echo "Current dir: $dir"
    for file in $(find "$dir" -type f -name "*.proto" -not -path '*/.git/*'); do
        echo "  Processing: $file"
       
        if ! protoc --proto_path=.:$protoImport:$protoImport/proto \
            --proto_path=$(go list -f '{{ .Dir }}' -m github.com/envoyproxy/protoc-gen-validate) \
            --go_out=paths=source_relative:./internal \
            --go-grpc_out=paths=source_relative:./internal \
            --validate_out="lang=go,paths=source_relative,register_reflection=true:./internal" \
            --go-http_out=paths=source_relative:./internal $file; then
            echo "❌ Failed to compile $file"
            exit 1
        fi
    done
done

find ./internal/proto -name "*.go" -exec perl -i -pe "s|bsi/proto|$goModName/internal/proto|g" {} \;
# bsi/kratos/panda
find ./internal/proto -name "*.go" -exec perl -i -pe "s|github.com/envoyproxy/protoc-gen-validate/|$goModName/internal/proto/|g" {} \;

#build tag
protoc-go-inject-tag -input="./internal/proto/*/*.pb.go"

echo "✅ All proto files compiled successfully"
exit 0
