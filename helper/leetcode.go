package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// 生成文件
func GenerateFiles(rootPath, titleSlug string) error {
	detail, err := GetProblemItem(titleSlug)
	if err != nil {
		return err
	}

	// 文件夹名
	dirName := trap(detail.Question.QuestionId) + "." + detail.Question.TitleSlug
	// 文件夹路径
	dirPath := path.Join(rootPath, dirName)
	err = os.MkdirAll(dirPath, os.ModeDir)
	if err != nil {
		return err
	}

	var goContent string
	for _, snippet := range detail.Question.CodeSnippets {
		if snippet.Lang == "Go" {
			goContent = snippet.Code
			break
		}
	}

	goFile := detail.Question.Title + ".go"
	readmeFile := "readme.md"
	// 写入go文件
	err = ioutil.WriteFile(path.Join(dirPath, goFile), []byte("package code\r\n\r\n"+goContent), os.ModePerm)
	if err != nil {
		return err
	}

	// 写入md文件
	content := "#### 题目\r\n" + detail.Question.TranslatedContent + "\r\n\r\n #### 题解"
	err = ioutil.WriteFile(path.Join(dirPath, readmeFile), []byte(content), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 获取所有需要生成的文件
func GenerateAllFiles() {
	nums, items, err := GetProblemListsTranlation()
	if err != nil {
		return
	}
	fmt.Printf("will generate total:%d questions\n", nums)
	for _, item := range items {
		err = GenerateFiles("algorithms", item.TitleSlug)
		if err != nil {
			fmt.Printf("ERROR:generate file [%s,%s] error:%v\n", item.QuestionId, item.TitleSlug, err)
		} else {
			fmt.Printf("generate question:[%s,%s] success", item.QuestionId, item.TitleSlug)
		}
	}
}

func trap(s string) string {
	var result string
	ret := 5 - len(s)
	if ret > 0 {
		for i := 0; i < ret; i++ {
			result += "0"
		}
	}
	result += s
	return result
}
