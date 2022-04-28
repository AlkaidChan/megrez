package admin

import (
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/config"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/errmsg"
	"github.com/megrez/vendor/gopkg.in/yaml.v2"
)

func GetThemeConfig(c *gin.Context) {
	var cfg = &config.ThemeConfig{}
	home, err := dir.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	theme, err := model.GetOptionByKey(vo.OptionKeyBlogTheme)
	if err != nil {
		log.Error("get option theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	b, err := ioutil.ReadFile(path.Join(home, theme, "config.yaml"))
	if err != nil {
		log.Error("read theme config file failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	for _, tab := range cfg.Tabs {
		for _, item := range tab.Items {
			value := string(b)
			item.Value = 
		}
	}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		log.Error("unmarshal config failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(cfg))
}
