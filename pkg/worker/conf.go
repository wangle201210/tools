package worker

const BaseUrl  = "files/"

//允许下载的列表
var fileAble = []string{
	"css",
	"js",
	"jpg",
	"png",
	"gif",
	"jpeg",
	"eot",
	"ttf",
	"svg",
	"woff",
	"ico",
}
//需要被替换为相对链接的文件
var needRep = []string{
	"css",
	"js",
}
