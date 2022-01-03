![Logo](./Tunaflow.png)

![Apache license](https://img.shields.io/badge/license-Apache-blue?style=flat-square)
[![Build](https://github.com/borgmon/tunaflow/actions/workflows/go.yml/badge.svg)](https://github.com/borgmon/tunaflow/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/borgmon/tunaflow?status.svg)](https://pkg.go.dev/github.com/borgmon/tunaflow?tab=doc)

# Tunaflow
Tunaflow is a backend-as-config tool. It can generate a data transformation micro-service ready to deploy, all with YAML. With power of TunaSchool(WIP), easyjson and gin, your data flow can scale infinity and light-weight. 

## Overview
How many hours do you spend on simple data transformation tasks/data mapping services? Say hi to Tunaflow.

With Tunaflow you can write a yaml and generate a deploy-ready package instantly. 

All you need to do is define `schema`, and use `flow` to link everything together.

With TunaSchool you can balance flows across different packages.

## Installation
```bash
go get -u github.com/borgmon/tunaflow
```

## Usage
After create a new empty folder, run
```bash 
tunaflow init {package path}
```
This will generate an example config file for you to start with.

After writing your YAML config, run
```bash
tunaflow apply
```
to generate your project.

![Logo](./build.png)

## YAML Config Example
```yaml
version: 1 # YAML Config version
name: example # name of the project
app-version: v1.0 # version of the project
package-path: example # package path of the project, this is for go.mod

schemas:
  - name: input # name of the schema
    payload: # json schema with [key](data type)
      stringField: string
      boolField: boolean
      nested:
        value: int

  - name: output
    payload:
      name: string
      isFish: boolean
      value: int

flows:
  - name: myflow # name of the schema
    upstream: input # schema name for the incoming data
    downstream: output # schema name for the outgoing data
    mapping: # actual transformation. [downstream field name](upstream field name aka where to get it)
      name: stringField
      isFish: boolField
      value: nested.value

```
## Road Map
- [ ] json example file as input
- [ ] subflow to handle one to many mappings
- [ ] code injection in YAML to handle complex mappings (time format etc)
- [ ] openapi doc generation
- [ ] finish up go doc
- [ ] validation for config
- [ ] more CLI options(clean, build)
- [ ] exclusive files
- [ ] tests
- [ ] refactor using interfaces to allow extension
- [ ] pass file path in yaml
- [ ] TunaSchool
- [ ] allow custom middleware 
- [ ] async handler option
- [ ] fuzzy match or use ML
