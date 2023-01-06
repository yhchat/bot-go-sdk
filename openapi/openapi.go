package openapi

type OpenApi struct {
	Token string `json:"token"`
}

func NewOpenApi(token string) *OpenApi {
	return &OpenApi{
		Token: token,
	}
}
