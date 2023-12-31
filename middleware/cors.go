package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		
		var headerKeys []string
		for k := range context.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		
		headerStr := strings.Join(headerKeys, ", ")
		
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		
		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 允许访问所有域
			context.Header("Access-Control-Allow-Origin", "*")
			
			// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			
			//  header的类型
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,"+
				"session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, "+
				"X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, "+
				"Content-Type, Pragma")
			// 允许跨域设置 可以返回其他子段
			
			// 跨域关键设置 让浏览器可以解析
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			
			// 缓存请求信息 单位为秒
			context.Header("Access-Control-Max-Age", "172800")
			
			//  跨域请求是否需要带cookie信息 默认设置为true
			context.Header("Access-Control-Allow-Credentials", "false")
			
			// 设置返回格式是json
			context.Set("content-type", "application/json")
			
		}
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		context.Next() //  处理请求
		
	}
}
