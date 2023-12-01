# My Go Project

This is a sample Go project that utilizes `godep` for managing dependencies.

## Table of Contents
- [Definition](#definition)
- [Usage](#usage)
    - [Installing Godep](#installing-godep)
    - [Initializing Godep for a Project](#initializing-godep-for-a-project)
    - [Restoring Dependencies](#restoring-dependencies)
    - [Running Your Go Application](#running-your-go-application)



## Definition

`godep` is a tool for managing Go application dependencies. It was designed to handle dependency management before the introduction of the official Go module system. It allows you to specify and lock the versions of your project's dependencies, providing a reproducible build environment.

## Usage

### Installing Godep

Make sure you have Go installed and your `GOPATH` set. Then, install `godep` using the following command:

```bash
go get -u github.com/tools/godep
```

### Initializing Godep for a Project
To initialize godep in your Go project, run:
```bash
godep init
```
This command will create a Godeps directory and a Godeps/Godeps.json file, where dependencies and their versions will be recorded.

### Adding Dependencies
To add a new dependency, use the godep save command:
```bash
godep save package/path@version
```
Replace package/path with the import path of the package you want to add and version with the desired version.

### Restoring Dependencies
To restore dependencies based on the Godeps/Godeps.json file, run:
```bash
godep restore
```
This will fetch and install the specified versions of your project's dependencies.

### Running Your Go Application
Here is a simple example demonstrating the use of godep:

1.Initialize godep:      
```bash
godep init
```
2.Add a dependency:
```bash
godep save github.com/example/package@v1.2.3
```
3.Restore dependencies:
```bash
godep restore
```

### Sample Godeps/Godeps.json File
An example of the generated Godeps/Godeps.json file:
```json
{
  "ImportPath": "github.com/yourusername/yourproject",
  "GoVersion": "go1.17",
  "GodepVersion": "v80",
  "Packages": [
    "./"
  ],
  "Deps": [
    {
      "ImportPath": "github.com/example/package",
      "Rev": "abc123",
      "Comment": "Version 1.2.3"
    }
  ]
}
```