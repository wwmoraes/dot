# Keys

All known [Graphviz attribute names](graphviz-attrs) are defined as constants,
so you can opt-in and use those to prevent typos on your code. They're all plain
`string`s named after the original attribute and prefixed with `Key`, so the
constant for the attribute `label` is `KeyLabel`.

```go
--8<-- "attributes/keysConstants.go"
```

[graphviz-attrs]: https://graphviz.org/doc/info/attrs.html
