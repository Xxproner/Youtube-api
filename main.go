package main

// https://yt.lemnoslife.com/videos?part=mostReplayed&id=XiCrniLQGYc

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"os"
	"io"
)

// set go json Unmarshal function
type ProcessJsonStruct struct {
	king string
	etag string
	items []interface{}
}

type VideoResourceStruct struct {
	king string
	etag string
	id   string
	mostReplayed interface{}
}

type ReplayedMarker struct {
	startMillis ulong
	intensityScoreNormalized float32
}

func GetMostReplayedSegment([]ReplayedMarker)

func upper_bound_Impl(nums []int, value int, startIdx, endInd uint32, comp func(int, int) bool) uint32

func upper_bound(nums []int, value int, comp func(int, int) bool) uint32

func main() {
	const youtubeDataApi3Url string = "https://yt.lemnoslife.com/videos?part=mostReplayed&id="

	// get video id
	var targetVideoId string = "XiCrniLQGYc"

	var targetYoutubeDataApi3Url string = youtubeDataApi3Url + targetVideoId

	resp, getRequestErr := http.Get(targetYoutubeDataApi3Url)

	if getRequestErr != nil {
		panic("Error: Get request!")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Error: Response code is not 200!")
	}

	var raw_json_map interface{}
	var json_map map[string]interface{}

	// ProcessJsonStruct mm;

	json_processing_err := json.Unmarshal(Reader(resp.Body), &raw_json_map)
	json_map = raw_json_map.(map[string]interface{})

	// []interface{} === []ReplayedMarker
	replayedMarkersList := json_map["items"][0]["mostReplayed"]["markers"].([]interface{}) // <--- array inteface{}

	point1, point2 := GetMostReplayedSegment(replayedMarkersList)



	// io.Copy(os.Stdout, resp.Body)
}

func GetMostReplayedSegment([]ReplayedMarker, duration ulong) (point1 *ReplayedMarker, point2 *ReplayedMarker, err error) {
	if len(ReplayedMarker) == 0 {
		return nil, nil, errors.New("Empty list")
	}

	var mostReplayedValue float = 0
	var mostReplayedIdx uint = 0

	for i = 0, i < len(ReplayedMarker), ++i {
		if ReplayedMarker[i].intensityScoreNormalized > mostReplayedValue{
			mostReplayedValue = ReplayedMarker[i].
			mostReplayedIdx = i
		}
	}

	if mostReplayed[mostReplayedIdx] + duration / 2 > mostReplayed[len(mostReplayed) - 1] ||
		 mostReplayed[mostReplayedIdx] - duration / 2 < 0 {

		 return nil, nil, errors.New("Too much duration")
	}


	var point1Value ulong = mostReplayed[mostReplayedIdx] - duration / 2;
	var point2Value ulong = mostReplayed[mostReplayedIdx] + duration / 2;

	// binary search for point1 and point2
}

func upper_bound(nums []int, value int, comp func(int, int) bool) uint32 {
	return upper_bound_Impl(nums, value, uint32(0), uint32(len(nums)), comp)
}

func upper_bound_Impl(nums []int, value int, startIdx, endInd uint32, comp func(int, int) bool) uint32 {
	var step uint32 = 0
	var count uint32 = uint32(len(nums))
	var it uint32 = 0

	for count > 0 {
		it = startIdx
		step = count / 2
		it += step

		if !comp(value, nums[it]) {
			it = it + 1
			startIdx = it
			count -= step + 1
		} else {
			count = step
		}
	}

	return it
}
