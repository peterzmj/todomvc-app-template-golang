package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todomvc-app-template-golang/db"
	"todomvc-app-template-golang/handler"
	)
//1. 新增一个任务
//2. 删除一个任务
//3. 标记【一个任务】为【完成状态】或者为【未完成状态】  标记全部任务为完成或未完成
//4. 查询全部任务 ，查询全部完成状态的任务，查询全部未完成状态的任务
//5. 按照关键词查询任务，按照关键词查询任务并且任务状态为完成，按照关键词查询任务并且任务状态为未完成

//定义一个服务器,监听并接收信息
func initServer(engine *gin.Engine) *http.Server{
	return &http.Server{
		Addr:              ":8080",
		Handler:           engine,
		ReadTimeout:       10 *time.Second,
		WriteTimeout:      10 *time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
//定义一个路由器组
func initRouter(engine *gin.Engine) {
	r := engine.Group("api")
	{
		r.POST("add", handler.Add)
		r.POST("del", handler.Del)
		r.POST("update", handler.Update)
		r.POST("findstatus", handler.FindStatus)
		r.POST("finditem", handler.FindItem)
	}
}
func main(){
	engine := gin.Default()
	engine.Use()
	initRouter(engine)
	server := initServer(engine)
	db.InitDB()
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}