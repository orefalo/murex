# _murex_ Shell Guide

## Data-Type Reference: `int` (integer)

> Whole number (primitive)

### Description

An integer is a whole number (eg 1, 2, 3, 4) rather than one with a decimal
point (such as 1.1).

Integers in _murex_ are sized based on the bit (or word) size of the target
CPU.

A 386, ARMv6 or other 32bit build of _murex_ would see the range of from
`-2147483648` (negative) through `2147483647` (positive).

AMD64 or other 64bit built of _murex_ would see the range from
`-9223372036854775808` (negative) through `9223372036854775807` (positive).

> Unless you specifically know you only want whole numbers, it is recommended
> that you use the default numeric data-type: `num`.



### Default Associations





### Supported Hooks

* `Marshal()`
    Supported
* `Unmashal()`
    Supported

### See Also

* [`Marshal()` ](../apis/marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` ](../apis/unmarshal.md):
  Converts a structured file format into structured memory
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [open](../commands/open.md):
  
* [str](../types/str.md):
  