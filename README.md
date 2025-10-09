ITO
===

`ito` is a very tiny file manager.

`ito` probides file paths in a specific directory.

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

`ito` is almost like `alias ito ls ~/.ito`.

Small degree of freedom makes users on focus it's task,
and remember what files are.


