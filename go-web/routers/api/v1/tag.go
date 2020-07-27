package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	_ "github.com/hearecho/go-pro/go-web/docs"
	"github.com/hearecho/go-pro/go-web/models"
	"github.com/hearecho/go-pro/go-web/pkg/resp"
	"github.com/hearecho/go-pro/go-web/pkg/setting"
	"github.com/hearecho/go-pro/go-web/pkg/utils"
	"github.com/unknwon/com"
)

// @Summary 获取全部标签
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} resp.R
// @Failure 500 {object} resp.R
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state = 1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	data["lists"] = models.GetTags(utils.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(200, resp.R{}.Ok().SetPath(c.Request.URL.Path).SetData(data))
}

// @Summary 添加标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} resp.R
// @Failure 500 {object} resp.R
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	r := resp.R{}.SetPath(c.Request.URL.Path)
	if ! valid.HasErrors() {
		if ! models.ExitTagByName(name) {
			r = r.Ok()
			models.AddTag(name, state, createdBy)
		} else {
			r = r.SetStatus(resp.ERROR_EXIST_TAG)
		}
	}
	c.JSON(200, r)
}

// @Summary Update article tag
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} resp.R
// @Failure 500 {object} resp.R
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	r := resp.R{}.SetPath(c.Request.URL.Path).Ok()
	if ! valid.HasErrors() {
		if models.ExitTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			r = r.SetStatus(resp.ERROR_NOT_EXIST_TAG)
		}
	}
	c.JSON(200, r)
}

// @Summary Delete article tag
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} resp.R
// @Failure 500 {object} resp.R
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	r := resp.R{}.SetPath(c.Request.URL.Path).Ok()
	if ! valid.HasErrors() {
		if models.ExitTagByID(id) {
			models.DeleteTag(id)
		} else {
			r = r.SetStatus(resp.ERROR_NOT_EXIST_TAG)
		}
	}
	c.JSON(200, r)
}
