/*
 * @Author: yhlyl
 * @Date: 2019-10-15 11:52:22
 * @LastEditTime: 2019-11-04 21:21:26
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/model/heartbeat.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package model

// 心跳
type Heartbeat struct {
	TimeStr string `json:"heart_beat"`
}
