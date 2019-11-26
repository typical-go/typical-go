# Typical Go

A Framework and Build Tool for Productive Go Development. <https://typical-go.github.io/>

## Install

Download the latest release from [releases page](https://github.com/typical-go/typical-go/releases)

For macOS:
```bash
curl -o typical-go -L https://github.com/typical-go/typical-go/releases/download/v0.9.2/typical-go_v0.9.2_darwin_amd64 && chmod +x typical-gos
```

## Usage

### New Project

```bash
typical-go new [PACKAGE]
```

### New Module

```bash
typical-go module [MODULE_NAME]
typical-go module [MODULE_NAME] -path=[PATH]
```

### Create Wrapper

```bash
typical-go wrapper
typical-go wrapper -path=[Path]
```


## Examples

- [RESTful Server](https://github.com/typical-go/typical-rest-server)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details




