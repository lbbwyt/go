package bloom

import (
	"hash"
	"math"
)

type Bloom interface {
	Add(key []byte) Bloom
	Check(key []byte) bool
	Count() uint
	PrintStats()
	SetHasher(hash.Hash)
	Reset()
	FillRatio() float64
	EstimatedFillRatio() float64
	SetErrorProbability(e float64)
}

//返回hash函数的个数
func K(e float64) uint {
	return uint(math.Ceil(math.Log2(1 / e)))
}
//返回位数组的容量
func M(n uint, p, e float64) uint {
	// m =~ n / ((log(p)*log(1-p))/abs(log e))
	return uint(math.Ceil(float64(n) / ((math.Log(p) * math.Log(1-p)) / math.Abs(math.Log(e)))))
}

func S(m, k uint) uint {
	return uint(math.Ceil(float64(m) / float64(k)))
}