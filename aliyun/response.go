package aliyun

import "fmt"

type Response struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
}

//ErrorResponse 错误信息结构体
type ErrorResponse struct {
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Recommend  string `json:"Recommend" xml:"Recommend"`
	HostId     string `json:"HostId" xml:"HostId"`
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	StatusCode int    `json:"StatusCode" xml:"StatusCode"` //Status Code of HTTP Response
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Aliyun API Error:\n RequestId: %s\n Status Code: %d\n Code: %s\n Message: %s\n", e.RequestId, e.StatusCode, e.Code, e.Message)
}
