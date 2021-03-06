/*
 * Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 *
 */

package resources

import (
	"reflect"

	"github.com/gardener/controller-manager-library/pkg/utils"
)

var cluster_key reflect.Type

func init() {
	cluster_key, _ = utils.TypeKey((*Cluster)(nil))
}

// _object is the standard implementation of the Object interface
// it uses the AbstractObject as base to provide standard implementations
// based on the internal object interface. (see _i_object)
type _object struct {
	AbstractObject
	resource Internal
}

var _ Object = &_object{}

func newObject(data ObjectData, resource Internal) Object {
	o := &_object{AbstractObject{}, resource}
	o.AbstractObject = NewAbstractObject(&_i_object{o}, data, resource.Resource())
	return o
}

func (this *_object) DeepCopy() Object {
	data := this.ObjectData.DeepCopyObject().(ObjectData)
	return newObject(data, this.resource)
}
