# go-learning

### Bypassing windows defender

Windows defender sometimes flags go executables as potential virus making running the files impossible. To bypass this we can exclude a certain folder from being
checked by windows defender and performing the build into this folder. Running it from there will therefore not be an issue. This is however an inconvenience since
we need to be in another folder in order to run the file. But it's more safe this way since we don't have to exclude for example $TEMP$ from windows defender. 

```console

> go build -o D:\go-builds\out.exe

```


TODO:

- Generics
- Packages
- Testing
- init function
- Look through Effective Go
- 11 go projects
- Blog posts
- Getting familiar with packages on the go website
- writing web applications
- input/output
- Solve one kattis problem maybe?
- data structures and algorithms in go
- the go mod tool and other command line commands that can be useful
- how to organise a go project.
- gophercises
- modules
- https://go.dev/doc/code
- https://go.dev/doc/
- Documentation
