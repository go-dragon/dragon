package dto

import (
	"github.com/go-dragon/util"
)

// 根据结构体的json标签将结构体转为map
func TStructToData(obj interface{}, keys []string) map[string]interface{} {
	res := util.StructJsonTagToMap(obj)
	return util.OnlyColumns(keys, res)
}

// todo 一般遍历自己转换
//func TStructsToListData(objs []interface{}, keys []string) ListData {
//	output := ListData{}
//	for _, v := range objs {
//		output = append(output, TStructToData(v, keys))
//	}
//	return output
//}
