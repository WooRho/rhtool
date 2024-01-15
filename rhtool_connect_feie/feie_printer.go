package printer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type IFeieParam interface {
	AddPrinter(snList string) (err error)
	Print(content string, sn string, times string) (orderId string, err error)
	Delete(sns string) (err error)
	PrinterStatus(sn string) (err error)
	IsPrintOk(strorderid string)
	PrinterLog(sn string, strdate string)
}

func NewFeiePrinter(user, ukey, furl string) IFeieParam {
	return &FeieParam{
		User: user,
		Ukey: ukey,
		Furl: furl,
	}
}

//清空待打印队列，Open_delPrinterSqs

// 添加打印机
func (f *FeieParam) AddPrinter(snList string) (err error) {
	var (
		resp = &PrinterStatusRespAdd{}
	)

	if strings.Trim(snList, " ") == "" {
		err = errors.New("请传入打印机信息")
		return
	}
	client, postValues := f.postValuesInit("Open_printerAddlist")
	postValues.Add("printerContent", snList) //打印机

	_res, _ := client.PostForm(f.Furl, postValues)
	if _res == nil {
		err = errors.New("url 可能存在错误")
		return
	}
	data, _ := ioutil.ReadAll(_res.Body)

	_res.Body.Close()

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	if resp.Msg != "ok" {
		err = errors.New(resp.Msg)
		return
	}
	if len(resp.Data.No) > 0 {
		errStr := resp.Data.No[0]
		err = errors.New(errStr)
		return
	}
	if len(resp.Data.NoGuide) > 0 {
		err = errors.New(resp.Data.NoGuide[0].Sn)
		return
	}

	return
}

// 打印
func (f *FeieParam) Print(content string, sn string, times string) (orderId string, err error) {

	var (
		resp = &PrinterStatusResp{}
	)

	client, postValues := f.postValuesInit("Open_printLabelMsg")
	postValues.Add("sn", sn)            //打印机编号
	postValues.Add("contents", content) //打印内容
	postValues.Add("times", times)      //打印次数

	_res, _ := client.PostForm(f.Furl, postValues)
	if _res == nil {
		err = errors.New("url 可能存在错误")
		return
	}
	data, _ := ioutil.ReadAll(_res.Body)
	_res.Body.Close()

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	if resp.Msg != "ok" {
		err = errors.New(resp.Msg)
		return
	}
	//orderId = resp.Data
	return
}

// 删除  打印机编号，多台打印机请用减号“-”连接起来。 （目前单删除）
func (f *FeieParam) Delete(sns string) (err error) {
	var (
		resp     = &PrinterStatusRespAdd{}
		_errorSn = make([]string, 0)
	)

	client, postValues := f.postValuesInit("Open_printerDelList")
	postValues.Add("snlist", sns) //订单ID由方法1返回

	res, _ := client.PostForm(f.Furl, postValues)
	if res == nil {
		err = errors.New("url 可能存在错误")
		return
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	if resp.Msg != "ok" {
		err = errors.New(resp.Msg)
		return
	}

	if len(resp.Data.No) > 0 {
		for _, v := range resp.Data.No {
			_errorSn = append(_errorSn, v)
		}
		err = errors.New(resp.Data.No[0])
		return
	}

	return
}

// 查询某订单是否打印成功
func (f *FeieParam) IsPrintOk(strorderid string) {
	client, postValues := f.postValuesInit("Open_queryOrderState")
	postValues.Add("orderid", strorderid) //订单ID由方法1返回

	res, _ := client.PostForm("", postValues)
	if res == nil {
		errors.New("url 可能存在错误")
		return
	}
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()
}

// 查询指定打印机某天的订单详情(推荐存本地)
func (f *FeieParam) PrinterLog(sn string, strdate string) {
	client, postValues := f.postValuesInit("Open_printMsg")
	postValues.Add("sn", sn)        //打印机编号
	postValues.Add("date", strdate) //日期字符串

	res, _ := client.PostForm(f.Furl, postValues)
	if res == nil {
		errors.New("url 可能存在错误")
		return
	}
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()
}

// 查询打印机的状态
func (f *FeieParam) PrinterStatus(sn string) (err error) {
	var (
		resp = &PrinterStatusResp{}
	)

	client, postValues := f.postValuesInit("Open_queryPrinterStatus")
	postValues.Add("sn", sn) //打印机编号

	res, _ := client.PostForm(f.Furl, postValues)
	if res == nil {
		err = errors.New("url 可能存在错误")
		return
	}

	data, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	//str := string(data)
	if !strings.Contains(resp.Data, "状态正常") {
		err = errors.New("打印机状态错误")
		return err
	}
	res.Body.Close()
	return
}

func (f *FeieParam) postValuesInit(apiName string) (http.Client, url.Values) {
	var (
		itime      = time.Now().Unix()
		stime      = strconv.FormatInt(itime, 10)
		sig        = SHA1(f.User + f.Ukey + stime) //生成签名
		client     = http.Client{}
		postValues = url.Values{}
	)

	postValues.Add("user", f.User)     //账号名
	postValues.Add("stime", stime)     //当前时间的秒数，请求时间
	postValues.Add("sig", sig)         //签名
	postValues.Add("apiname", apiName) //固定

	fmt.Println(f.User)
	fmt.Println(f.Ukey)
	fmt.Println(f.Furl)

	return client, postValues
}
