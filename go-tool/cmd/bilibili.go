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
var mid string


var bilibiliCmd = &cobra.Command{
	Use:"bili",
	Short:"查询bilibili相关信息",
	Long:bilidesc,
	Run: func(cmd *cobra.Command, args []string) {
		var u utils.UserInfo
		switch mode {
		case ModeUserInfo:
			u = utils.GetUserInfo(mid)
			fmt.Println(u)
		default:

		}
	},
}

func init()  {
	bilibiliCmd.Flags().StringVarP(&mid,"mid","","6089090","用户mid号")
	bilibiliCmd.MarkFlagRequired("mid")
}

