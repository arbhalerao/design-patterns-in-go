#!/bin/bash

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <category> <pattern>"
  exit 1
fi

CATEGORY=$1
PATTERN=$2

PATTERN_LOWER=$(echo "$PATTERN" | tr '[:upper:]' '[:lower:]')
CATEGORY_LOWER=$(echo "$CATEGORY" | tr '[:upper:]' '[:lower:]')

# Create directory structure
mkdir -p "$CATEGORY_LOWER/$PATTERN_LOWER/cmd"

# Create cmd/main.go with package main
cat <<EOF > "$CATEGORY_LOWER/$PATTERN_LOWER/cmd/main.go"
package main

func main() {
	
}
EOF

# Create pattern.go with the pattern's package name
cat <<EOF > "$CATEGORY_LOWER/$PATTERN_LOWER/$PATTERN_LOWER.go"
package $PATTERN_LOWER
EOF

echo "$PATTERN pattern structure created successfully in $CATEGORY_LOWER/$PATTERN_LOWER"
