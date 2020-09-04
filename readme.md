# 🐦 Scarecrow

Scarecrow is a simple static site generator.

The entire idea behind Scarecrow is that it **stays out of your way**. Often static site generators are just far too complex, including plethoras of features and folders that you really don't need.

* Easy to customize.
* Allows you to write your site's content in Markdown. 
* Built-in blogging support.
* Supports Handlebars templating in Markdown.
* Compiles down to purely static files.

## Installation

Scarecrow is currently unpublished, but when it is it will be available via NPM and could be installed with this command:

```command
$ npm install -g @lukewhrit/scarecrow
```

## Usage

```
$ scarecrow --help

  🐦 Scarecrow is a simple static site generator

  Usage
    $ foo <input>

  Options
    --rainbow, -r  Include a rainbow

  Examples
    $ foo unicorns --rainbow
    🌈 unicorns 🌈
```

## Directory Structure

Scarecrow uses a very simple directory structure that allows for easy customization and development. The only required files are `layout.html` and `pages/index.md`.

```
.
 ├── assets/
 │    └── logo.jpg
 ├── posts/
 │    └── hello-world.md
 ├── pages/
 │    ├── contact.md
 │    ├── projects.md
 │    └── index.md
 └── layout.html
```

## Contributors

Scarecrow is built entirely from contributors to the Open-Source community.

* [Luke Whrit <lukewhrit@gmail.com>](https://github.com/lukewhrit) - Lead developer and maintainer.

## License

Scarecrow is publicly available under the terms of the [Apache license v2.0](license).
