## flextime-linter

flextime-linter restrict to use time package, and suggest code using [Songmu/flextime](https://github.com/Songmu/flextime).

## Install

```bash
# MacOS 
curl -L https://github.com/wreulicke/flextime-linter/releases/download/v0.0.1/flextime-linter_0.0.1_darwin_amd64 -o /usr/local/bin/flextime-linter

# Linux
curl -L https://github.com/wreulicke/flextime-linter/releases/download/v0.0.1/flextime-linter_0.0.1_linux_amd64 -o /usr/local/bin/flextime-linter

# Windows
curl -L https://github.com/wreulicke/flextime-linter/releases/download/v0.0.1/flextime-linter_0.0.1_windows_amd64.exe -o <path-directory>/flextime-linter.exe
```

## Usage

flextime-litner replace to `flextime` from `time` automatically.
You can try below.

```
flextime-linter -fix ./...
```