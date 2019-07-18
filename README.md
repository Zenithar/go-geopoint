# go-geopoint

GPS Point encoding library with privacy in mind.

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
## Privacy encoding

:alert: **Experimental** DON'T USE IT !!!
No proof to work, ATM, just an idea implemented.

```go
import "go.zenithar.org/geopoint"

var (
  key = []byte(.....)
)

func foo() {
  p := geopoint.Encode(43.603574, 1.442917)

  fmt.Printf("%d\n", uint64(p))
  // 31659983379082793

  fmt.Printf("%s\n", p.Code())
  // 0707A:964EE:7AE29

  // Anonymize the point
  cpan, _ := geopoint.NewCryptoPan(key)
  ano := cpan.Anonymize(p)

  fmt.Printf("%d\n", uint64(ano))
  // 40028174716349995

  fmt.Printf("%s\n", ano.Code())
  // 08E35:69AEF:9AE2B
}
```

## Fun ideas

:star2: It could be used to represents GPS coords on other planets !

```
$ <planet>:<tile>:<inner-x>:<inner-y>
```

