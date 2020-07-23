package cmd

import (
	"fmt"
	"github.com/hearecho/go-pro/go-tool/utils"
	"github.com/spf13/cobra"
)

var url string
var path string
var avatar string

var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "生成二维码",
	Long:  `将输入url转化为二维码进行保存`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := utils.CreateAvatar(avatar, path, url)
		if err != nil {
			fmt.Println("二维码生成失败,地址:", path)
		}
		fmt.Println("二维码生成成功,地址:", path)
	},
}

func init() {
	qrcodeCmd.Flags().StringVarP(&url, "url", "u", "", "二维码的内容")
	qrcodeCmd.MarkFlagRequired("url")
	qrcodeCmd.Flags().StringVarP(&path, "path", "p", "resu.png", "生成二维码地址")
	qrcodeCmd.Flags().StringVarP(&avatar, "avatar", "a", "image/avatar.png", "贴图地址,默认使用小恐龙")
}
