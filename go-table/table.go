package go_table

import (
	"fmt"
	"reflect"
	"strconv"
)

type Test struct {
	Name string
	Age  int
	Sex  string
}

func CreatTable(headers []string,data interface{}) {
	//输出表头
	var ls []int
	ls = getMaxLength(data)
	for i:=0;i<len(ls) ;i++  {
		if ls[i] < len(headers[i]) {
			ls[i] = len(headers[i])
		}
		fmt.Printf("|%-*s",ls[i],headers[i])
	}
	fmt.Printf("|\n")
	value := reflect.ValueOf(data)
	length := value.Len()
	fields := value.Index(0).NumField()
	for i:=0;i< length; i++ {
		for j:=0; j<fields;j++  {
			fmt.Printf("|%-*v",ls[j],value.Index(i).Field(j))
		}
		fmt.Printf("|\n")
	}
}

func getMaxLength(data interface{}) []int{
	var lengths []int
	value := reflect.ValueOf(data)
	if value.Type().Kind() == reflect.Slice {
		length := value.Len()
		fields := value.Index(0).NumField()
		for j:=0;j<fields;j++ {
			max := 0
			for i:=0;i<length;i++ {
				temp := value.Index(i).Field(j).Interface()
				l := getInterfaceLength(temp)
				if max < l {
					max = l
				}
			}
			lengths = append(lengths, max)
		}
	}
	return lengths
}

func getInterfaceLength(data interface{}) int {
	l := 0
	t := reflect.TypeOf(data).Kind()
	switch t {
	case reflect.String :
		l = len(data.(string))
	case reflect.Int:
		l = len(strconv.Itoa(data.(int)))
	case reflect.Float64:
		l = len(strconv.FormatFloat(data.(float64),'f',-1,64))
	case reflect.Uint:
		l = len(strconv.FormatUint(uint64(data.(uint)),10))
	default:
		l = 10
	}
	return l
}

func main() {
	t := Test{
		Name: "echo45645",
		Age:  20,
		Sex:  "男",
	}
	t1 := Test{
		Name: "ech11346456156",
		Age:  21,
		Sex:  "男",
	}
	t2 := Test{
		Name: "ech3",
		Age:  21,
		Sex:  "男",
	}
	data := []Test{t,t1,t2}
	heades := []string{"name","age","sex"}
	CreatTable(heades,data)
}
