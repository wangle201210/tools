package worker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var Wg sync.WaitGroup

func Imitate(url string) (filePath string) {
	u := []byte(url)
	pName := regexp.MustCompile(`[^:]*://([A-Za-z0-9.]*)`)
	p := pName.FindSubmatch(u)
	pn := string(p[1])
	//创建项目文件
	if e := os.MkdirAll(BaseUrl+pn,os.ModePerm); e != nil {
		fmt.Println(e)
	}
	// 获取首页内容并保存
	c := GetIndexAndCreateFile(url, pn+"/index.html")
	// 查找有哪些文件需要下载
	files := GetInfo(c)
	//files是匹配到的需要被下载的文件列表
	//url是传入的全网址
	//pn 是不好http的请求地址
	//最后一个参数记录当前地址
	CreateFiles(files,url,pn,"")
	Wg.Wait()
	filePath = BaseUrl+pn+".zip"
	Zip(BaseUrl+pn,filePath)
	return

}
// 首页处理
func GetIndexAndCreateFile(url string,path string) (c []byte)  {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified  {
		fmt.Println(err)
		panic(err)
	}
	c, err = ioutil.ReadAll(resp.Body)
	//都改为相对链接
	c = []byte(strings.Replace(string(c), url, "./", -1))
	c = []byte(strings.Replace(string(c), `'/`, `'./`, -1))
	c = []byte(strings.Replace(string(c), `"/`, `"./`, -1))
	check(err)
	_, ce := FileCreate(path)
	check(ce)
	err = ioutil.WriteFile(BaseUrl+path, c, 0644)
	check(err)
	return
}

//files是匹配到的需要被下载的文件列表
//url是传入的全网址
//pn 是不带http的请求地址
//最后一个参数记录当前地址
func CreateFiles(files []string,url string,pn string,nowPath string)  {

	for _,file := range files{
		//如果是以http开头的即是外链，不用下载（但是是自己网址下的就要下载）
		if len(file) > 7 && file[0:len(header(url))] == header(url) || file == "/" || file == "" || len(file) < 3 {
			if len(file) >= len(url) && string(file[0:len(url)]) == url {
				file = file[len(url):]
			}else {
				continue
			}
		}
		//判断是否在允许下载的列表里
		canDown := false
		for _,able := range fileAble {
			if strings.Contains(file, "."+able) {
				canDown = true
			}
		}

		if canDown {
			Wg.Add(1)
			go GetContentAndCreateFile(file,url,pn,nowPath)
		}
	}

}
// 生成文件并保存
// pn projectName 是项目存放的根目录
// url 是项目的首页网址
func GetContentAndCreateFile(file,url string,pn string,path string) (c []byte)  {
	defer Wg.Done()
	fmt.Printf("downloading:%s\n",header(url)+ "://"+pn+"/"+file)
	resp, err := http.Get(header(url)+ "://"+pn+"/"+file)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	c, err = ioutil.ReadAll(resp.Body)

	check(err)
	_, ce := FileCreate(pn+"/"+file)
	check(ce)
	c = NeedRep(file,c)
	err = ioutil.WriteFile(BaseUrl+pn+"/"+file, c, 0644)
	check(err)
	//深层处理
	files := GetFileInfo(c)
	split := strings.Split(pn+"/"+file, "/")
	fmt.Printf("need info is :%s\n",pn+"/"+file)
	newPn := strings.Join(split[:len(split)-1],"/")
	CreateFiles(files,url,newPn,path)
	//ch<-file
	return
}


// 判断文件夹是否存在
func FileCreate(filePath string) (f *os.File,e error) {
	fileInfo := strings.Split(filePath,"/")
	path := strings.Join(fileInfo[0:len(fileInfo)-1],"/")
	e = os.MkdirAll(BaseUrl+path,os.ModePerm)
	if e == nil {
		return os.Create(BaseUrl+filePath)
	}
	return nil,e
}

func check(e error) {
	if e != nil {
		fmt.Printf("this is a err:%s",e)
	}
}


