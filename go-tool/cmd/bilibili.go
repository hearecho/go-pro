package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-tool/utils"
)

const (
	ModeUserInfo = iota + 1
	ModeVideoInfo
)
var bilidesc = ""
var name string


var bilibiliCmd = &cobra.Command{
	Use:"bili",
	Short:"查询bilibili相关信息",
	Long:bilidesc,
	Run: func(cmd *cobra.Command, args []string) {
		switch mode {
		case ModeUserInfo:
			users := utils.GetUserInfo(name)
			for _,v := range users {
				fmt.Printf("|%-20s|%-15s|%-6.0f|%-15.0f|%-2.0f|%-4s|%-4s|\n",v.Name,v.Mid,v.VideoNum,v.Fans,v.Level,v.PlayNum,v.Likes)
			}
		default:

		}
	},
}

func init()  {
	bilibiliCmd.Flags().StringVarP(&name,"name","n","bilibili","用户名称")
	bilibiliCmd.MarkFlagRequired("name")
}

