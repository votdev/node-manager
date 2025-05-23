/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/harvester/node-manager/pkg/apis/node.harvesterhci.io/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	CloudInit() CloudInitController
	Ksmtuned() KsmtunedController
	NodeConfig() NodeConfigController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) CloudInit() CloudInitController {
	return generic.NewNonNamespacedController[*v1beta1.CloudInit, *v1beta1.CloudInitList](schema.GroupVersionKind{Group: "node.harvesterhci.io", Version: "v1beta1", Kind: "CloudInit"}, "cloudinits", v.controllerFactory)
}

func (v *version) Ksmtuned() KsmtunedController {
	return generic.NewNonNamespacedController[*v1beta1.Ksmtuned, *v1beta1.KsmtunedList](schema.GroupVersionKind{Group: "node.harvesterhci.io", Version: "v1beta1", Kind: "Ksmtuned"}, "ksmtuneds", v.controllerFactory)
}

func (v *version) NodeConfig() NodeConfigController {
	return generic.NewController[*v1beta1.NodeConfig, *v1beta1.NodeConfigList](schema.GroupVersionKind{Group: "node.harvesterhci.io", Version: "v1beta1", Kind: "NodeConfig"}, "nodeconfigs", true, v.controllerFactory)
}
