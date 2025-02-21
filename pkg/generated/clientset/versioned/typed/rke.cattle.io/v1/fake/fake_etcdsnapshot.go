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

// fakeETCDSnapshots implements ETCDSnapshotInterface
type fakeETCDSnapshots struct {
	*gentype.FakeClientWithList[*v1.ETCDSnapshot, *v1.ETCDSnapshotList]
	Fake *FakeRkeV1
}

func newFakeETCDSnapshots(fake *FakeRkeV1, namespace string) rkecattleiov1.ETCDSnapshotInterface {
	return &fakeETCDSnapshots{
		gentype.NewFakeClientWithList[*v1.ETCDSnapshot, *v1.ETCDSnapshotList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("etcdsnapshots"),
			v1.SchemeGroupVersion.WithKind("ETCDSnapshot"),
			func() *v1.ETCDSnapshot { return &v1.ETCDSnapshot{} },
			func() *v1.ETCDSnapshotList { return &v1.ETCDSnapshotList{} },
			func(dst, src *v1.ETCDSnapshotList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ETCDSnapshotList) []*v1.ETCDSnapshot { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ETCDSnapshotList, items []*v1.ETCDSnapshot) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
