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
			fmt.Println("|用户名 |用户ID |投稿数 |播放数 |点赞数 |关注数 |粉丝数 |")
			fmt.Printf("|%-7s|%-7s|%-7.0f|%-7s|%-7s|%-7.0f|%-7.0f|\n",u.Name,u.Mid,u.VideoNum,u.PlayNum,u.Likes,u.Following,u.Follower)
		default:

		}
	},
}

func init()  {
	bilibiliCmd.Flags().StringVarP(&mid,"mid","","6089090","用户mid号")
	bilibiliCmd.MarkFlagRequired("mid")
}

