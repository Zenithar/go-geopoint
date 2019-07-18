# go-geopoint

GPS Point encoding library

```go
import "go.zenithar.org/geopoint"

func foo() {
  p := geopoint.Encode(43.603574, 1.442917)

  fmt.Printf("%d\n", uint64(p))
  // 75071988303315493

  fmt.Printf("%s\n", p.Code())
  // 10AB5:935B6:6C225
}
```

## Fun ideas

:star2: It could be used to represents GPS coords on other planets !

```
$ <planet>:<tile>:<inner-x>:<inner-y>
```

