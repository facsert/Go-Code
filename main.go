/*
 * @Author: facsert
 * @Date: 2023-08-01 21:41:41
 * @LastEditTime: 2023-08-06 22:17:11
 * @LastEditors: facsert
 * @Description:
 */

package main

import (
    "learn/common"
	_ "learn/flags"
	_ "learn/logger"
	_ "learn/file"
	_ "time"
)

func main() {
    // common.Title(" first title", 0)
	common.Title("first title", 5)
	// common.Title(" first title", 2)
	// common.Title(" first title", 3)
}
