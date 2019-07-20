# go-geopoint

GPS Point encoding library with privacy in mind.

```go
import "go.zenithar.org/geopoint"

func foo() {
  p := geopoint.Encode(43.603574, 1.442917)

  fmt.Printf("%d\n", uint64(p))
  // 75071809151126838

  fmt.Printf("%s\n", p.Code())
  // 10AB5:69A51:94D36
}
```
## What's the difference with geohash ?

For compatibility, you should use geohash. 

The main objectives is to find a way to encode point (lat,lon) as an uint64 
(like geohash) and make it sortable.

## Brain dump

I tryied to apply CryptoPan anonimyzer on the uint64 encoded coordinate, I know that
this algorithm is told to be prefix-preserving, but after application of the algorithm
point are not keeping their geospace properties at first sight.

I'm trying to find a cryptographic keyed hash function for GPS point, and the result
should be a GPS point too. Maybe all GPS related arithmetic should be influenced by 
the cryptographic hash, like in homothecy. GPS Point anonymization are based on 
precision loss.

A real GPS point is projected in another `planet` GPS point. The other planet is 
generated with the `key` of the hash function. An origin point as always the same 
other planet point value (injective function but collision must be proved also). 
It could be visualized with a homothetic transformation of the original planet 
were the `key` could be part of the homogeneous dilatation factor, but not only
countries of the generated earth should have different shapes.

Maybe I should investigate in Format Preserving Encryption field...
I'm not a mathematician or a scientist, just an intuitive person.

Maybe this is a full bowl of shit, waiting for being flushed !

## Fun ideas

:star2: It could be used to represents GPS coords on other planets !

```
$ <planet>:<tile>:<inner-x>:<inner-y>
```

