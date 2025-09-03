package main

import "reflect"

func main() {

}

// SliceToMapSlice struct切片转为map切片
// 用gconv.Structs可以解决`map切片`转`struct切片`
func SliceToMapSlice(source interface{}) []map[string]interface{} {
	// 判断，interface转为[]interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		panic("ERROR: Unknown type, slice expected.")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	// 转换之后的结果变量
	res := make([]map[string]interface{}, 0)

	// 通过遍历，每次迭代将struct转为map
	for _, elem := range ret {
		data := make(map[string]interface{})
		objT := reflect.TypeOf(elem)
		objV := reflect.ValueOf(elem)
		for i := 0; i < objT.NumField(); i++ {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
		res = append(res, data)
	}

	return res
}
