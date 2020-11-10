/*
Copyright 2020 The Nocalhost Authors.
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

package user

import (
	"context"

	"nocalhost/internal/nocalhost-api/service"
	"nocalhost/pkg/nocalhost-api/app/api"
	"nocalhost/pkg/nocalhost-api/pkg/errno"
	"nocalhost/pkg/nocalhost-api/pkg/log"

	"github.com/gin-gonic/gin"
)

// Get 获取用户信息
// @Summary 通过用户id获取用户信息
// @Description Get an user by user id
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")

	userID, _ := c.Get("userId")
	if userID == 0 {
		api.SendResponse(c, errno.ErrParam, nil)
		return
	}

	// Get the user by the `user_id` from the database.
	u, err := service.Svc.UserSvc().GetUserByID(context.TODO(), userID.(uint64))
	if err != nil {
		log.Warnf("get user info err: %v", err)
		api.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	api.SendResponse(c, nil, u)
}

// Get 获取用户列表
// @Summary 获取用户列表
// @Description Get userlist
// @Tags 用户
// @Accept  json
// @Produce  json
// @param Authorization header string true "Authorization"
// @Success 200 {object} model.UserList "用户列表"
// @Router /v1/users/list [get]
func GetList(c *gin.Context) {
	u, _ := service.Svc.UserSvc().GetUserList(context.TODO())
	api.SendResponse(c, nil, u)
}