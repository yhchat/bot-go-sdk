/*
 * @Date: 2022-07-24 22:57:40
 * @LastEditTime: 2023-01-06 10:07:54
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package utils

import (
	"github.com/go-resty/resty/v2"
)

//发送POST请求
//重试3次
func HttpPost(url string, data []byte) (resp *resty.Response, err error) {
	client := resty.New()
	resp, err = client.SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(url)
	return resp, err
}
