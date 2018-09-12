package standard

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/spaolacci/murmur3"
	"github.com/zhenjl/cityhash"
	"hash"
	"hash/crc64"
	"hash/fnv"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"

	"testing"
)

func TestNew(t *testing.T) {

	var n uint  = 100000
	h := []hash.Hash{fnv.New64(), crc64.New(crc64.MakeTable(crc64.ECMA)), murmur3.New64(), cityhash.New64(), md5.New(), sha1.New()}
	l := []string{"fnv.New64()", "crc64.New()", "murmur3.New64()", "cityhash.New64()", "md5.New()", "sha1.New()"}


		for j := range h {
			fmt.Printf("\n\nTesting %d with size %s\n", n, l[j])
			bf := New(n)
			bf.SetHasher(h[j])
			//bf.PrintStats()
		}

}


//使用布隆过滤器对URL去重


func TestStandardBloom_Add(t *testing.T) {
	var n uint  = 100000
	bf := New(n)


	urlsList := getUrl("url.txt")

	fmt.Println(len(urlsList))

	for _, item := range urlsList {
		if !bf.Check([]byte(item)) {
			bf.Add([]byte(item))
		}

	}
	fmt.Println(bf.Count())

}


// 获取url
func getUrl(path string) (urlsList []string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	sreg := regexp.MustCompile(`(https.*)\n?`)
	sall := sreg.FindAllSubmatch(data, -1)
	for _, sitem := range sall {
		fmt.Println(string(sitem[0]))
		surl, err := url.Parse(string(sitem[0]))
		if err != nil {
			panic(err)
		}
		urlsList = append(urlsList, surl.String())
	}

	return urlsList
}
