# Go types

This package provides alias/types for maps and slices.  
Thread-safe types also supported.

## Support atomic

### Maps :

|  Alias     |      Type                    |
|:----------:|:----------------------------:|
| Map        |  `map[string]any`    |

### Slices :

|  Alias     |      Type      |
|:----------:|:--------------:|
| Ints       |  `[]int`       |
| Uints      |  `[]uint`      |
| Int64s     |  `[]int64`     |
| Uint64s    |  `[]uint64`    |
| Floats     |  `[]float64`   |
| Strings    |  `[]string`    |
| Bytes      |  `[]byte`      |
| Bools      |  `[]bool`      |
| Slice      | `[]any`|

### Time & Date :
|  Alias     |      Wrapper   |      Type                    |
|:----------:|:---------------:|:------------:|
| Date       |  `NewDate()`    |  `Date`      |
| DateTime   | `NewDateTime()` |  `DateTime`  |


## Support sync

You can use the wrapper `types.SyncXXX()` :
```go
m := types.SyncMap()
```

|  Alias     |      Wrapper   |     Type                 |
|:----------:|:--------------:|:------------------------:|
| TSafeMap   | `SyncMap()`    | `map[string]any` |
| TSafeStrings   | `SyncStrings()`    | `[]string` |
| TSafeInts   | `SyncInts()`    | `[]int` |
| TSafeUints   | `SyncUints()`    | `[]uint` |
| TSafeInt64s   | `SyncInt64s()`    | `[]int64` |
| TSafeUint64s   | `SyncUint64s()`    | `[]uint64` |

