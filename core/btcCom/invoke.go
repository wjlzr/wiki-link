package btcCom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"wiki-link/common"
)

//根据方法，获取对应接口返回数据
//method  		方法名称
//result		获取返回值
//query 		查询条件
//params		参数
func Invoke(method string, result interface{}, query map[string]interface{}, params ...interface{}) error {
	//获取地址链接
	url := toUrl(Url[method], params...)

	//拼接带参数的链接
	if query != nil {
		url = url + "?" + getValues(query)
	}
	// fmt.Println(url)
	bs, err := getFrom(url)
	if err != nil {
		return err
	}
	// fmt.Println(string(bs))
	if err := json.Unmarshal(bs, &result); err != nil {
		return err
	}
	return nil
}

//获取数据
func getFrom(url string) ([]byte, error) {
	req, err := http.NewRequest(common.GET, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//拼接地址
func toUrl(url string, params ...interface{}) string {
	return fmt.Sprintf(url, params...)
}

func getValues(query map[string]interface{}) string {
	var buf bytes.Buffer
	for k, v := range query {
		switch v.(type) {
		case int:
			buf.WriteString("&" + k + "=" + strconv.Itoa(v.(int)))
		case string:
			buf.WriteString("&" + k + "=" + v.(string))
		case int64:
			buf.WriteString("&" + k + "=" + strconv.FormatInt(v.(int64), 10))
		}
	}
	return buf.String()[1:]
}
