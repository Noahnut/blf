[![Build Status](https://travis-ci.com/Noahnut/blf.svg?branch=main)](https://travis-ci.com/Noahnut/blf)

# Bloom Filter
Simple Bloom Filter implement by golang
Bloom Filter is the simple and space-efficient probabilistic data structure to check the data is exist or not. but Bloom Filter have the false positive rate. false positive rate mean the data is not in the set or storage but Bloom Filter return true.
This implement is the **Count Bloom Filter** provide
1. Add element to Bloom Filter
2. Query element from Bloom Filter
3. Remove element from Bloom Filter

## Detail
Use the murmur3 hash as the Bloom Filter hash function

## Install 
```shell
go get github.com/Noahnut/blf
```

## Usage
```go
expectFalsePositiveRate := 0.5
// contruct the bloomFilter with expect element number and False Positive Rate
blf := ContructbloomFilter(100, expectFalsePositiveRate)

// Add the AAAA to the bloom filter
blf.Add([]byte("AAAA"))

// return true  "AAAA" exist in the bloom filter
blf.Query([]byte("AAAA")) 

// return false "BBB" not exist in the bloom filter
blf.Query([]byte("BBB"))

// remove "AAAA" from the bloom filter
blf.Delete([]byte("AAAA"))

// return false "AAAA" not exist in the bloom filter
blf.Query([]byte("AAAA"))
```