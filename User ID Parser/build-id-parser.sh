#!/usr/bin/env bash

go mod download

package="./IdParser.go"

package_split=(${package//\// })
package_name="IdParser"

platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386" "linux/arm64" "linux/arm")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ "$GOOS" = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS="$GOOS" GOARCH=$GOARCH go build -o "./build/$output_name" $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done