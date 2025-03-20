# goweakpointer

# Output

```bash
$ goweakpointer-example 
main : value: 1 (count down: 5)
worker: want to work forever
worker: value: 2
worker: value: 3
main : value: 4 (count down: 4)
worker: value: 5
main : value: 6 (count down: 3)
worker: value: 7
main : value: 8 (count down: 2)
worker: value: 9
main : value: 10 (count down: 1)
worker: value: 11
main : finished
main : waiting worker
worker: value: 12
worker: lost the pointer!
worker: wont work forever anymore, exiting
main : waiting worker - done
```
