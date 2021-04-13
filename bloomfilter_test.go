package blf

import (
	"strconv"
	"testing"
)

func TestBasicblf(t *testing.T) {
	blf := ContructbloomFilter(100, 0.1)
	blf.Add([]byte("AAAA"))

	if blf.Query([]byte("AAAA")) != true {
		t.Error("AAAA should in filter")
	}

	if blf.Query([]byte("BBB")) != false {
		t.Error("BBB should not in filter")
	}

	blf.Delete([]byte("AAAA"))

	if blf.Query([]byte("AAAA")) != false {
		t.Error("AAAA should not in filter")
	}
}

func TestSimpleBenchmark(t *testing.T) {

	blf := ContructbloomFilter(100, 0.5)
	var inputList = []int{120, 150, 121, 422, 456, 121, 987, 111, 121}

	for _, e := range inputList {
		blf.Add([]byte(strconv.Itoa(e)))
	}

	for _, e := range inputList {
		if blf.Query([]byte(strconv.Itoa(e))) == false {
			t.Errorf("%d should not in filter", e)
		}
	}
}

func TestBenchmark100data(t *testing.T) {

	blf := ContructbloomFilter(100, 0.5)
	inputList := make([]int, 100)

	for i := 0; i < len(inputList); i++ {
		inputList[i] = i * 10
	}

	for _, e := range inputList {
		blf.Add([]byte(strconv.Itoa(e)))
	}

	for _, e := range inputList {
		if blf.Query([]byte(strconv.Itoa(e))) == false {
			t.Errorf("%d should not in filter", e)
		}
	}
}

func TestFalsePositiveRate(t *testing.T) {
	expectFalsePositiveRate := 0.5
	blf := ContructbloomFilter(100, expectFalsePositiveRate)
	inputList := make([]int, 100)

	for i := 0; i < len(inputList); i++ {
		inputList[i] = i
	}

	for _, e := range inputList {
		blf.Add([]byte(strconv.Itoa(e)))
	}
	count := 0
	for i := 101; i <= 200; i++ {
		if blf.Query([]byte(strconv.Itoa(i))) == true {
			count++
		}
	}

	falsePositiveRate := float64(count) / float64(100)
	if falsePositiveRate < expectFalsePositiveRate-0.1 && falsePositiveRate >= expectFalsePositiveRate+0.1 {
		t.Error("falsePositiveRate is not correct")
	}
}
