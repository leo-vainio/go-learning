# go-learning

### Bypassing windows defender

Windows Defender sometimes flags go executables as a virus, making building and running the files impossible. To bypass this we can exclude a certain folder 
from being monitored by Windows Defender and building the .exe file into this folder by using the command shown below. This is however a minor inconvenience since
we need navigate to another folder in order to run the executable. It is however more safe to do it this way since we don't have to exclude for example $TEMP$
from windows defender. 

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
