/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var mode int8
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-tool",
	Short: "go命令行自定义工具包",
	Long: `go命令行自定义工具包，包括简单curl实现，生成二维码，格式化json文件等等工具`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowHelp(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-tool.yaml)")
	rootCmd.PersistentFlags().Int8VarP(&mode, "mode", "m", 1, "请输入单词转换的模式")
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(qrcodeCmd)
	rootCmd.AddCommand(bilibiliCmd)
}

func initConfig() {
	//读取配置文件， 如果没有
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-tool.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-tool")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
