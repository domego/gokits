package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	// 支持gif,jpeg,png的处理
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	// GZipMinSize 要启用gzip的最小值
	GZipMinSize = 1024
)

// ToGzipJSON 转换成json并gzip
func ToGzipJSON(obj interface{}) (gziped bool, data []byte, err error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return
	}
	if len(bs) <= GZipMinSize {
		return false, bs, nil
	}
	buf := &bytes.Buffer{}
	gzipWriter := gzip.NewWriter(buf)
	_, err = gzipWriter.Write(bs)
	gzipWriter.Close()
	if err != nil {
		return
	}
	return true, buf.Bytes(), nil
}

// ToGzip gzip
func ToGzip(bs []byte) (gziped bool, data []byte, err error) {
	if len(bs) <= GZipMinSize {
		return false, bs, nil
	}
	buf := &bytes.Buffer{}
	gzipWriter := gzip.NewWriter(buf)
	_, err = gzipWriter.Write(bs)
	gzipWriter.Close()
	if err != nil {
		return
	}
	return true, buf.Bytes(), nil
}

// FromGzip from gzip
func FromGzip(data []byte) (bs []byte, err error) {
	buf := bytes.NewBuffer(data)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return
	}
	defer gzipReader.Close()
	bs, err = ioutil.ReadAll(gzipReader)
	return
}

// FromGzipToJSON from gzip json
func FromGzipToJSON(data []byte, obj interface{}) (err error) {
	buf := bytes.NewBuffer(data)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return
	}
	defer gzipReader.Close()
	bs, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return
	}
	return json.Unmarshal(bs, obj)
}

// ToJSONString to json string
func ToJSONString(data interface{}) string {
	bs, err := json.Marshal(data)
	if err != nil {
		return "<can't serialize to json>"
	}
	return string(bs)
}

// JSONStringToMap json string to map
func JSONStringToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonBytes := []byte(jsonStr)
	dec := json.NewDecoder(bytes.NewReader(jsonBytes))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ToGBK to gbk
func ToGBK(s string) (string, error) {
	enc := simplifiedchinese.GBK
	trans := enc.NewEncoder()
	r, _, err := transform.String(trans, s)
	return r, err
}

// FromGBK from gbk
func FromGBK(s string) (string, error) {
	enc := simplifiedchinese.GBK
	trans := enc.NewDecoder()
	r, _, err := transform.String(trans, s)
	return r, err
}

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// RandString rand string
func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[rand.Intn(len(chars))])
	}
	return string(bs)
}

// IsNumber is number
func IsNumber(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

// ToPercent to percent
func ToPercent(f float64) string {
	if f < 0 {
		f = 0.0
	} else if f > 1.0 {
		f = 1.0
	}
	t := decimal.NewFromFloat(f)
	r := t.Mul(decimal.New(100, 0))
	return strconv.Itoa(int(r.IntPart()))
}

// BoolToByte bool to byte
func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// InArray 判断指定数组中包含指定项
func IsInArray(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// IsEmpty 判断是否为空
func IsEmpty(val interface{}) bool {
	switch val.(type) {
	case string:
		return strings.TrimSpace(val.(string)) == ""
	case byte:
		return val == byte(0)
	case int:
		return val == int(0)
	case int32:
		return val == int32(0)
	case int64:
		return val == int64(0)
	default:
		return false
	}
}
