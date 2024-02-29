/*
Copyright 2021 The DnsJia Authors.
WebSite:  https://github.com/dnsjia/luban
Email:    OpenSource@dnsjia.com

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

package routers

import (
	"github.com/foce123/kubespace/controller"
	"github.com/gin-gonic/gin"
)

func InitCasBinRouter(Router *gin.RouterGroup) {

	CasBinRouter := Router.Group("casbin")
	{
		CasBinRouter.POST("", controller.AddCasBin)
	}
}
