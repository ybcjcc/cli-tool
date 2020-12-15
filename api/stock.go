package api

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Stock(in string) error  {
	tpl := "https://hq.sinajs.cn/list=s_%s"
	r, err := http.Get(fmt.Sprintf(tpl, in))
	if err != nil {
		return err
	}
	txt, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()
	if err != nil {
		return err
	}

	str := fmt.Sprintf("%s", txt)
	err = strToUtf8(&str)
	if err != nil {
		return err
	}

	parts := strings.Split(str, "=")
	if len(parts) != 2 {
		return errors.New("response error")
	}

	s := parts[1]
	s = strings.Replace(s, `"`, "", 2)
	s = strings.Replace(s, `;`, "", 2)
	s = strings.Replace(s, "\n", "", 2)
	s2 := strings.Split(s, ",")

	fmt.Println("名字: " + s2[0])
	fmt.Println("代码: " + in)
	fmt.Println("现价: " + s2[1])
	fmt.Println("涨跌: " + s2[2])
	fmt.Println("幅度: " + s2[3])
	fmt.Println("时间: " + time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func gbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func strToUtf8(str *string) error {
	b, err := gbkToUtf8([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}