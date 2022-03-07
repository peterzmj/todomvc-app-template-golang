package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todomvc-app-template-golang/db"
	"todomvc-app-template-golang/model"
)
//1. 新增一个任务
func Add(c *gin.Context) {
	var p model.TodomvcAdd
	c.ShouldBindJSON(&p)
	var m = &model.Todomvc{Item: p.Item, Status: 0}
	db.DB.Create(&m)
	c.JSON(http.StatusOK, nil)
}
//2. 删除一个任务
func Del(c *gin.Context) {
	var p model.TodomvcDel
	c.ShouldBindJSON(&p)
	db.DB.Delete(&model.Todomvc{}, p.Id)
	c.JSON(http.StatusOK, nil)
}
//3. 标记【一个任务】为【完成状态】或者为【未完成状态】,更新任务内容
func Update (c *gin.Context) {
	var p []model.TodomvcUpdate
	c.ShouldBindJSON(&p)
	for _, t := range p {
		db.DB.Model(&model.Todomvc{}).Where("id", t.Id).Select("item","status","updated_at").Updates(model.Todomvc{
			Item:   t.Item,
			Status: t.Status,
		})
	}
	c.JSON(http.StatusOK, nil)
}
//4. 查询全部完成状态的任务，查询全部未完成状态的任务，查询全部任务
func FindStatus (c *gin.Context) {
	var p model.TodomvcFindStatus
	c.ShouldBindJSON(&p)
	var m []model.Todomvc
	var tx =db.DB
	if p.Status == -1 {
		tx.Find(&m)
	}else{
		tx = tx.Where("status", p.Status)
		tx.Find(&m)
	}
	c.JSON(http.StatusOK, &m)
}
//5. 按照关键词查询任务，按照关键词查询任务并且任务状态为完成，按照关键词查询任务并且任务状态为未完成
func FindItem (c *gin.Context) {
	var p model.TodomvcFindItem
	c.ShouldBindJSON(&p)
	var m []model.Todomvc
	var tx =db.DB
	if p.Status == -1 {
		tx = tx.Where("item LIKE ?", "%"+p.Item+"%")
		tx.Find(&m)
	}else{
		tx = tx.Where("item LIKE ? AND Status = ?", "%"+p.Item+"%", p.Status)
		tx.Find(&m)
	}
	c.JSON(http.StatusOK, &m)
}