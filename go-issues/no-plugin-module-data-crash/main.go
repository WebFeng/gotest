package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"plugin"
)

func main() {

	const pluginAFilePath = "a.so"
	const pluginBFilePath = "b.so"
	const pluginBCopiedPath = "b-copied.so"

	{
		// 打开a插件
		_, err := plugin.Open(pluginAFilePath)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	{
		// 打开b插件,报错, 依赖包版本不一致
		_, err := plugin.Open(pluginBFilePath)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	{
		// 再次打开b插件, 再次报错
		_, err := plugin.Open(pluginBFilePath)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	{
		// 拷贝b插件
		buf, err := ioutil.ReadFile(pluginBFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err = ioutil.WriteFile(pluginBCopiedPath, buf, 0644); err != nil {
			fmt.Println(err)
			return
		}
	}

	{
		// 打开拷贝版b插件,将会crash
		_, err := plugin.Open(pluginBCopiedPath)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	// 执行不到
	_ = os.Remove("b-copy.so")

}
