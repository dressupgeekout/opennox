# Building for Linux

This is the main build procedure for Linux version.

## Installing Go compiler

### Fastest way

```bash
sudo snap install --classic go
```

### Longer way

See [official docs](https://golang.org/doc/install).

## Installing C compiler

```bash
sudo apt install gcc
```

## Installing other dependencies

```bash
sudo dpkg --add-architecture i386
sudo apt update
sudo apt install libsdl2-dev libsdl2-dev:i386 libopenal-dev libopenal-dev:i386
```

## Building

```bash
cd src
go run ./internal/noxbuild
```

This should produce `opennox`, `opennox-hd` and `opennox-server` binaries.
