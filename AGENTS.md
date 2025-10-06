ITO
===

`ito` is a tiny file manager.

`ito` probides just file paths in a specific directory.

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


