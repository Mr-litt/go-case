package app

import (
	"net/http"
	"fmt"
	"strings"
)

// 获取短链接服务
func GetShortUrl(w http.ResponseWriter, r *http.Request) {

	// 获取长链接参数
	url := r.URL.Query().Get("url")
	if url == "" {
		fmt.Fprintln(w, "url参数错误")
		return
	}

	// 长链接转短链接并返回给客户端
	shortUrl := longToShort(url)
	if shortUrl == "" {
		fmt.Fprintln(w, "链接生成失败")
	}
	fmt.Fprintln(w, shortUrl)
}

// 跳转长链接服务
func RedirectLongUrl(w http.ResponseWriter, r *http.Request) {

	// 获取短链接路径
	shortPath := strings.TrimLeft(r.URL.Path, "/")
	if shortPath == "" {
		fmt.Fprintln(w, "短链接路由错误")
		return
	}

	// 路径转长链接并重定向
	longUrl := shortToLong(shortPath)
	http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
}
