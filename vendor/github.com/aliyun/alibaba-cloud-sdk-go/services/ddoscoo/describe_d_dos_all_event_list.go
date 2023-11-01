package ddoscoo

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDDosAllEventList invokes the ddoscoo.DescribeDDosAllEventList API synchronously
func (client *Client) DescribeDDosAllEventList(request *DescribeDDosAllEventListRequest) (response *DescribeDDosAllEventListResponse, err error) {
	response = CreateDescribeDDosAllEventListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDDosAllEventListWithChan invokes the ddoscoo.DescribeDDosAllEventList API asynchronously
func (client *Client) DescribeDDosAllEventListWithChan(request *DescribeDDosAllEventListRequest) (<-chan *DescribeDDosAllEventListResponse, <-chan error) {
	responseChan := make(chan *DescribeDDosAllEventListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDDosAllEventList(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDDosAllEventListWithCallback invokes the ddoscoo.DescribeDDosAllEventList API asynchronously
func (client *Client) DescribeDDosAllEventListWithCallback(request *DescribeDDosAllEventListRequest, callback func(response *DescribeDDosAllEventListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDDosAllEventListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDDosAllEventList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDDosAllEventListRequest is the request struct for api DescribeDDosAllEventList
type DescribeDDosAllEventListRequest struct {
	*requests.RpcRequest
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	EventType  string           `position:"Query" name:"EventType"`
}

// DescribeDDosAllEventListResponse is the response struct for api DescribeDDosAllEventList
type DescribeDDosAllEventListResponse struct {
	*responses.BaseResponse
	Total        int64         `json:"Total" xml:"Total"`
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	AttackEvents []AttackEvent `json:"AttackEvents" xml:"AttackEvents"`
}

// CreateDescribeDDosAllEventListRequest creates a request to invoke DescribeDDosAllEventList API
func CreateDescribeDDosAllEventListRequest() (request *DescribeDDosAllEventListRequest) {
	request = &DescribeDDosAllEventListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeDDosAllEventList", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDDosAllEventListResponse creates a response to parse from DescribeDDosAllEventList response
func CreateDescribeDDosAllEventListResponse() (response *DescribeDDosAllEventListResponse) {
	response = &DescribeDDosAllEventListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
