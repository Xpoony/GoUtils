package config
// read any json config file or json file
import (
	j "encoding/json"
	e "errors"
	f "fmt"
	i "io/ioutil"
	r "reflect"
	sc "strconv"
	st "strings"
	s "sync"
)

var c *configs
var oc s.Once

type configs struct {
	data map[string]interface{}
}


func GetObject() *configs {
	oc.Do(func() {
		var err error
		c, err = c.init()
		if err != nil {
			f.Println(err)
		}
	})
	return c
}
func (c *configs) init() (*configs, error) {
	var data map[string]interface{}
	path, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	filedata, err := i.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = j.Unmarshal(filedata, &data)
	if err != nil {
		return nil, err
	}
	return &configs{data}, nil
}
func getConfigFilePath() (string, error) {
	dir_list, err := i.ReadDir(".")
	if err != nil {
		return "", err
	}
	names := make([]string, len(dir_list))
	n := 0
	var array []string
	var ftype string
	for _, f := range dir_list {
		if !f.IsDir() {
			array = st.Split(f.Name(), ".")
			ftype = array[len(array)-1]
			if ftype == "json" {
				names[n] = f.Name()
				n = n + 1
			}
		}
	}
	if names[1] != "" {
		return "", e.New("find too many config file")
	}
	return names[0], nil
}

func (c *configs) SearchValue(key ...interface{}) (interface{}, error) {
	var result interface{}
	result = c.data
	for index := 0; index < len(key); index++ {
		keyType := r.TypeOf(key[index]).String()
		switch keyType {
		case "string":
			{
				temp, ok := result.(map[string]interface{})
				if !ok {
					return nil, e.New("The target is no a map, key index:" + sc.Itoa(index))
				}
				str, _ := key[index].(string)
				result = temp[str]
			}
		case "int":
			{
				arr, ok := result.([]interface{})
				if !ok {
					return nil, e.New("The target is no a array, key index:" + sc.Itoa(index))
				}
				in, _ := key[index].(int)
				result = arr[in]
			}
		default:
			{
				return nil, e.New("wrong type of key")
			}
		}
	}
	return result, nil
}
