/*
 * 说明：
 * 作者：zhe
 * 时间：2018-09-01 6:12 PM
 * 更新：
 */

package json_iterator

import (
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Node struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (n *Node) String() string {
	b, err := json.Marshal(n)
	if err != nil {
		logrus.Errorf("marshal error: %v", err)
		return ""
	}
	return string(b)
}

func (n *Node) Initialize(data []byte) *Node {
	if err := json.Unmarshal(data, n); err != nil {
		logrus.Errorf("unmarshal error: %v", err)
		return nil
	}
	return n
}
