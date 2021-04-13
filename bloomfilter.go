package blf

import (
	"fmt"
	"math"

	"github.com/twmb/murmur3"
)

type blf struct {
	bitArraySize uint64
	bitArray     []int
	hashNumber   uint64
}

func murmurHash(data []byte, seed uint64) uint64 {
	h64 := murmur3.SeedNew64(seed)
	h64.Write(data)
	return h64.Sum64()
}

func ContructbloomFilter(itemSize uint64, falsePositiveRate float64) blf {
	if falsePositiveRate > 1.0 {
		fmt.Print("falsePositiveRate can not greater than One")
		return blf{}
	}
	blfm := blf{}
	blfm.countTheArraySizeAndHashNumber(itemSize, falsePositiveRate)
	blfm.bitArray = make([]int, blfm.bitArraySize)
	return blfm
}

func (this *blf) countTheArraySizeAndHashNumber(itemSize uint64, falsePositiveRate float64) {
	// m = ceil((n * log(p)) / log(1 / pow(2, log(2)))) number of bit array
	// k = round((m / n) * log(2)); number of hash function
	fNumber := float64(itemSize)
	this.bitArraySize = uint64(math.Ceil((fNumber * math.Log(falsePositiveRate)) / math.Log(1/math.Pow(2, math.Log(2)))))
	this.hashNumber = uint64(math.Round(float64(this.bitArraySize/itemSize) * math.Log(2)))
}

func (this *blf) Query(value []byte) bool {
	for i := uint64(0); i < this.hashNumber; i++ {
		value := murmurHash(value, i)
		if this.bitArray[value%this.bitArraySize] == 0 {
			return false
		}
	}
	return true
}

func (this *blf) Add(value []byte) error {
	for i := uint64(0); i < this.hashNumber; i++ {
		value := murmurHash(value, i)
		this.bitArray[value%this.bitArraySize]++
	}
	return nil
}

func (this *blf) Delete(value []byte) {
	for i := uint64(0); i < this.hashNumber; i++ {
		value := murmurHash(value, i)
		if this.bitArray[value%this.bitArraySize] != 0 {
			this.bitArray[value%this.bitArraySize]--
		}
	}
}
