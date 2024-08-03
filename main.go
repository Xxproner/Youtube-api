package main

// https://yt.lemnoslife.com/videos?part=mostReplayed&id=XiCrniLQGYc

import (
	"fmt"
	"net/http"
	"encoding/json"
	// "strings"
	_ "os"
	"io"
	"errors"
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
	startMillis uint64 // change to StartMillis
	intensityScoreNormalized float64 // change to IntensityScoreNormalized
}

/**
 * comp must return true if lhs is less than rhs 
 * clamp similar to c++ clamp but returns negative if left and else if right positive
 * */
func BinarySearchOrNearest(markersList []ReplayedMarker, target *ReplayedMarker, startIdx, endIdx int, comp (func(*ReplayedMarker, *ReplayedMarker) bool), 
	clamp (func (*ReplayedMarker, *ReplayedMarker, *ReplayedMarker)int)) int {
	var midIdx int 
	for startIdx < endIdx {
		midIdx = (startIdx + endIdx) / 2

		if !comp(&markersList[midIdx], &target) && !comp(&target, &markersList[midIdx]) {
			return midIdx
		}

		if comp(&target, &markersList[midIdx]) {
			endIdx = midIdx - 1
		} else {
			startIdx = midIdx + 1
		}
	}

	// find minimum value

	whatIdx := clamp(&markersList[startIdx - 1], &target, &markersList[startIdx])

	if whatIdx < 0 {
		return startIdx - 1
	}

	return startIdx // always index of greater value
}

/**
 * return values are first, second idxes and error if any 
 * */
func GetMostReplayedSegment(markersList []ReplayedMarker, duration uint64) (int, int, error) {
	if len(markersList) == 0 {
		return -1, -1, errors.New("Empty list")
	}

	var mostReplayedValue float64 = 0
	var mostReplayedIdx   int     = 0

	// find max intensity in the markers list
	for i := 0; i < len(markersList); i++ {
		if markersList[i].intensityScoreNormalized > mostReplayedValue {
			mostReplayedValue = markersList[i]
			mostReplayedIdx   = i
		}
	}

	// check out-of-range 
	if markersList[mostReplayedIdx] + duration / 2 > markersList[len(markersList) - 1] ||
		markersList[mostReplayedIdx] - duration / 2 < 0 {
			return -1, -1, errors.New("Too long duration")
	}

	var point1Value uint64 = markersList[mostReplayedIdx] - duration / 2;
	var point2Value uint64 = markersList[mostReplayedIdx] + duration / 2;

	var target1 ReplayedMarker = {point1Value}
	var target2 ReplayedMarker = {point2Value}

	ReplayedMarkerComparer := func(lhs *ReplayedMarker, rhs *ReplayedMarker) bool {
		return lhs->startMillis < rhs->startMillis
	}

	// mid guaranteed is between lhs and rhs
	ReplayedMarkerClamp := func(lhs *ReplayedMarker, mid *ReplayedMarker, rhs *ReplayedMarker) int {
		if (rhs->startMillis - mid->startMillis < mid->startMillis - lhs->startMillis) {
			return 1;
		}

		return -1;
	}

	var idx1 int = BinarySearchOrNearest(markersList, &target1, 0                  , mostReplayedIdx , ReplayedMarkerComparer, ReplayedMarkerClamp);

	var idx2 int = BinarySearchOrNearest(markersList, &target2, mostReplayedIdx + 1, len(markersList), ReplayedMarkerComparer, ReplayedMarkerClamp);

	_ = idx1; 
	_ = idx2;
}

func GetYoutubeVideoId() string{
	return string("XiCrniLQGYc")
}

func PanicIfError(err error) {
	if (err != nil) {
		fmt.Println("Error:", err)
		panic("")			
	}
}

func main() {

	const youtubeDataApi3Url string = "https://yt.lemnoslife.com/videos?part=mostReplayed&id="

	var targetVideoId string = GetYoutubeVideoId()

	var targetYoutubeDataApi3Url string = youtubeDataApi3Url + targetVideoId

	resp, getRequestErr := http.Get(targetYoutubeDataApi3Url)
	PanicIfError(getRequestErr)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		PanicIfError(errors.New("Response status code not OK"))
	}

	rawJsonData, readJsonDataError := io.ReadAll(resp.Body) 
	PanicIfError(readJsonDataError);

	var rawJsonMap interface{}

	jsonProcessingErr := json.Unmarshal(rawJsonData, &rawJsonMap)
	PanicIfError(jsonProcessingErr)

	jsonMap := rawJsonMap.(map[string]interface{})

	// type of this variable is []interface{} or implicilty []ReplayedMarker
	rawReplayedMarkersList := jsonMap["items"].([]interface{})[0].(map[string]interface{})["mostReplayed"].(map[string]interface{})["markers"].([]interface{}) // <---- array of markers
	
	// _ = rawReplayedMarkersList

	var markersList []ReplayedMarker = make([]ReplayedMarker, len(rawReplayedMarkersList))

	for i := 0; i < len(rawReplayedMarkersList); i++ {
		
		// conversion ReplayedMarker from interface{}
		rawReplayedMarker := rawReplayedMarkersList[i].(map[string]interface{})
		
		markersList[i] = ReplayedMarker{uint64(rawReplayedMarker["startMillis"].(float64)), rawReplayedMarker["intensityScoreNormalized"].(float64)}
	} // get list of replayed markers

	point1, point2, getMostReplayedSegmentErr := GetMostReplayedSegment(replayedMarkersList)
	PanicIfError(getMostReplayedSegmentErr);

	// io.Copy(os.Stdout, resp.Body)
}
