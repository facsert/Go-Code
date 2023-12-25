/*
 * @Author: facsert
 * @Date: 2023-08-06 22:18:09
 * @LastEditTime: 2023-08-06 22:23:27
 * @LastEditors: facsert
 * @Description:
 */

package common

import (
	"fmt"
	"strings"
)

func Title(title string, level int) string {
	if level > 3 { level = 3 }
	if level < 0 { level = 0 }
	separator := [...]string{"#", "=", "*", "-"}[level]
	space := [...]string{"\n\n", "\n", "", ""}[level]

	line := strings.Repeat(separator, 80)
	fmt.Printf("%s%s %s %s\n", space, line, title, line)
	return title
}
