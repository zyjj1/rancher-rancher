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

package fake

import (
	v1 "github.com/rancher/rancher/pkg/apis/rke.cattle.io/v1"
	rkecattleiov1 "github.com/rancher/rancher/pkg/generated/clientset/versioned/typed/rke.cattle.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCustomMachines implements CustomMachineInterface
type fakeCustomMachines struct {
	*gentype.FakeClientWithList[*v1.CustomMachine, *v1.CustomMachineList]
	Fake *FakeRkeV1
}

func newFakeCustomMachines(fake *FakeRkeV1, namespace string) rkecattleiov1.CustomMachineInterface {
	return &fakeCustomMachines{
		gentype.NewFakeClientWithList[*v1.CustomMachine, *v1.CustomMachineList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("custommachines"),
			v1.SchemeGroupVersion.WithKind("CustomMachine"),
			func() *v1.CustomMachine { return &v1.CustomMachine{} },
			func() *v1.CustomMachineList { return &v1.CustomMachineList{} },
			func(dst, src *v1.CustomMachineList) { dst.ListMeta = src.ListMeta },
			func(list *v1.CustomMachineList) []*v1.CustomMachine { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.CustomMachineList, items []*v1.CustomMachine) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
