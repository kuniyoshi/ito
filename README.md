ITO
===

`ito` is an extremely lightweight file manager.

`ito` provides quick access to files within a specific directory.

USAGE
=====

```
$ ls -1 ~/.ito
foo
bar
baz
$ ito list
foo
bar
baz
$ ls $(ito foo)
~/.ito/foo
```

PHILOSOPHY
==========

`ito` behaves almost like `alias ito='ls ~/.ito'`.

Its deliberately tiny feature set helps you stay focused on the task at hand
and remember what each file is for.

