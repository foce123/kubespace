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

package k8s

import (
	"net/http"

	"github.com/foce123/kubespace/controller"
	"github.com/foce123/kubespace/controller/response"
	"github.com/foce123/kubespace/models/k8s"
	"github.com/foce123/kubespace/pkg/k8s/Init"
	"github.com/foce123/kubespace/pkg/k8s/parser"
	"github.com/foce123/kubespace/pkg/k8s/statefulset"
	"github.com/gin-gonic/gin"
)

func GetStatefulSetListController(c *gin.Context) {
	client, err := Init.ClusterID(c)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}
	dataSelect := parser.ParseDataSelectPathParameter(c)
	nsQuery := parser.ParseNamespacePathParameter(c)

	data, err := statefulset.GetStatefulSetList(client, nsQuery, dataSelect)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}

	response.OkWithData(data, c)
	return
}

func DeleteCollectionStatefulSetController(c *gin.Context) {
	client, err := Init.ClusterID(c)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}
	var statefulSetList []k8s.StatefulSetData

	err = controller.CheckParams(c, &statefulSetList)
	if err != nil {
		response.FailWithMessage(http.StatusNotFound, err.Error(), c)
		return
	}

	err = statefulset.DeleteCollectionStatefulSet(client, statefulSetList)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func DeleteStatefulSetController(c *gin.Context) {
	client, err := Init.ClusterID(c)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}
	namespace := parser.ParseNamespaceParameter(c)
	name := parser.ParseNameParameter(c)

	err = statefulset.DeleteStatefulSet(client, namespace, name)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func RestartStatefulSetController(c *gin.Context) {
	client, err := Init.ClusterID(c)
	if err != nil {
		response.FailWithMessage(response.ParamError, err.Error(), c)
		return
	}
	var statefulSetData k8s.StatefulSetData
	err2 := controller.CheckParams(c, &statefulSetData)
	if err2 != nil {
		response.FailWithMessage(response.ParamError, err2.Error(), c)
		return
	}
	err3 := statefulset.RestartStatefulSet(client, statefulSetData.Name, statefulSetData.Namespace)
	if err3 != nil {
		response.FailWithMessage(response.InternalServerError, err3.Error(), c)
		return
	}
	response.Ok(c)
	return
}

func ScaleStatefulSetController(c *gin.Context) {
	client, err := Init.ClusterID(c)
	if err != nil {
		response.FailWithMessage(response.ParamError, err.Error(), c)
		return
	}
	var scaleData k8s.ScaleStatefulSet

	err2 := controller.CheckParams(c, &scaleData)
	if err2 != nil {
		response.FailWithMessage(response.ParamError, err2.Error(), c)
		return
	}

	err = statefulset.ScaleStatefulSet(client, scaleData.Namespace, scaleData.Name, *scaleData.ScaleNumber)
	if err != nil {
		response.FailWithMessage(response.InternalServerError, err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func DetailStatefulSetController(c *gin.Context) {
	client, err := Init.ClusterID(c)
	if err != nil {
		response.FailWithMessage(response.ParamError, err.Error(), c)
		return
	}
	namespace := parser.ParseNamespaceParameter(c)
	name := parser.ParseNameParameter(c)
	dataSelect := parser.ParseDataSelectPathParameter(c)

	data, err := statefulset.GetStatefulSetDetail(client, dataSelect, namespace, name)

	if err != nil {
		response.FailWithMessage(response.ERROR, err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}
