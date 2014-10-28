# Introduction

Naive performance test of doing type assertion using:

 * The static way

or...

 * Using reflection to avoid code repetition

# Results

```
BenchmarkAppendNativeTypeAssertion       5000000           666 ns/op
BenchmarkAppendReflectionTypeAssertion   5000000           529 ns/op
```

# Thanks to

 * Dan Kortschak <dan.kortschak@adel....edu.au>
 * Dmitry Vyukov <dvyukov@goo....com>
 * Xingtao Zhao  <zhaoxingtao@gm....com>
 * Egon          <egonelbre@gm....com>
 * Chris Dollin  <ehog.hedge@goo...mail.com>
