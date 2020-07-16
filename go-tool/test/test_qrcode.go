package main

import (
	"fmt"
	"go-tool/utils"
)

func main() {
	image,_ := utils.CreateAvatar("image/avatar.png","resu.png")
	fmt.Println(image)
}
