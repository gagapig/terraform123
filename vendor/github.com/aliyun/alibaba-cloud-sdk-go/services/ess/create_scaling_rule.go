package ess

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

// CreateScalingRule invokes the ess.CreateScalingRule API synchronously
func (client *Client) CreateScalingRule(request *CreateScalingRuleRequest) (response *CreateScalingRuleResponse, err error) {
	response = CreateCreateScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScalingRuleWithChan invokes the ess.CreateScalingRule API asynchronously
func (client *Client) CreateScalingRuleWithChan(request *CreateScalingRuleRequest) (<-chan *CreateScalingRuleResponse, <-chan error) {
	responseChan := make(chan *CreateScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScalingRule(request)
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

// CreateScalingRuleWithCallback invokes the ess.CreateScalingRule API asynchronously
func (client *Client) CreateScalingRuleWithCallback(request *CreateScalingRuleRequest, callback func(response *CreateScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateScalingRule(request)
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

// CreateScalingRuleRequest is the request struct for api CreateScalingRule
type CreateScalingRuleRequest struct {
	*requests.RpcRequest
	AlarmDimension           *[]CreateScalingRuleAlarmDimension `position:"Query" name:"AlarmDimension"  type:"Repeated"`
	StepAdjustment           *[]CreateScalingRuleStepAdjustment `position:"Query" name:"StepAdjustment"  type:"Repeated"`
	ScalingGroupId           string                             `position:"Query" name:"ScalingGroupId"`
	DisableScaleIn           requests.Boolean                   `position:"Query" name:"DisableScaleIn"`
	InitialMaxSize           requests.Integer                   `position:"Query" name:"InitialMaxSize"`
	ScalingRuleName          string                             `position:"Query" name:"ScalingRuleName"`
	Cooldown                 requests.Integer                   `position:"Query" name:"Cooldown"`
	PredictiveValueBehavior  string                             `position:"Query" name:"PredictiveValueBehavior"`
	ScaleInEvaluationCount   requests.Integer                   `position:"Query" name:"ScaleInEvaluationCount"`
	ScalingRuleType          string                             `position:"Query" name:"ScalingRuleType"`
	MetricName               string                             `position:"Query" name:"MetricName"`
	PredictiveScalingMode    string                             `position:"Query" name:"PredictiveScalingMode"`
	ResourceOwnerAccount     string                             `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue          requests.Integer                   `position:"Query" name:"AdjustmentValue"`
	EstimatedInstanceWarmup  requests.Integer                   `position:"Query" name:"EstimatedInstanceWarmup"`
	OwnerAccount             string                             `position:"Query" name:"OwnerAccount"`
	PredictiveTaskBufferTime requests.Integer                   `position:"Query" name:"PredictiveTaskBufferTime"`
	AdjustmentType           string                             `position:"Query" name:"AdjustmentType"`
	OwnerId                  requests.Integer                   `position:"Query" name:"OwnerId"`
	PredictiveValueBuffer    requests.Integer                   `position:"Query" name:"PredictiveValueBuffer"`
	ScaleOutEvaluationCount  requests.Integer                   `position:"Query" name:"ScaleOutEvaluationCount"`
	MinAdjustmentMagnitude   requests.Integer                   `position:"Query" name:"MinAdjustmentMagnitude"`
	TargetValue              requests.Float                     `position:"Query" name:"TargetValue"`
}

// CreateScalingRuleAlarmDimension is a repeated param struct in CreateScalingRuleRequest
type CreateScalingRuleAlarmDimension struct {
	DimensionValue string `name:"DimensionValue"`
	DimensionKey   string `name:"DimensionKey"`
}

// CreateScalingRuleStepAdjustment is a repeated param struct in CreateScalingRuleRequest
type CreateScalingRuleStepAdjustment struct {
	MetricIntervalUpperBound string `name:"MetricIntervalUpperBound"`
	MetricIntervalLowerBound string `name:"MetricIntervalLowerBound"`
	ScalingAdjustment        string `name:"ScalingAdjustment"`
}

// CreateScalingRuleResponse is the response struct for api CreateScalingRule
type CreateScalingRuleResponse struct {
	*responses.BaseResponse
	ScalingRuleAri string `json:"ScalingRuleAri" xml:"ScalingRuleAri"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ScalingRuleId  string `json:"ScalingRuleId" xml:"ScalingRuleId"`
}

// CreateCreateScalingRuleRequest creates a request to invoke CreateScalingRule API
func CreateCreateScalingRuleRequest() (request *CreateScalingRuleRequest) {
	request = &CreateScalingRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingRule", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateScalingRuleResponse creates a response to parse from CreateScalingRule response
func CreateCreateScalingRuleResponse() (response *CreateScalingRuleResponse) {
	response = &CreateScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
