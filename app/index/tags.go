package index

import (
	"github.com/gin-gonic/gin"
	"myWeb/model"
	"net/http"
)

func TagsIndex(c *gin.Context) {
	var optPage model.OptPage
	optPage.Length = 10
	optPage.Start = 0
	var searchWhere model.SearchWhere
	id := c.DefaultQuery("id", "")
	searchWhere.TagId = id
	article, total := model.ApiArticleList(&searchWhere, &optPage)
	assginData := model.MergeAssignData(c, gin.H{
		"articles": article,
		"total":    total,
	})
	c.HTML(http.StatusOK, "index/tags/index.html", assginData)

}
