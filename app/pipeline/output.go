package pipeline

import (
	"sort"

	"qtrx.io/lycosa/app/pipeline/collector"
	"qtrx.io/lycosa/common/kafka"
	"qtrx.io/lycosa/common/mgo"
	"qtrx.io/lycosa/common/mysql"
	"qtrx.io/lycosa/runtime/cache"
)

// 初始化输出方式列表collector.DataOutputLib
func init() {
	for out, _ := range collector.DataOutput {
		collector.DataOutputLib = append(collector.DataOutputLib, out)
	}
	sort.Strings(collector.DataOutputLib)
}

// 刷新输出方式的状态
func RefreshOutput() {
	switch cache.Task.OutType {
	case "mgo":
		mgo.Refresh()
	case "mysql":
		mysql.Refresh()
	case "kafka":
		kafka.Refresh()
	}
}
