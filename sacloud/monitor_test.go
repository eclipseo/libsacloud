// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCPUResourceMonitorJSON = `
    {
        "2016-05-07T12:15:00+09:00": {
            "CPU-TIME": 0.5
        },
        "2016-05-07T12:20:00+09:00": {
            "CPU-TIME": 0.1
        },
        "2016-05-07T12:25:00+09:00": {
            "CPU-TIME": 0
        },
        "2016-05-07T12:30:00+09:00": {
            "CPU-TIME": null
        }
    }
`
var testNICResourceMonitorJSON = `
    {
        "2016-05-07T12:15:00+09:00": {
            "Receive": 305.02666667,
            "Send": 17.39
        },
        "2016-05-07T12:20:00+09:00": {
            "Receive": 283.50333333,
            "Send": 14.336666667
        },
        "2016-05-07T12:25:00+09:00": {
            "Receive": 304.03,
            "Send": 14.643333333
        },
        "2016-05-07T12:30:00+09:00": {
            "Receive": null,
            "Send": null
        }
    }
`

var tesDiskResourceMonitorJSON = `
    {
        "2016-05-07T12:15:00+09:00": {
            "Read": 0,
            "Write": 286.72
        },
        "2016-05-07T12:20:00+09:00": {
            "Read": 0,
            "Write": 204.8
        },
        "2016-05-07T12:25:00+09:00": {
            "Read": 0,
            "Write": 81.92
        },
        "2016-05-07T12:30:00+09:00": {
            "Read": null,
            "Write": null
        }
    }
`

var testInternetMonitorJSON = `
{
	"2016-09-22T11:00:00+09:00": {
	    "In": 0,
	    "Out": 0
	},
	"2016-09-22T11:05:00+09:00": {
	    "In": 0,
	    "Out": 0
	},
	"2016-09-22T11:10:00+09:00": {
	    "In": 0,
	    "Out": 0
	},
	"2016-09-22T11:15:00+09:00": {
	    "In": null,
	    "Out": null
	}
}
`

var testDatabaseMonitorJSON = `
    {
	"2016-09-22T08:20:00+09:00": {
	    "Total-Memory-Size": null,
	    "Used-Memory-Size": null,
	    "Total-Disk1-Size": null,
	    "Used-Disk1-Size": null,
	    "Total-Disk2-Size": null,
	    "Used-Disk2-Size": null
	},
	"2016-09-22T08:25:00+09:00": {
	    "Total-Memory-Size": null,
	    "Used-Memory-Size": null,
	    "Total-Disk1-Size": null,
	    "Used-Disk1-Size": null,
	    "Total-Disk2-Size": null,
	    "Used-Disk2-Size": null
	},
	"2016-09-22T08:30:00+09:00": {
	    "Total-Memory-Size": 1884224,
	    "Used-Memory-Size": 239448,
	    "Total-Disk1-Size": 18544316,
	    "Used-Disk1-Size": 15761644,
	    "Total-Disk2-Size": 100982868,
	    "Used-Disk2-Size": 1880016
	},
	"2016-09-22T08:35:00+09:00": {
	    "Total-Memory-Size": 1884224,
	    "Used-Memory-Size": 239340,
	    "Total-Disk1-Size": 18544316,
	    "Used-Disk1-Size": 15761596,
	    "Total-Disk2-Size": 100982868,
	    "Used-Disk2-Size": 1880064
	}
    }
`

var testNFSMonitorJSON = `
	{
	"2017-09-07T12:20:00+09:00": {
      "Free-Disk-Size": null
    },
    "2017-09-07T12:25:00+09:00": {
      "Free-Disk-Size": 97776352
    },
    "2017-09-07T12:30:00+09:00": {
      "Free-Disk-Size": 97776352
    },
    "2017-09-07T12:35:00+09:00": {
      "Free-Disk-Size": 97776352
    }
    }
`

var testEmptyMonitorJSON = `
{
	"Data": {	
		"2017-09-07T12:20:00+09:00": [],
    	"2017-09-07T12:25:00+09:00": []
    },
    "is_ok": true

}
`

var testEmptyMonitorDataJSON = `
{
	"Data": [],
    "is_ok": true
}
`
var testResourceMonitorResponseJSON = `
{
    "Data": ` + testCPUResourceMonitorJSON + `,
    "is_ok" : true
}
`

func TestMarshalResourceMonitorJSON(t *testing.T) {
	var m MonitorValues
	err := json.Unmarshal([]byte(testCPUResourceMonitorJSON), &m)

	assert.NoError(t, err)
	assert.NotEmpty(t, m)
}

func TestMarshalCPUResourceMonitorJSON(t *testing.T) {
	var m ResourceMonitorResponse
	err := json.Unmarshal([]byte(testEmptyMonitorJSON), &m)

	assert.NoError(t, err)
	assert.NotEmpty(t, m)
	assert.NotEmpty(t, m.Data)

}

func TestMarshalEmptyMonitorJSON(t *testing.T) {
	var m ResourceMonitorResponse
	err := json.Unmarshal([]byte(testResourceMonitorResponseJSON), &m)

	assert.NoError(t, err)
	assert.NotEmpty(t, m)
	assert.NotEmpty(t, m.Data)
}

func TestMarshalEmptyMonitorDataJSON(t *testing.T) {
	var m ResourceMonitorResponse
	err := json.Unmarshal([]byte(testEmptyMonitorDataJSON), &m)

	assert.NoError(t, err)
	assert.NotEmpty(t, m)
	assert.Empty(t, m.Data)
}

func TestFlattenMonitorValues(t *testing.T) {
	var monitor MonitorValues
	json.Unmarshal([]byte(testCPUResourceMonitorJSON), &monitor)

	res, err := monitor.FlattenCPUTimeValue()
	assert.NoError(t, err)
	assert.Len(t, res, 3)

	// 順不同なため以下テストは通らない
	//t.Logf("values : %#v", res)
	//assert.Equal(t, res[0].Value, 0.5)
	//assert.Equal(t, res[1].Value, 0.1)
	//assert.Equal(t, res[2].Value, 0)
}

func TestResourceMonitorCalc(t *testing.T) {
	var monitor MonitorValues
	json.Unmarshal([]byte(testCPUResourceMonitorJSON), &monitor)

	var sum = 0.6
	var count float64 = 3
	var max = 0.5
	var min float64

	calcResult := monitor.Calc()
	assert.NotNil(t, calcResult)
	assert.NotNil(t, calcResult.CPU)
	assert.Equal(t, calcResult.CPU.Avg, sum/count)
	assert.Equal(t, calcResult.CPU.Count, count)
	assert.Equal(t, calcResult.CPU.Max, max)
	assert.Equal(t, calcResult.CPU.Min, min)

}
