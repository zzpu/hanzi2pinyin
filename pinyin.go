package main

import (
	//
	"hanzi2pinyin/data"


	"fmt"
	//"os"
	//"io"
	//"bufio"
	//"strings"
	//"strconv"
	//"encoding/json"
	"encoding/json"
	"strconv"
	"strings"
)

var(
	dict map[int] string
)
func init(){


	dict = make(map[int]string)
	//f, err := os.Open("dict/Mandarin.dat")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//rd := bufio.NewReader(f)
	//for {
	//	line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
	//
	//	if err != nil || io.EOF == err {
	//		break
	//	}
	//	pair := strings.Split(line,"\t")
	//	k,_ := strconv.Atoi(pair[0])
	//	dict[k]=pair[1]
	//	fmt.Print(pair[0]," ",pair[1])
	//}
	//
	//b, err := json.Marshal(dict)
	//
	//f1, err := os.Create("dict/dict.dat")
	//if err != nil {
	//	panic(err)
	//}
	//f1.Write(b)

	bytes, err := data.Asset("dict/dict.dat")  // 根据地址获取对应内容
	if err != nil {
		panic(err)
	}

	//wordPairs := strings.Split(string(bytes),"\n")

	json.Unmarshal(bytes,&dict)

	for k,d := range dict{
		s:=""
		for _,c := range []byte(d){
			if string(c) >="A" && string(c)<= "Z"{
				cc := string(c)
				s += cc
			}

		}
		dict[k]=s
		//fmt.Print(k," ",s,"\n")

	}
	//for _,word :=range  wordPairs{
	//	pair := strings.Split(word,"\t")
	//	k,_ := strconv.Atoi(pair[0])
	//	dict[k]=pair[1]
	//	fmt.Print(pair[0],pair[1])
	//
	//}

}

func GetPinYin(s string) string{
	textQuoted := strconv.QuoteToASCII(s)
	t := strings.Split(textQuoted,"\\u")
	i,_:=strconv.Atoi(t[0])
	
	fmt.Print(dict[i])
	return textQuoted
}
func main(){
	GetPinYin("你好")

}