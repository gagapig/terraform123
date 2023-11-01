package polardb

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

// ItemsItem is a nested struct in polardb response
type ItemsItem struct {
	MemoryClass        string `json:"MemoryClass" xml:"MemoryClass"`
	Region             string `json:"Region" xml:"Region"`
	DBType             string `json:"DBType" xml:"DBType"`
	ResultInfo         string `json:"ResultInfo" xml:"ResultInfo"`
	ClassGroup         string `json:"ClassGroup" xml:"ClassGroup"`
	Psl4MaxIOPS        string `json:"Psl4MaxIOPS" xml:"Psl4MaxIOPS"`
	MaxStorageCapacity string `json:"MaxStorageCapacity" xml:"MaxStorageCapacity"`
	ClassCode          string `json:"ClassCode" xml:"ClassCode"`
	Deadline           string `json:"Deadline" xml:"Deadline"`
	ModifiedTime       string `json:"ModifiedTime" xml:"ModifiedTime"`
	CreatedTime        string `json:"CreatedTime" xml:"CreatedTime"`
	SwitchTime         string `json:"SwitchTime" xml:"SwitchTime"`
	ReferenceExtPrice  string `json:"ReferenceExtPrice" xml:"ReferenceExtPrice"`
	ClassTypeLevel     string `json:"ClassTypeLevel" xml:"ClassTypeLevel"`
	Id                 int    `json:"Id" xml:"Id"`
	Status             int    `json:"Status" xml:"Status"`
	DBVersion          string `json:"DBVersion" xml:"DBVersion"`
	Cpu                string `json:"Cpu" xml:"Cpu"`
	DBClusterId        string `json:"DBClusterId" xml:"DBClusterId"`
	ReferencePrice     string `json:"ReferencePrice" xml:"ReferencePrice"`
	MaxIOPS            string `json:"MaxIOPS" xml:"MaxIOPS"`
	PrepareInterval    string `json:"PrepareInterval" xml:"PrepareInterval"`
	Psl5MaxIOPS        string `json:"Psl5MaxIOPS" xml:"Psl5MaxIOPS"`
	StartTime          string `json:"StartTime" xml:"StartTime"`
	MaxConnections     string `json:"MaxConnections" xml:"MaxConnections"`
	Pl3MaxIOPS         string `json:"Pl3MaxIOPS" xml:"Pl3MaxIOPS"`
	Pl1MaxIOPS         string `json:"Pl1MaxIOPS" xml:"Pl1MaxIOPS"`
	TaskType           string `json:"TaskType" xml:"TaskType"`
	Pl2MaxIOPS         string `json:"Pl2MaxIOPS" xml:"Pl2MaxIOPS"`
}
