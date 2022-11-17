# 🐦 Scarecrow

[![Go Report Card](https://goreportcard.com/badge/github.com/lukewhrit/scarecrow)](https://goreportcard.com/report/github.com/lukewhrit/scarecrow) [![Documentation](https://pkg.go.dev/badge/github.com/lukewhrit/scarecrow)](https://pkg.go.dev/github.com/lukewhrit/scarecrow)

Scarecrow is a simple static site generator.

The entire idea behind Scarecrow is that it **stays out of your way**. Often static site generators are just far too complex, including a plethora of features and folders that you really don't need.

**Abilities and Features:**

* Easy to customize.
* Allows you to write your site's content in Markdown. 
* Built-in blogging support.
* Supports Mustache templating in Markdown files.
* Compiles down to purely static files; JavaScript-free.
* Fast; Scarecrow compiles sites faster than Hugo or Jekyll.

## Table of Contents

- [🐦 Scarecrow](#-scarecrow)
	- [Table of Contents](#table-of-contents)
	- [Installation](#installation)
	- [Documentation](#documentation)
		- [Usage](#usage)
		- [Directory Structure](#directory-structure)
		- [Asset Handling](#asset-handling)
		- [Templates](#templates)
			- [Scarecrow-specific variables](#scarecrow-specific-variables)
				- [`posts`](#posts)
	- [TODO](#todo)
	- [License](#license)

## Installation

Scarecrow is pre-alpha software, it shouldn't be used in production or really even testing. However, you can still install Scarecrow right now and use it if you wish. Use this command: 

```
$ go get -v -u github.com/lukewhrit/scarecrow
```

## Documentation

### Usage

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

### Directory Structure

Scarecrow uss a simply directory structure, allowing for easy customization and development. The only required files are `pages/index.md` and `layout.html`.

```
.
 ├── assets/
 │    └── ...
 ├── posts/
 │    └── hello-world.md
 ├── pages/
 │    ├── contact.md
 │    └── index.md
 └── layout.html
```

`pages/index.md` is an ordinary markdown file, containing the home page of your website.

`layout.html`, on the other hand, is more unique. It's the base HTML code encasing the parsed Markdown in your files. Scarecrow injects the content of your site depending on which file it's in wherever it sees the `{{ body }}` tag (or `{{ Body }}`/`{{body}}`/etc).

For the time being, Scarecrow does not support custom per-file layouts. This will however change in future releases.

Although the `posts/` directory is 100% optional, it does hold special significance in that an array of its files and their front-matter are compiled into a `posts` object available to Mustache templates. For more information see the [Scarecrow-specific Variables](#scarecrow-specific-variables) section of the Templates documentation.

### Asset Handling

In the same fashion as many other static site generators, assets are held in a specific folder. This can cause issues with pages in sub-folders.

The afore mentioned specific folder is the `/assets/` directory, in this directory any and all images, styles, or scripts pertaining to your website should be placed.

The contents of the `/assets/` directory are not modified at all, with the only exceptions being CSS and JS. These assets are minified during the compile sequence.

To solve this, most static site generators implement a symbol which will reference their assets directory in paths and Scarecrow is no different.Scarecrow implements a `+` symbol that references the assets directory.

### Templates

Scarecrow implements support for Mustache templating, via [`cbroglie/mustache`](https://github.com/cbroglie/mustache). This library implements almost the entire Mustache specification, but includes only experimental suppot for lambdas.

Please note that the `layout.html` file does not support standard Mustache templates, and only allows for the sole `{{ body }}` tag.

Documentation and reference material on Mustache are available on the [mustache.github.io website](https://mustache.github.io/mustache.5.html).

#### Scarecrow-specific variables

Scarecrow implements numerous custom variables and functions that are made available within Mustache templates.

##### `posts`

`posts` is an array containing an object:

* `title: string`: title of the post
* `href: string`: link to the post

Each of these objects represents a post in the `posts/` directory.

## TODO

* [ ] Blogging support.
* [X] Use `layout.html` to determine file content and output to a `dist/` directory.
* [X] Passing to a Mustache compiler before outputting.
* [ ] Front matter support and parsing.
* [X] HTML output minifying.
* [ ] Built-in web server with automatic source rebuilding.
* [ ] Custom, per-file layouts.
* [X] Automatically create required directories on output.
* [ ] Asset Handling
  * [ ] Automatic minifying of CSS and JS assets.
  * [ ] Support for referencing the asset directory in HTML files via `+` symbol.
  * [X] Basic asset handling (moving assets to output directory, etc.).

## License

Scarecrow is available to the public under the terms of the [Apache license v2.0](license).
