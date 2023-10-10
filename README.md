# Short URL Generator
Go implementation for generating Tiny URL- and bit.ly-like URLs, based on
Python's [short_url](https://pypi.org/project/short_url).

The intended use is to obfuscate integer IDs into short strings which can be
used in your URLs.  A bit-shuffling approach is used to avoid generating
consecutive, predictable strings.  Furthermore, the algorithm is deterministic
and guarantees that no collisions will occur.

The URL alphabet is fully customizable and may contain any number of
characters.  By default, digits and lower-case letters are used, with
some removed to avoid confusion between characters like `o`, `O` and `0`.  The
default alphabet is shuffled and has a prime number of characters to further
improve the results of the algorithm.

The block size specifies how many bits will be shuffled.  The lower `blockSize`
bits are reversed.  Any bits higher than `blockSize` will remain as is.
`blockSize` of `0` will leave all bits unaffected and the algorithm will simply
be converting your integer to a different base.  The `minLength` parameter
allows you to pad the string if you want it to be a specific length.

Sample Usage:

```
import "shortcode"
code := shortcode.EncodeID(12) // "LhKA"
id := shortcode.DecodeID(code) // 12
```

Use the top-level functions of the module to use the default encoder settings.
Otherwise, you may create your own custom `Encoder` and use its `EncodeID` and
`DecodeID` methods.
