package shmscoring

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	urlFormat string = "http://%s:%s%s" // {0}=ip, {1}=port, {2}=endpoint
)

type ShmmConnector struct {
	ShmmAddrPorts map[string]string
}

func NewShmmConnector(addrs []string, ports []string) *ShmmConnector {
	s := &ShmmConnector{
		ShmmAddrPorts: make(map[string]string),
	}
	for idx, addr := range addrs {
		s.ShmmAddrPorts[addr] = ports[idx]
	}

	return s
}

func (s *ShmmConnector) GetNodeMetric(nodeName string) (NodeMetric, error) {
	result := NodeMetric{}

	if _, ok := s.ShmmAddrPorts[nodeName]; !ok {
		return result, fmt.Errorf("nodeName \"%s\" not found in ShmmConnector.", nodeName)
	}

	url := fmt.Sprintf(urlFormat, nodeName, s.ShmmAddrPorts[nodeName], "/metrics/mem")
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var jsonResult struct {
		Data metricsResp `json:"data"`
	}
	if err = json.Unmarshal(body, &jsonResult); err != nil {
		return result, err
	}

	result.ShmSum = jsonResult.Data.ShmSum
	result.ShmDiskUsed = jsonResult.Data.ShmDisk.Used
	result.ShmDiskFree = jsonResult.Data.ShmDisk.Free
	result.ShmDiskAll = jsonResult.Data.ShmDisk.All
	result.ModelCount = jsonResult.Data.ModelCount
	result.Models = make([]string, result.ModelCount)
	for i, m := range jsonResult.Data.Models {
		result.Models[i] = fmt.Sprintf("%s:%s", m.ModelName, m.TagName)
	}

	return result, nil
}
