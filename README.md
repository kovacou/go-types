# Go types

This package provides alias/types for maps and slices.  
Thread-safe types also supported.

## Support atomic

* Maps :
    * `Map`

* Slices :
    * `Uints`
    * `Int64s`
    * `Uint64s`

## Support sync

You need to call `Init()` before using the type.
```go
m := SyncMap{}
m.Init(5) // Initialize mutex and internal value.
```

* Maps
    * `SyncMap`