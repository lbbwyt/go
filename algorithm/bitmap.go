//BitMap算法
package algorithm

import "fmt"

const (
	BitSize = 8 //一个字节8位
)

type Bitmap struct {
	BitArray  []byte
	ArraySize uint32
}

//max 为最大数，
//bitmap排序的时间复杂度不是O(N)的，而是取决于待排序数组中的最大值max，
func NewBitmap(max uint32) *Bitmap {
	var r uint32
	switch {
	case max <= BitSize:
		r = 1
	default:
		r = max / BitSize
		if max%BitSize != 0 {
			r++
		}

	}
	fmt.Println("数组大小:", r)
	return &Bitmap{
		BitArray:  make([]byte, r),
		ArraySize: r,
	}
}

func (bitmap *Bitmap) Set(i uint32) {
	idx, pos := bitmap.calc(i)
	//	|=是位操作运算符的一种,
	//	代表的含义为a=a|b;即把a和b做按位或(|)操作，结果赋值给a。
	bitmap.BitArray[idx] |= 1 << (pos - 1)
	fmt.Println(bitmap.BitArray[idx])
	fmt.Println("set value=", i, " idx=", idx, " pos=", pos, ByteToBinaryString(bitmap.BitArray[idx]))
}

func (bitmap *Bitmap) Clear(i uint32) {
	idx, pos := bitmap.calc(i)
	//&= 是按位与之后赋值，^=是按位异或之后赋值，|=是按位或之后赋值。
	bitmap.BitArray[idx] &^= 1 << (pos - 1)
	fmt.Println("clear value=", i, " idx=", idx, " pos=", pos, ByteToBinaryString(bitmap.BitArray[idx]))
}

//计算一个无符号32位数字的位置
func (bitmap *Bitmap) calc(i uint32) (idx, pos uint32) {

	idx = i >> 3 //相当于i / 8,即字节位置
	if idx >= bitmap.ArraySize {
		panic("数组越界.")
		return
	}
	pos = i % BitSize //位位置
	return
}

//转出byte变量的二进制字符串
func ByteToBinaryString(data byte) (str string) {
	var a byte
	for i := 0; i < 8; i++ {
		a = data
		data <<= 1
		data >>= 1

		switch a {
		case data:
			str += "0"
		default:
			str += "1"
		}

		data <<= 1
	}
	return str
}
