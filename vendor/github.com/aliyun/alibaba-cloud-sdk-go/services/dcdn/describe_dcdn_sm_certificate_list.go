package dcdn

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

// DescribeDcdnSMCertificateList invokes the dcdn.DescribeDcdnSMCertificateList API synchronously
func (client *Client) DescribeDcdnSMCertificateList(request *DescribeDcdnSMCertificateListRequest) (response *DescribeDcdnSMCertificateListResponse, err error) {
	response = CreateDescribeDcdnSMCertificateListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnSMCertificateListWithChan invokes the dcdn.DescribeDcdnSMCertificateList API asynchronously
func (client *Client) DescribeDcdnSMCertificateListWithChan(request *DescribeDcdnSMCertificateListRequest) (<-chan *DescribeDcdnSMCertificateListResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnSMCertificateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnSMCertificateList(request)
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

// DescribeDcdnSMCertificateListWithCallback invokes the dcdn.DescribeDcdnSMCertificateList API asynchronously
func (client *Client) DescribeDcdnSMCertificateListWithCallback(request *DescribeDcdnSMCertificateListRequest, callback func(response *DescribeDcdnSMCertificateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnSMCertificateListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnSMCertificateList(request)
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

// DescribeDcdnSMCertificateListRequest is the request struct for api DescribeDcdnSMCertificateList
type DescribeDcdnSMCertificateListRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnSMCertificateListResponse is the response struct for api DescribeDcdnSMCertificateList
type DescribeDcdnSMCertificateListResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	CertificateListModel CertificateListModel `json:"CertificateListModel" xml:"CertificateListModel"`
}

// CreateDescribeDcdnSMCertificateListRequest creates a request to invoke DescribeDcdnSMCertificateList API
func CreateDescribeDcdnSMCertificateListRequest() (request *DescribeDcdnSMCertificateListRequest) {
	request = &DescribeDcdnSMCertificateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnSMCertificateList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnSMCertificateListResponse creates a response to parse from DescribeDcdnSMCertificateList response
func CreateDescribeDcdnSMCertificateListResponse() (response *DescribeDcdnSMCertificateListResponse) {
	response = &DescribeDcdnSMCertificateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}