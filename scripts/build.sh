#!/usr/bin/env bash

if [[ -z "$1" ]]; then
    echo "usage: $0 <package-name>"
    exit 1
fi

package=$1

platforms=(
    "android/arm64"
    "darwin/amd64"
    "linux/amd64"
    "windows/amd64"
)

for platform in "${platforms[@]}"
do
    platform=(${platform//\// })
    GOOS=${platform[0]}
    GOARCH=${platform[1]}

    output=$(basename "$package" .go)'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o $output $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred on $GOOS/$GOARCH! Aborting...'
        exit 1
    fi
done