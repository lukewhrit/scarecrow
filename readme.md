# 🐦 Scarecrow

Scarecrow is a simple static site generator.

The entire idea behind Scarecrow is that it **stays out of your way**. Often static site generators are just far too complex, including plethoras of features and folders that you really don't need.

**Abilities/Features:**

* Easy to customize.
* Allows you to write your site's content in Markdown. 
* Built-in blogging support.
* Supports Go's `html/template` in Markdown.
* Compiles down to purely static files; JavaScript-free.

## Installation

Scarecrow is currently unpublished, but when it is it will be available via the Go tool-chain and could be installed with this command:

```
$ go get -v -u github.com/lukewhrit/scarecrow
```

## Usage

```
$ scarecrow --help
Scarecrow is a modern and simple static site generator

Usage:
  scarecrow [command]

Available Commands:
  help        Help about any command
  make        Compile a Scarecrow project

Flags:
  -h, --help      help for scarecrow
  -v, --version   version for scarecrow

Use "scarecrow [command] --help" for more information about a command.
```

## Directory Structure

Scarecrow uses a very simple directory structure that allows for easy customization and development. The only required files are `layout.html` and `pages/index.md`.

```
.
 ├── posts/
 │    └── hello-world.md
 ├── pages/
 │    ├── contact.md
 │    ├── projects.md
 │    └── index.md
 └── layout.html
```

`pages/index.md` is pretty plain, `layout.html` however has some special things going on with it. `layout.html` serves as the base HTML file for all content on your site. You should define the styles and layout of your site here.

Scarecrow will inject your sites content into this file where it finds a `<scarecrow-body />` tag.

Scarecrow does not yet support custom layouts per file. We are interested in possibly implementing this in the future.

## TODO

* [ ] Blogging support
* [ ] Use `layout.html` to determine file content and output to a `dist/` directory
* [ ] Pass to `html/template` before outputting
* [ ] Front matter support and parsing

## Contributors

Scarecrow is built entirely from contributors to the Open-Source community. Here are some of the notable contributors to Scarecrow:

* [Luke Whrit <lukewhrit@gmail.com>](https://github.com/lukewhrit) - Lead developer and maintainer.

Scarecrow is **not** supported by and has no interest in pandering to any corporation.

## License

Scarecrow is publicly available under the terms of the [Apache license v2.0](license).
