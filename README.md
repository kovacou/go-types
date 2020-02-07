# Go types

This package provides alias/types for maps and slices.  
Thread-safe types also supported.

## Support atomic

### Maps :

|  Alias     |      Type                    |
|:----------:|:----------------------------:|
| Map        |  `map[string]interface{}`    |




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

## Support sync

You can use the wrapper `types.SyncXXX()` :
```go
m := types.SyncMap()
```

Or use the struct definition `types.TSafeXXX` :
```go
m := types.TSafeMap{}
m.Init(5)
```

|  Alias     |      Wrapper   |     Type                 |
|:----------:|:--------------:|:------------------------:|
| TSafeMap   | `SyncMap()`    | `map[string]interface{}` |
