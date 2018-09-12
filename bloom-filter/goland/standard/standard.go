package standard

import (
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"goland/bloom"
	"hash"
	"hash/fnv"
	"math"
)

type StandardBloom struct {

	//hash函数
	h hash.Hash


	//位数组的大小，即容量
	// m =~ n / ((log(p)*log(1-p))/abs(log e)) （见维基百科）
	m uint


	//hash函数的个数
	// k = log2(1/e)
	// Given that our e is defaulted to 0.001, therefore k ~= 10, which means we need 10 hash values
	k uint

	// s is the size of the partition, or slice.
	// s = m / k
	s uint

	// p is the fill ratio of the filter partitions. It's mainly used to calculate m at the start.
	// p is not checked when new items are added. So if the fill ratio goes above p, the likelihood
	// of false positives (error rate) will increase.
	//
	// By default we use the fill ratio of p = 0.5
	//填充率，即位数组中设为1的占比
	p float64

	//错误率，
	// By default we use the error rate of e = 0.1% = 0.001. In some papers this is P (uppercase P)
	e float64

	// n is the number of elements the filter is predicted to hold while maintaining the error rate
	// or filter size (m). n is user supplied. But, in case you are interested, the formula is
	// n =~ m * ( (log(p) * log(1-p)) / abs(log e) )
	n uint

	// b is the set of bit array holding the bloom filters. There will be k b's.
	b *bitset.BitSet

	// c is the number of items we have added to the filter
	c uint

	// bs holds the list of bits to be set/check based on the hash values
	//添加一个元素时，首先计算该元素对应每个hash函数的hash值，
	bs []uint
}

//构造函数
func New(n uint) bloom.Bloom {
	var (
		p float64 = 0.5
		e float64 = 0.001
		k uint    = bloom.K(e)
		m uint    = bloom.M(n, p, e)
	)

	return &StandardBloom{
		h:  fnv.New64(),
		n:  n,
		p:  p,
		e:  e,
		k:  k,
		m:  m,
		b:  bitset.New(m),
		bs: make([]uint, k),
	}
}


//设置hash函数
func (this *StandardBloom) SetHasher(h hash.Hash) {
	this.h = h
}

func (this *StandardBloom) Reset() {
	this.k = bloom.K(this.e)
	this.m = bloom.M(this.n, this.p, this.e)
	this.b = bitset.New(this.m)
	this.bs = make([]uint, this.k)

	if this.h == nil {
		this.h = fnv.New64()
	} else {
		this.h.Reset()
	}
}

//设置错误率
func (this *StandardBloom) SetErrorProbability(e float64) {
	this.e = e
}

//计算填充率，
func (this *StandardBloom) FillRatio() float64 {
	return float64(this.b.Count()) / float64(this.m)
}

//求取k位的位置索引之后，分别置1
func (this *StandardBloom) Add(item []byte) bloom.Bloom {
	this.bits(item)
	for _, v := range this.bs[:this.k] {
		this.b.Set(v)
	}
	this.c++
	return this
}

func (this *StandardBloom) Count() uint {
	return this.c
}


func (this *StandardBloom) Check(item []byte) bool {
	this.bits(item)
	for _, v := range this.bs[:this.k] {
		if !this.b.Test(v) {
			return false
		}
	}

	return true
}

//预估填充率
func (this *StandardBloom) EstimatedFillRatio() float64 {
	return 1 - math.Exp((-float64(this.c)*float64(this.k))/float64(this.m))
}

func (this *StandardBloom) PrintStats() {
	fmt.Printf("m = %d, n = %d, k = %d, s = %d, p = %f, e = %f\n", this.m, this.n, this.k, this.s, this.p, this.e)
	fmt.Println("Total items:", this.c)
	c := this.b.Count()
	fmt.Printf("Total bits set: %d (%.1f%%)\n", c, float32(c)/float32(this.m)*100)
}

//将item放置到bloom中时，位数组需置1的位置索引， k个。
func (this *StandardBloom) bits(item []byte) {
	this.h.Reset()
	this.h.Write(item)
	s := this.h.Sum(nil)
	a := binary.BigEndian.Uint32(s[4:8])
	b := binary.BigEndian.Uint32(s[0:4])

	// Reference: Less Hashing, Same Performance: Building a Better Bloom Filter
	// URL: http://www.eecs.harvard.edu/~kirsch/pubs/bbbf/rsa.pdf
	for i, _ := range this.bs[:this.k] {
		this.bs[i] = (uint(a) + uint(b)*uint(i)) % this.m
	}
}