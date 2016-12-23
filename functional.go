/*
+ 程序维护一个float64切片和一个字符串切片，并支持以下命令:
  + use [ float64 | string ]                          : 切换到整型切片或者 字符串切片为当前切片
  + append newElement1 newElement2  ...          : 追加新的元素
  + insert index newElement1  newElement2, ...    : 插入新的元素
  + delete fromIndex ［toIndex］                   : 删除元素
  + invert                                        : 反转
  + sort                                          : 排序
  + unique                                        : 去重
  以上命令执行完毕后，打印出整个当前切片。
  + find elementValue : find命令打印出被查找元素的下标。
*/
package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

//import "github.com/asiainfoLDP/datahub_commons/log"

func main() {
	salaries := make([]float64, 0, 30)
	names := make([]string, 0, 10)
	usenFloat := true
	reader := bufio.NewReader(os.Stdin)
	for {
		bs, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}
		line := string(bs)
		//fmt.Println(line, b, err)
		index := strings.Index(line, " ")
		if index < 0 {
			index = len(line)
		}
		command := line[:index]
		paramsText := line[index:]
		paramsText = strings.TrimSpace(paramsText)
		switch command {
		case "use":
			if paramsText == "float64" {
				usenFloat = true
				fmt.Println("切换至浮点数切片")
			} else if paramsText == "string" {
				usenFloat = false
				fmt.Println("切换至字符串")
			} else {
				fmt.Println("未知的切片类型")
			}

		case "append":
			// _ = salaries
			// _ = names
			//_ = usenFloat
			params := strings.Split(paramsText, " ")

			if usenFloat {
				for _, item := range params {
					item = strings.TrimSpace(item)
					if item == "" {
						continue
					}
					//fmt.Println(items)
					data, err := strconv.ParseFloat(item, 64)
					if err != nil {
						fmt.Println("转化为float64错误", err)
						break
					}
					salaries = append(salaries, data)
				}
				fmt.Println(salaries)
			} else {
				for _, item := range params {
					item = strings.TrimSpace(item)
					if item == "" {
						continue
					}
					names = append(names, item)
				}
				fmt.Println(strings.Join(names, ","))
			}

		case "insert":
			index := strings.Index(paramsText, " ")
			if index < 0 {
				index = len(paramsText)
			}
			insertindex, err := strconv.ParseInt(paramsText[:index], 10, 64)
			if err != nil {
				fmt.Println("插入下标转化为int错误", err)
				continue
			}
			params := strings.Split(paramsText[index:], " ")
			if usenFloat {
				if insertindex < 0 || insertindex > int64(len(salaries)) {
					fmt.Println("非法的下标")
					return
				}
				for _, item := range params {
					item = strings.TrimSpace(item)
					if item == "" {
						continue
					}
					data, err := strconv.ParseFloat(item, 64)
					if err != nil {
						fmt.Println("转化为float64错误", err)
						break
					}
					tail := append([]float64{}, salaries[insertindex:]...)
					head := append(salaries[:insertindex], data)
					salaries = append(head, tail...)
					// salaries = append(salaries, data)
				}
				fmt.Println(salaries)

			} else {
				for _, item := range params {
					item = strings.TrimSpace(item)
					if item == "" {
						continue
					}
					tail := append([]string{}, names[insertindex:]...)
					head := append(names[:insertindex], item)
					names = append(head, tail...)
					//names = append(names, item)

				}
				fmt.Println(strings.Join(names, ","))

			}

		case "delete":
			indexs := strings.Split(paramsText, " ")
			if len(indexs) > 2 {
				fmt.Println("下标参数过多")
			}
			stat_str := indexs[0]
			end_str := indexs[1]
			if usenFloat {
				startindex, err := strconv.ParseInt(stat_str, 10, 64)
				if err != nil {
					fmt.Println("开始删除下标转化为int64错误", err)
					continue
				}
				endindex, err := strconv.ParseInt(end_str, 10, 64)
				if err != nil {
					fmt.Println("结束删除下标转化为int64错误", err)
					continue
				}
				salaries = append(salaries[startindex:endindex], salaries[endindex+1:]...)
				fmt.Println(salaries)
			} else {
				startindex, err := strconv.ParseInt(stat_str, 10, 64)
				if err != nil {
					fmt.Println("开始删除下标转化为int64错误", err)
					continue
				}
				endindex, err := strconv.ParseInt(end_str, 10, 64)
				if err != nil {
					fmt.Println("结束删除下标转化为int64错误", err)
					continue
				}
				names = append(names[startindex:endindex], names[endindex+1:]...)
				fmt.Println(strings.Join(names, ","))
			}
		case "invert":
			if usenFloat {
				afte_invert := make([]float64, 0, 30)
				save := make([]float64, 0, 30)
				for i := len(salaries) - 1; i >= 0; i-- {
					afte_invert = append([]float64{}, salaries[i])
					save = append(save, afte_invert...)
				}

				fmt.Println(save)

			} else {
				afte_invert := make([]string, 0, 30)
				save := make([]string, 0, 30)
				for i := len(names) - 1; i >= 0; i-- {
					afte_invert = append([]string{}, names[i])
					save = append(save, afte_invert...)
				}
				fmt.Println(save)

			}
		case "sort":
			if usenFloat {
				for i := 1; i < len(salaries); i++ {
					for j := 0; j < i; j++ {
						if salaries[i] < salaries[j] {
							tem := salaries[i]
							salaries[i] = salaries[j]
							salaries[j] = tem
						}

					}
				}
				fmt.Println(salaries)

			} else {
				for i := 1; i < len(names); i++ {
					for j := 1; j < i; j++ {
						if names[i] < names[j] {
							tem := names[i]
							names[i] = names[j]
							names[j] = tem
						}

					}
				}
				fmt.Println(names)

			}

		case "unique":
			if usenFloat {
				for i := 1; i < len(salaries); i++ {
					for j := 0; j < i; j++ {
						if salaries[i] == salaries[j] {
							tail := append([]float64{}, salaries[i:]...)
							salaries = append(salaries[:i-1], tail...)
						}

					}
				}
				fmt.Println(salaries)

			} else {
				for i := 1; i < len(names); i++ {
					for j := 0; j < i; j++ {
						if names[i] == names[j] {
							tail := append([]string{}, names[i:]...)
							names = append(names[:i-1], tail...)

						}

					}
				}
				fmt.Println(names)
			}

		case "find":

			if usenFloat {
				param, err := strconv.ParseFloat(paramsText, 64)
				if err != nil {
					fmt.Println("find:转化为float64失败", err)
					continue
				}
				for i, ele := range salaries {
					if ele == param {
						fmt.Println(i)
					}
				}

			} else {
				for i, ele := range names {
					if ele == paramsText {
						fmt.Println(i)
					}
				}

			}

		}
	}
}
