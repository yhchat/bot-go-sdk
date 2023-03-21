/*
 * @Date: 2023-01-06 10:45:46
 * @LastEditTime: 2023-03-21 15:46:30
 *
 * Copyright (c) 2023 by 北京九万智达科技有限公司, All Rights Reserved.
 */
package openapi

const (
	API_BASE_URL = "https://chat-go.jwzhd.com/open-apis/v1"
)

type OpenApi struct {
	Token string `json:"token"`
}

func NewOpenApi(token string) *OpenApi {
	return &OpenApi{
		Token: token,
	}
}
