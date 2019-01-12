package worker

import (
	"regexp"
	"strings"
)

// 获取首页所有该要的链接
//func GetInfo(index []byte) (files []string) {
//	//css 等
//	link := regexp.MustCompile(`<link[^>]*>`)
//	linkMatchs := link.FindAllSubmatch(index,-1)
//	for _,m := range linkMatchs{
//		href := regexp.MustCompile(`href=[\"\']([^\"^\'^\?]*)[\"\'\?]`)
//		hrefMatchs := href.FindSubmatch(m[0])
//		if len(hrefMatchs) < 2 {
//			continue
//		}
//		files=append(files,string(hrefMatchs[1]))
//	}
//	//js等
//	jsLink := regexp.MustCompile(`<script[^>]*>`)
//	jsLinkMatchs := jsLink.FindAllSubmatch(index,-1)
//	for _,m := range jsLinkMatchs{
//		js := regexp.MustCompile(`src=[\"\']([^\"^\'^\?]*)[\"\'\?]`)
//		jsMatchs := js.FindSubmatch(m[0])
//		if len(jsMatchs) < 2 {
//			continue
//		}
//		files = append(files,string(jsMatchs[1]))
//	}
//	//img等
//
//	//<img width="300" height="142" src="/attachment/core/46/2018_10/26_17/26750561b9a57e81.jpg" title="广告" border="0">
//	//<a href="./455.html"><img src="./wp-content/themes/begin/timthumb.php?src=/wp-content/uploads/2017/03/coscmd_configure.png&w=280&h=210&a=&zc=1" alt="腾讯云COS上传、批量删除工具(Python)" /></a>						<span class="cat"><a href="./category/python/">Python</a></span>
//
//	imgLink := regexp.MustCompile(`<[Ii][Mm][Gg][^>]*>`)
//	imgLinkMatchs := imgLink.FindAllSubmatch(index,-1)
//	for _,m := range imgLinkMatchs{
//		//"XXXX"
//		img := regexp.MustCompile(`[\"\']([^\"^\']*)[\"\']`)
//		imgMatchs := img.FindAllSubmatch(m[0],-1)
//		for _,i := range imgMatchs {
//			files = append(files,string(i[1]))
//		}
//	}
//	return
//}
func GetInfo(index []byte) (files []string) {
	//href="index/fw/bsdt.css"
	href := regexp.MustCompile(`href=[\"\']([^\"^\']*)[\"\']`)
	linkMatchs := href.FindAllSubmatch(index,-1)
	for _,m := range linkMatchs{
		if len(m) < 2 {
			continue
		}
		files=append(files,string(m[1]))
	}
	//js等

	//src="images/close.png"
	src := regexp.MustCompile(`src=[\"\']([^\"^\']*)[\"\']`)
	srcMatchs := src.FindAllSubmatch(index,-1)
	for _,m := range srcMatchs{
		if len(m) < 2 {
			continue
		}
		files=append(files,string(m[1]))
	}

	//css 等
	cssLink := regexp.MustCompile(`<link[^>]*>`)
	cssLinkMatchs := cssLink.FindAllSubmatch(index,-1)
	for _,m := range cssLinkMatchs{
		css := regexp.MustCompile(`href=[\"\']([^\"^\'^\?]*)[\"\'\?]`)
		cssMatchs := css.FindSubmatch(m[0])
		if len(cssMatchs) < 2 {
			continue
		}
		files=append(files,string(cssMatchs[1]))
	}
	//js等
	jsLink := regexp.MustCompile(`<script[^>]*>`)
	jsLinkMatchs := jsLink.FindAllSubmatch(index,-1)
	for _,m := range jsLinkMatchs{
		js := regexp.MustCompile(`src=[\"\']([^\"^\'^\?]*)[\"\'\?]`)
		jsMatchs := js.FindSubmatch(m[0])
		if len(jsMatchs) < 2 {
			continue
		}
		files = append(files,string(jsMatchs[1]))
	}
	//img等

	//<img width="300" height="142" src="/attachment/core/46/2018_10/26_17/26750561b9a57e81.jpg" title="广告" border="0">
	//<a href="./455.html"><img src="./wp-content/themes/begin/timthumb.php?src=/wp-content/uploads/2017/03/coscmd_configure.png&w=280&h=210&a=&zc=1" alt="腾讯云COS上传、批量删除工具(Python)" /></a>						<span class="cat"><a href="./category/python/">Python</a></span>
	imgLink := regexp.MustCompile(`<[Ii][Mm][Gg][^>]*>`)
	imgLinkMatchs := imgLink.FindAllSubmatch(index,-1)
	for _,m := range imgLinkMatchs{
		img := regexp.MustCompile(`[\"\']([^\"^\']*)[\"\']`)
		imgMatchs := img.FindAllSubmatch(m[0],-1)
		for _,i := range imgMatchs {
			files = append(files,string(i[1]))
		}
	}
	// 去重
	files = RemoveRep(files)

	return
}
//子文件内的链接
func GetFileInfo(c []byte) (files []string) {
	//url('iconfont.woff?t=1501323857226')
	imgLink := regexp.MustCompile(`url\([\'\"]([^\"^\']*)[\"\']\)`)
	imgLinkMatchs := imgLink.FindAllSubmatch(c,-1)
	for _,m := range imgLinkMatchs{
		files = append(files,string(m[1]))
	}
	img2Link := regexp.MustCompile(`url\(([A-Za-z0-9\.\_\-\\\/\?\#\&\=]*)\)`)
	img2LinkMatchs := img2Link.FindAllSubmatch(c,-1)
	for _,m := range img2LinkMatchs{
		files = append(files,string(m[1]))
	}
	files = RemoveRep(files)
	return
}
// 看下开头是http 还是https
func header(url string) (h string) {
	pName := regexp.MustCompile(`([^:]*)://([A-Za-z0-9.]*)`)
	p := pName.FindSubmatch([]byte(url))
	h = string(p[1])
	return
}

func NeedRep(fileName string,c []byte) (r []byte) {
	canRep := false
	for _,able := range needRep {
		if strings.Contains(fileName, "."+able) {
			canRep = true
		}
	}
	if canRep {
		r = []byte(strings.Replace(string(c), `'/`, `'./`, -1))
		r = []byte(strings.Replace(string(r), `"/`, `"./`, -1))
	} else {
		r = c
	}
	return

}
func RemoveRep(s []string) []string {
	result := []string{}
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return reduce(result)
}

//去除多余标志比如？&#@等
func reduce(s []string) (r []string) {
	for _,f := range s{
		re := regexp.MustCompile(`[A-Za-z0-9\/\_\.\-\\]*`)
		ma := re.FindSubmatch([]byte(f))
		r = append(r,string(ma[0]))
	}
	return
}