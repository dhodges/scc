#### Seek coding test

An exercise in modelling simple pricing of customer ads.

#### Install

This project uses the go progamming language. 
Once you have go installed, run:

```
$ go get github.com/dhodges/scc
```

Then:

```
$ cd $GOPATH/src/github.com/dhodges/scc
$ go run main.go
```

You should then see output like the following:

![screen showing example session](https://github.com/dhodges/scc/raw/master/scc_screenshot.png)

#### Testing

```
$ go test ./checkout
$ go test ./pricing
```

#### Assumptions

The spec suggests two types of deals, e.g. a 3-for-2 deal on "Classic" Ads or a discount on "Premium" Ads.

This code assumes that the two deals are mutually exclusive, e.g. if customer "SecondBite" has a 3-for-2 deal on Classic Ads then they will *not* also have a discount on Classic Ads.
