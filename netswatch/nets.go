// Copyright 2015 flannel authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netswatch

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/coreos/flannel/subnet/etcdv2"
)

type IP uint32

type NodeMeta struct {
	OrgName  string
	NodeType string
	NodeName string
	HostIP   net.IP
}

func Hello() {
	fmt.Println("一哭二闹三上悠亚")
}

func WatchNets(ctx context.Context, sm *etcdv2.LocalManager) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done netswatch")
			return
		default:
			fmt.Println("watching nets")
			fmt.Printf("%T\n", sm.GetSubnets(ctx))
			time.Sleep(2 * time.Second)
			// case <-time.After(2 * time.Second):
			// 	fmt.Println("1024")
		}

	}
}

func ExtendNodeMeta(meta *NodeMeta) *NodeMeta {
	// If meta.NodeName is not set, then use hostname for node name.
	if len(meta.NodeName) == 0 {
		name, err := os.Hostname()
		if err != nil {
			fmt.Println("get hostname error")
			fmt.Printf("%v", err)
			name = "default-node"
		}
		meta.NodeName = name
	}

	return meta
}
