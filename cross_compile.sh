#! /usr/bin/bash

platforms=(
  "windows/amd64"
  "darwin/amd64"
  "linux/amd64"
)

# Get name of output directory
package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  output_name="${package}_${GOOS}_${GOARCH}"

  if [ $GOOS = "windows" ]; then
    output_name+='.exe'
  fi

  # execute build
  env GOOS=$GOOS GOARCH=$GOARCH go build -o ./bin/$output_name
done