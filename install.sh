#!/bin/bash

echo "Installing Go tools and dependencies..."
go install golang.org/x/lint/golint@latest
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install github.com/conventionalcommit/commitlint@latest
go install mvdan.cc/gofumpt@latest
go install github.com/mrtazz/checkmake@latest
go install github.com/editorconfig-checker/editorconfig-checker@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/yoheimuta/protolint/cmd/protolint@latest

# Set up Lefthook
echo "Setting up Lefthook..."
go install github.com/evilmartians/lefthook@latest
lefthook install

# Create commitlint config
echo "Creating commitlint configuration..."
cat > .commitlint.yml <<EOL
extends:
  - default
rules:
  type-enum:
    - 2
    - always
    - [feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert, dep]
  subject-case:
    - 2
    - always
    - [lower-case]
EOL

echo "Installation complete!"
