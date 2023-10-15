## Presentation regarding "Mutation Testing In Go"

This repository contains a presentation regarding "Mutation Testing In Go" and
the source code used in the examples. 

## Run Tests

Standard tests can be run just like you are used to:

```shell
go test ./...
```

Mutation tests can be run via a specific tag, to avoid always running them:

```shell
go test -tags mutation ./...
```

## View Presentation

The presentation is located at [SLIDES.md](/SLIDES.md) and can be viewed via
any text editor, as it is markdown. However, it is a reveal.js slide and can
be viewed as such. I recommend simply using the VSCode extension
`evilz.vscode-reveal`.

## Windows

To run this example on windows, you need to get a special fork of ooze first.

```shell
git clone --branch temp_win_support git@github.com:Bios-Marcel/ooze.git 
```

The [go.mod](/go.mod) reference `../ooze`, so make sure to clone / move into
the right spot.
