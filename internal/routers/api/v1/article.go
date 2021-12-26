package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/koopa0/blog-service/pkg/app"
	"github.com/koopa0/blog-service/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// Get @Summary 取得單個文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// List @Summary 取得多個文章
// @Produce json
// @Param name  query string false "文章名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0,1) default
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {

}

// Create @Summary 新增文章
// @Produce json
// @Param name body string true "文章名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0,1) default(1)
// @Param created_by body string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {

}

// Update @Summary 更新文章
// @Produce json
// @Param id path int true "文章ID"
// @Param name body string false "文章名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {

}

// Delete @Summary 刪除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {

}
