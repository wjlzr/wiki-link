package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/oldfritter/sidekiq-go"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"wiki-link/common"
	"wiki-link/config"
	"wiki-link/db"
)

var (
	worker sidekiq.WorkerI
	// emails = []string{"leon.zcf@gmail.com"}
	emails = []string{} // 接口不通，发email通知
)

func initWorker() {
	for _, w := range common.AllWorkers {
		if w.Name == "SendEmail" {
			worker = common.AllWorkerIs[w.Name](&w)
			worker.SetClient(db.RedisClient())
		}
	}
}

type Result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
}

//根据方法，获取对应接口返回数据
//method  		方法名称
//result		获取返回值
//query 		查询条件
//params		参数
func OKLinkInvoke(method string, r Result, query map[string]interface{}, params ...interface{}) error {
	//获取地址链接
	url := toOKLinkUrl(OKLinkUrl[method], params...)

	//拼接带参数的链接
	if query != nil {
		url = url + "?" + getValues(query)
	}
	// fmt.Println(url)

	bs, err := GetFromOKLink(url)
	if err != nil {
		return err
	}

	// fmt.Println(string(bs))

	if err := json.Unmarshal(bs, &r); err != nil {
		return err
	}

	if r.GetCode() != 0 {
		if r.GetCode() == 429 {
			return fmt.Errorf("OKLink %w", ErrVisitLimit)
		}
		err = fmt.Errorf(r.GetMsg())
	}
	return err
}

func OKLinkPost(method string, r Result, query map[string]interface{}, params ...interface{}) error {
	//获取地址链接
	url := toOKLinkUrl(OKLinkUrl[method], params...)

	// fmt.Println("url: ", url)
	// fmt.Println("query: ", query)

	bs, err := postFromOKLink(url, query)
	if err != nil {
		return err
	}

	// fmt.Println(string(bs))
	if err := json.Unmarshal(bs, &r); err != nil {
		return err
	}
	if r.GetCode() != 0 {
		if r.GetCode() == 429 {
			return fmt.Errorf("OKLink %w", ErrVisitLimit)
		}
		err = fmt.Errorf(r.GetMsg())
	}
	return err
}

//获取数据
func GetFromOKLink(url string) ([]byte, error) {
	req, err := http.NewRequest(common.GET, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(common.XApiKey, config.Conf().GetOkLinkApiKey())
	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true, //禁止
	}

	client := http.Client{
		Transport: &transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		notify500(url)
	}
	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func postFromOKLink(url string, query map[string]interface{}) ([]byte, error) {
	q, _ := json.Marshal(query)
	params := bytes.NewReader(q)
	req, err := http.NewRequest(common.POST, url, params)
	if err != nil {
		return nil, err
	}

	req.Header.Set(common.XApiKey, config.Conf().GetOkLinkApiKey())
	req.Header.Set("Content-Type", "application/json")

	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true, //禁止
	}

	client := http.Client{
		Transport: &transport,
	}

	// client := &http.Client{}
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

func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, time.Second*10)
	if err != nil {
		return conn, err
	}

	tcp_conn := conn.(*net.TCPConn)
	tcp_conn.SetKeepAlive(false)

	return tcp_conn, err
}

//拼接地址
func toOKLinkUrl(url string, params ...interface{}) string {
	return fmt.Sprintf(url, params...)
}

//获取query参数
func getUrlValues(info map[string]interface{}) url.Values {
	params := url.Values{}
	for k, v := range info {
		switch v.(type) {
		case int:
			params.Add(k, strconv.Itoa(v.(int)))
		case string:
			params.Add(k, v.(string))
		}
	}
	return params
}

//
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

func notify500(url string) {
	if worker == nil {
		initWorker()
	}
	for _, e := range emails {
		worker.Priority(map[string]string{"mailer": "notify", "method": "OKLink500", "locale": "zh-CN", "email": e, "content": url})
	}
}
