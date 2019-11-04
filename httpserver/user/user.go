/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:02:14
 * @LastEditTime: 2019-11-04 21:14:30
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/httpserver/user/user.go
 * @https://github.com/android-coco/gin_micro
 */
package user

import (
	"gin_micro/db"
	"gin_micro/httpserver/servermiddleware"
	"gin_micro/module"
	"gin_micro/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RegisteredReq 代理id,账号,手机号码,密码,ip,设备类型,机器序列号
type RegisteredReq struct {
	servermiddleware.BaseReq
	Account  string `form:"account" json:"account" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" `
	Password string `form:"password" json:"password" binding:"required"`
	AgentID  int    `form:"agent_id" json:"agent_id"`
}

//RegisteredRsp RegisteredRsp
type RegisteredRsp struct {
	UserID      int64   `json:"user_id"`      //uid
	UserGold    float64 `json:"user_gold"`    //金币
	UserDiamond int64   `json:"user_diamond"` //砖石
	PhoneNumber string  `json:"phone_number"` //手机号码
	MemberOrder int64   `json:"member_order"` //等级
}

//Registered 注册
func Registered(c *gin.Context) {
	var resReq RegisteredReq

	err := c.ShouldBindJSON(&resReq)

	if err != nil {
		util.Logger.Errorf("Registered 接口  参数绑定 出错 err: %s ", err.Error())
		c.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.ErrorLackCode, ErrorMsg: err.Error()})
		return
	}

	//执行注册的存储过程
	//(IN `strAccounts` varchar(32),IN `strPhoneNumber` varchar(11),IN
	// `strLogonPass` varchar(32),IN `numAgentID`
	// int,IN `numRegisterOrigin`
	// int,IN `strMachineID`
	// varchar(32),IN `strClientIP` varchar(15),OUT `errorCode` int,OUT `errorDescribe` varchar(127))
	rows, err := db.GetDB().Raw("CALL WSP_PW_RegisterAccounts(?,?,?,?,?,?,?)",
		resReq.Account,
		resReq.Mobile,
		resReq.Password,
		resReq.AgentID,
		resReq.NumRegisterOrigin,
		resReq.MachineID,
		strings.Split(c.Request.RemoteAddr, ":")[0],
	).Rows()
	if err != nil {
		util.Logger.Errorf("Registered 接口  查询存储过程 err: %s ", err.Error())
		c.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.ErrorSQLCode, ErrorMsg: err.Error()})
		return
	}
	var errorCode int64
	var errorMsg string
	rows.Next()
	err = rows.Scan(&errorCode, &errorMsg)
	if err != nil {
		util.Logger.Errorf("Registered 接口  查询存储过程 err: %s ", err.Error())
		c.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.ErrorSQLCode, ErrorMsg: err.Error()})
		return
	}
	if errorCode != util.SuccessCode {
		c.JSON(http.StatusOK, module.ApiResp{ErrorNo: errorCode, ErrorMsg: errorMsg})
		return
	}
	//查询映射到结构体
	var rsq RegisteredRsp
	rows.NextResultSet()
	rows.Next()
	//	SELECT 0 AS errorCode, '' AS errorMsg,
	//	paraUserID AS UserID,
	//	paraGoldCoin AS UserGold,
	//	paraDiamond  AS UserDiamond,
	//	strPhoneNumber AS PhoneNumber,
	//	paraLevelNum AS MemberOrder;
	err = rows.Scan(&rsq.UserID,
		&rsq.UserGold,
		&rsq.UserDiamond,
		&rsq.PhoneNumber,
		&rsq.MemberOrder)
	if err != nil {
		util.Logger.Errorf("Registered 接口  查询存储过程 err: %s ", err.Error())
		c.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.ErrorSQLCode, ErrorMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, Data: rsq})
}
