/*
 * @Author:
 * @Date: 2024-06-30 10:54:42
 * @LastEditors:
 * @LastEditTime: 2024-06-30 10:55:37
 * @Description:
 */
package main

import "go-admin/router"

func main() {
	r := router.App()
	r.Run(":8081")
}
