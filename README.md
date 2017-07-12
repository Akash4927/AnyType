# AnyType

Templates, and ready-to-use go lang source files - generated with dotgo.

## Subjects
This repo provides it's subjects (projects) in it's immediate subdirectories:

-  [chan](chan/ReadMe.md) - Gain concurrency - use go-channels for piping (and more)

## Directories
Each directory has many subdirectories.

Most contain 'target directories' for the generation and thus become ready-to-use go packages, accessed e.g. via

	  go get github.com/GoLangsam/AnyType/chan/sss/test

or

	  go get github.com/GoLangsam/AnyType/chan/sss/standard/io

## Documentation
You'll find many `*.md` files - such as this one.

They may help to understand how these simple, selfcontained and ready-to-use go packages came into existence.

Hint: The sources in the packages are documented - You may use `go doc`.

The `*.md` documenation is mostly *meta* and may give better background for underlying desing principles and concepts.

Happy reading :-)

Mind You: *Simplicity is complicated* :-)

## Templates
You'll find templates such as

- `dot.go.tmpl` - in target directories
- `*.dot.go.tmpl` - in subject directories
- other `*.tmpl` - all the way down into sub-subject directories

##### `dot.go.tmpl` - in target directories
- Required to exist in any target directory
  (otherwise, dotgo would not write output; a security measure)
- Defines context for a target directory - we keep data close to results.
  
##### `*.dot.go.tmpl` - in subject directories
- Defines files (and their basename) to be generated into some target directory.

##### other `*.tmpl - all the way down into sub-subject directories
- Define further portions / subtemplates. 

Hint: Especially `*{{.}}*.tmpl` are usually just simple go code sniplets.

**They** are the meat here, where the **architecture** are the bones.

## Source files

Well, it's about [go](http://golang.org). So we have 

### `*.go` - ready-to-use compilable packages

It started as a proof-of-concept, and grows into a ready-to-use collection which directly extends the standard package, and even beyond.

### `*.ugo` - **ugly** go souces

Results generated from the templates - still **ugly**, before the pass thru `gofmt`.

You may ignore them, or `diff` them agains their `*.go` twin in order to see the templates and sniplets to be well formated.
