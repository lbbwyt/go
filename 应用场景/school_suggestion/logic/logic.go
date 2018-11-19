// logic
package logic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var SchoolList []*SchoolInfo
var trieTree *Trie = NewTrie()

func init() {
	Init("./data/school.dat")
}

//统一错误处理
func handleError(err error, msg string) {
	if err != nil {
		fmt.Printf(" %s, err:%v\n", msg, err)
		panic(err)
	}
}

func Search(keyword string, limit int) (schools []*SchoolInfo) {

	nodes := trieTree.PrefixSearch(keyword, limit)
	//fmt.Printf("len:%d\n", len(nodes))
	for _, v := range nodes {
		school, ok := v.Data.(*SchoolInfo)
		if !ok {
			fmt.Printf("invalid school data:%v", v)
			continue
		}

		schools = append(schools, school)
	}
	return
}

func Init(filename string) (err error) {
	file, err := os.Open(filename)
	handleError(err, "打开文件失败")
	defer file.Close()

	var id int
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		handleError(err, "读取失败")

		strSplit := strings.Split(line, "\t")
		if len(strSplit) != 4 {
			fmt.Printf("无效的学校信息, line:%s\n", line)
			continue
		}

		var schoolInfo SchoolInfo
		id++
		schoolInfo.SchoolId = id
		schoolInfo.Province = strings.TrimSpace(strSplit[0])
		schoolInfo.City = strings.TrimSpace(strSplit[1])
		schoolInfo.SchoolName = strings.TrimSpace(strSplit[2])

		schoolType, err := strconv.Atoi(strings.TrimSpace(strSplit[3]))
		handleError(err, "读取学校信息失败")

		schoolInfo.SchoolType = schoolType
		SchoolList = append(SchoolList, &schoolInfo)
		//避免学校名字重复加上id
		trieTree.Add(fmt.Sprintf("%s%d", schoolInfo.SchoolName, schoolInfo.SchoolId), &schoolInfo)
	}
	return err
}
