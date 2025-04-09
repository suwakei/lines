<img src="https://raw.githubusercontent.com/suwakei/logo/main/lines/lines_logo.png" alt="image-logo" width="400px" height="150px">

# lines

[![Build Status](https://github.com/suwakei/lines/actions/workflows/build.yml/badge.svg)](https://github.com/suwakei/lines/actions/workflows/build.yml)
[![Test Status](https://github.com/suwakei/lines/actions/workflows/test.yml/badge.svg)](https://github.com/suwakei/lines/actions/workflows/test.yml)
[![Lint Status](https://github.com/suwakei/lines/actions/workflows/lint.yml/badge.svg)](https://github.com/suwakei/lines/actions/workflows/lint.yml)

## Overview
lines is a CLI application that counts the number of lines, blanks, comments, files, bytes. in the file of the input path and outputs them.

## Table of Contents
- [Overview](#overview)
- [Example](#example)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
    - [Supported Extentions](#supported-extentions)
    - [Supported Options](#supported-options)

## Example
<img src="https://raw.githubusercontent.com/suwakei/logo/main/lines/lines_example.gif" alt="example-image" width="1000px" height="350px">

## Features

- lines is **accurate**, lines handles multi line comments correctly,
and not counting comments that are in strings.


## Installation

## Usage

### Supported Extentions
> [!IMPORTANT]
> Files not included in the following extensions will not be colored,
> but will be displayed in the count results.

[Ext] [Detail] [Color]
```console
	.asm          Assembly(.asm) Red
	.bat          Batch File(.bat) Cyan
	.bash         BASH(.bash) HiGreen
	.c            C(.c) Yellow
	.cc           C++(.cc) Yellow
	.cs           C#(.h) Cyan
	.css          CSS(.css) Yellow
	.cfg          Configuration File(.cfg) HiBlack
	.cpp          C++(.cpp) Yellow
	.clj          Clojure(.clj) HiGreen
	.coffee       CoffeeScript(.coffee) Yellow
	.d            D(.d) Red
	.dart         Dart(.dart) Cyan
	.dockerfile   Dockerfile(.dockerfile) HiBlue
	.Dockerfile   Dockerfile(.Dockerfile) HiBlue
	Dockerfile    Dockerfile HiBlue
	.dockerignore Docker ignore file(.dockerignore) HiBlue
	.fs           F#(.fs) Cyan
	.f90          Fortran(.f90) HiWhite
	.erl          Erlang(.erl) HiWhite
	.ex           Elixir(.ex) Magenta
	.exs          Elixir Script(.exs) Magenta
	.go           Go(.go) Blue
	.groovy       Groovy(.groovy) Yellow
	.gitignore    Git ignore file(.gitignore) HiWhite
	.h            C(.h) Yellow
	.hpp          C++(.hpp) Yellow
	.html         HTML(.html) HiYellow
	.ini          Initialization File(.ini) HiWhite
	.js           JavaScript(.js) Yellow
	.jsx          JSX(.jsx) Yellow
	.java         Java(.java) Red
	.json         JSON File(.json) Yellow
	.jsonc        JSONC File(.jsonc) Yellow
	.kt           Kotlin(.kt) Green
	.lua          Lua(.lua) Cyan
	.log          Log File(.log) HiWhite
	.less         Less(.less) Cyan
	LICENSE       LICENSE File Yellow
	.m            Objective-C(.m) Yellow
	.ml           ML(.ml) HiYellow
	.md           Markdown(.md) Cyan
	.mk           Makefile(.mk) HiRed
	.mod          Go modules file(.mod) Blue
	Makefile      Makefile HiRed
	.nim          Nim(.nim) YellowGreen
	.py           Python(.py) HiCyan
	.pl           Perl(.pl) Cyan
	.pas          Pascal(.pas) HiWhite
	.php          PHP(.php) Blue
	.r            R(.r) Cyan
	.rb           Ruby(.rb) Red
	.rs           Rust(.rs) HiBlack
	.rtf          Rich Text Format(.rtf) HiWhite
	.raku         Raku(.raku) HiWhite
	.s            Assembly(.s) Red
	.svg          SVG Magenta
	.sh           Shell Script(.sh) Green
	.sql          SQL(.sql) Pink
	.sum          Go sum file(.sum) Blue
	.sass         SASS(.sass) Red
	.scss         SCSS(.scss) Red
	.swift        Swift(.swift) HiYellow
	.scala        Scala(.scala) Red
	.ts           TypeScript(.ts) Blue
	.tcl          Tcl(.tcl) HiWhite
	.tsx          TSX(.tsx) Blue
	.txt          Plain Text(.txt) HiWhite
	.toml         TOML(.toml) HiBlack
	.v            Verilog(.v) HiWhite
	.vue          Vue file(.vue) Green
	.vhdl         VHDL(.vhdl) HiWhite
	.wasm         WebAssembly Magenta
	.xml          XML(.xml) Blue
	.xsl          XSLT(.xsl) HiWhite
	.yml          YAML File(.yml) Magenta
	.yaml         YAML File(.yaml) Magenta
	.zsh           ZSH Green
	.zig          Zig(.zig) HiYellow
```

### Supported Options
```bash
Usage:
  lines <PATH> [OPTIONS]

Flags:
  -d, --dist strings    input filepath to output. output format [.json, .jsonc, .yml, .yaml, .toml, .txt]
  -e, --ext strings     input extension you don't want to count "-e=test.json, .js, .go" or "-e=test.json -e=.js -e=.go". (default: .exe, .com, .dll, .so, .dylib, .xls, .xlsx, .pdf, .doc, .docx, .ppt, .pptx)
  -h, --help            help for lines
  -i, --ignore string   input your .gitignore file path. ignore extentions in .gitignore file. (default: .gitignore)
  -o, --only string     By specifying an extension or file name, only files with that extension or name are targeted. "-o=.go" or "-o .go" or "-o=test.txt"
  -v, --version         Print version of this app
```