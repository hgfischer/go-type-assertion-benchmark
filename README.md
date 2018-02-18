# Introduction

Native performance test of doing type assertion using:

 * The static way

or...

 * Using reflection to avoid code repetition

# Results

```
goos: darwin
goarch: amd64
BenchmarkAppendNativeTypeAssertion-8       10000000       129 ns/op      56 B/op       3 allocs/op
BenchmarkAppendReflectionTypeAssertion-8   10000000       144 ns/op      56 B/op       3 allocs/op
```

# Thanks to

 * Dan Kortschak <dan.kortschak@adel....edu.au>
 * Dmitry Vyukov <dvyukov@goo....com>
 * Xingtao Zhao  <zhaoxingtao@gm....com>
 * Egon          <egonelbre@gm....com>
 * Chris Dollin  <ehog.hedge@goo...mail.com>
