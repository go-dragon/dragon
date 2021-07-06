/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
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
 * limitations under the License.
 */

package model

import "time"

type Instance struct {
	Valid       bool              `json:"valid"`
	Marked      bool              `json:"marked"`
	InstanceId  string            `json:"instanceId"`
	Port        uint64            `json:"port"`
	Ip          string            `json:"ip"`
	Weight      float64           `json:"weight"`
	Metadata    map[string]string `json:"metadata"`
	ClusterName string            `json:"clusterName"`
	ServiceName string            `json:"serviceName"`
	Enable      bool              `json:"enabled"`
	Healthy     bool              `json:"healthy"`
	Ephemeral   bool              `json:"ephemeral"`
}

type Service struct {
	Dom             string            `json:"dom"`
	CacheMillis     uint64            `json:"cacheMillis"`
	UseSpecifiedURL bool              `json:"useSpecifiedUrl"`
	Hosts           []Instance        `json:"hosts"`
	Checksum        string            `json:"checksum"`
	LastRefTime     uint64            `json:"lastRefTime"`
	Env             string            `json:"env"`
	Clusters        string            `json:"clusters"`
	Metadata        map[string]string `json:"metadata"`
	Name            string            `json:"name"`
}

type ServiceDetail struct {
	Service  ServiceInfo `json:"service"`
	Clusters []Cluster   `json:"clusters"`
}

type ServiceInfo struct {
	App              string            `json:"app"`
	Group            string            `json:"group"`
	HealthCheckMode  string            `json:"healthCheckMode"`
	Metadata         map[string]string `json:"metadata"`
	Name             string            `json:"name"`
	ProtectThreshold float64           `json:"protectThreshold"`
	Selector         ServiceSelector   `json:"selector"`
}

type ServiceSelector struct {
	Selector string
}

type Cluster struct {
	ServiceName      string               `json:"serviceName"`
	Name             string               `json:"name"`
	HealthyChecker   ClusterHealthChecker `json:"healthyChecker"`
	DefaultPort      uint64               `json:"defaultPort"`
	DefaultCheckPort uint64               `json:"defaultCheckPort"`
	UseIPPort4Check  bool                 `json:"useIpPort4Check"`
	Metadata         map[string]string    `json:"metadata"`
}

type ClusterHealthChecker struct {
	Type string `json:"type"`
}

type SubscribeService struct {
	ClusterName string            `json:"clusterName"`
	Enable      bool              `json:"enable"`
	InstanceId  string            `json:"instanceId"`
	Ip          string            `json:"ip"`
	Metadata    map[string]string `json:"metadata"`
	Port        uint64            `json:"port"`
	ServiceName string            `json:"serviceName"`
	Valid       bool              `json:"valid"`
	Weight      float64           `json:"weight"`
}

type BeatInfo struct {
	Ip          string            `json:"ip"`
	Port        uint64            `json:"port"`
	Weight      float64           `json:"weight"`
	ServiceName string            `json:"serviceName"`
	Cluster     string            `json:"cluster"`
	Metadata    map[string]string `json:"metadata"`
	Scheduled   bool              `json:"scheduled"`
	Period      time.Duration     `json:"-"`
	Stopped     bool              `json:"-"`
}

type ExpressionSelector struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type ServiceList struct {
	Count int64    `json:"count"`
	Doms  []string `json:"doms"`
}
