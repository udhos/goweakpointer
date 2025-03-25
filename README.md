# goweakpointer

# Output

```bash
$ goweakpointer-example 
main   : value:  1 (count down:  5)
worker : want to work forever
worker : value:  2
worker : value:  3
main   : value:  4 (count down:  4)
main   : value:  5 (count down:  3)
worker : value:  6
worker : value:  7
main   : value:  8 (count down:  2)
worker : value:  9
main   : value: 10 (count down:  1)
worker : value: 11
main   : finished
main   : waiting worker
worker : value: 12
worker : lost the pointer!
worker : won't work forever anymore, exiting to release resources
main   : waiting worker - done
```
