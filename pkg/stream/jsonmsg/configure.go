/*
   Copyright (C) Nippon Telegraph and Telephone Corporation.

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

package jsonmsg

import (
	"net"

	"github.com/norouter/norouter/pkg/version"
)

const (
	OpConfigure Op = "configure"
)

type ConfigureRequestArgs struct {
	Me       net.IP        `json:"me"` // Required
	Forwards []Forward     `json:"forwards,omitempty"`
	Others   []IPPortProto `json:"others,omitempty"`
}

type ConfigureResultData struct {
	Features []version.Feature `json:"features,omitempty"`
	Version  string            `json:"version,omitempty"`
}

type Forward struct {
	// listenIP is "me"
	ListenPort  uint16 `json:"listen_port"`
	ConnectIP   net.IP `json:"connect_ip"`
	ConnectPort uint16 `json:"connect_port"`
	Proto       string `json:"proto"`
}

type IPPortProto struct {
	IP    net.IP `json:"ip"`
	Port  uint16 `json:"port"`
	Proto string `json:"proto"`
}
