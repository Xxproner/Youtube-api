package main

// https://yt.lemnoslife.com/videos?part=mostReplayed&id=XiCrniLQGYc

import (
	"fmt"
	_ "net/http"
	_ "encoding/json"
	"strings"
	"os/exec"
	_ "io"
	"runtime"
	"errors"
	"bytes"
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


// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
// openURL opens the specified URL in the default browser of the user.
func OpenBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		// Check if running under WSL
		if isWSL() {
			// Use 'cmd.exe /c start' to open the URL in the default Windows browser
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			// Use xdg-open on native Linux environments
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
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

		if !comp(&markersList[midIdx], target) && !comp(target, &markersList[midIdx]) {
			return midIdx
		}

		if comp(target, &markersList[midIdx]) {
			endIdx = midIdx - 1
		} else {
			startIdx = midIdx + 1
		}
	}

	// find minimum value

	whatIdx := clamp(&markersList[startIdx - 1], target, &markersList[startIdx])

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
			mostReplayedValue = markersList[i].intensityScoreNormalized
			mostReplayedIdx   = i
		}
	}

	// check out-of-range 
	if markersList[mostReplayedIdx].startMillis + duration / 2 > markersList[len(markersList) - 1].startMillis ||
		markersList[mostReplayedIdx].startMillis - duration / 2 < 0 {
			return -1, -1, errors.New("Too long duration")
	}

	var point1Value uint64 = markersList[mostReplayedIdx].startMillis - duration / 2;
	var point2Value uint64 = markersList[mostReplayedIdx].startMillis + duration / 2;

	var target1 ReplayedMarker = ReplayedMarker{point1Value, 0.}
	var target2 ReplayedMarker = ReplayedMarker{point2Value, 0.}

	ReplayedMarkerComparer := func(lhs *ReplayedMarker, rhs *ReplayedMarker) bool {
		return (lhs.startMillis) < (rhs.startMillis)
	}

	// mid guaranteed is between lhs and rhs
	ReplayedMarkerClamp := func(lhs *ReplayedMarker, mid *ReplayedMarker, rhs *ReplayedMarker) int {
		if (rhs.startMillis - mid.startMillis) < (mid.startMillis - lhs.startMillis) {
			return 1;
		}

		return -1;
	}

	var idx1 int = BinarySearchOrNearest(markersList, &target1, 0                  , mostReplayedIdx , ReplayedMarkerComparer, ReplayedMarkerClamp);

	var idx2 int = BinarySearchOrNearest(markersList, &target2, mostReplayedIdx + 1, len(markersList), ReplayedMarkerComparer, ReplayedMarkerClamp);

	return idx1, idx2, nil;
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


type GoogleClientData struct {
	type Web struct {
		Client_id string
		Project_id string
		Auth_uri string 
		Token_uri string
		Auth_provider string
		Client_secret string
		Redirect_uris []string
	}
}

/**
 * TODO: process errors 
 * */
func GoogleAuth2_0ScopeYoutubeDataApi_v3(cliend_id string, redirectUrl string) string {

	// https://accounts.google.com/o/oauth2/v2/auth?
	var buffer *Buffer = NewBufferString("https://accounts.google.com/o/oauth2/v2/auth?")
	
	// nessesary google api query params
	buffer.WriteString("client_id="); buffer.WriteString(cliend_id)
	buffer.WriteString("&redirect_uri="); buffer.WriteString(redirectUrl)
	buffer.WriteString("&response_type="); buffer.WriteString("code")
	buffer.WriteString("&scope="); buffer.WriteString("https://www.googleapis.com/auth/youtube")

	var request_uri = buffer.String()

	return request_uri
}





type FileReader struct {
	file *os.File
	readedData []bytes
}

func (fileReader FileReader) Read(buffer []bytes) (int, error){
	var n int64;
	var read_error error;
	n, read_error = fileReader.file.ReadFrom(buffer);
	return int(n), read_error
}

func (fileReader FileReader) OpenFile(string) error {
	var some_error error
	fileReader.file, some_error = os.Open(string)
	return some_error
}

func (fileReader FileReader) CloseFile() error {
	return fileReader.file.Close()
}

type ParserGoogleAuth2_0Token struct {
	buffer bytes.Buffer
	ServeHTTP(ResponseWriter, *Request)
}

func (parser ParserGoogleAuth2_0Token) (repWriter ResponseWriter, req *Request) {
	
}

func main() {

	var pathToClientSecret string = "client_secret.json"
	var googleClientDataFile FileReader;
	PanicIfError(google_client_data.Open(pathToClientSecret))
	defer googleClientDataFile.CloseFile()

	var decoderGoogleClientData *json.Decoder = json.NewDecoder(googleClientDataFile)
	var googleClientData GoogleClientData;
	PanicIfError(decoderGoogleClientData.Decode(&googleClientData))

	// https proto
	var requestUriAuth2_0 string = GoogleAuth2_0ScopeYoutubeDataApi_v3(googleClietnData.Web.Client_id, googleClientData.Web.Redirect_uris[0])

	var googlTokenAauth_2_0 ParserGoogleAuth2_0Token;
	// PanicIfError(http.ListenAndServeTLS(":8000", "cert.txt", "key.txt", googleTokenAuth_2_0)) // create cert and private key
	PanicIfError(http.ListenAndServer(":8000", googleTokenAuth_2_0))

	if error := OpenBrowser(requestUriAuth2_0); error != nil {
		panic(error)
	}
	// wait answer from console google

	// const youtubeDataApi3Url string = "https://yt.lemnoslife.com/videos?part=mostReplayed&id="

	// var targetVideoId string = GetYoutubeVideoId()

	// var targetYoutubeDataApi3Url string = youtubeDataApi3Url + targetVideoId

	// resp, getRequestErr := http.Get(targetYoutubeDataApi3Url)
	// PanicIfError(getRequestErr)

	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	PanicIfError(errors.New("Response status code not OK"))
	// }

	// rawJsonData, readJsonDataError := io.ReadAll(resp.Body) 
	// PanicIfError(readJsonDataError);

	// var rawJsonMap interface{}

	// jsonProcessingErr := json.Unmarshal(rawJsonData, &rawJsonMap)
	// PanicIfError(jsonProcessingErr)

	// jsonMap := rawJsonMap.(map[string]interface{})

	// // type of this variable is []interface{} or implicilty []ReplayedMarker
	// rawReplayedMarkersList := jsonMap["items"].([]interface{})[0].(map[string]interface{})["mostReplayed"].(map[string]interface{})["markers"].([]interface{}) // <---- array of markers
	
	// // _ = rawReplayedMarkersList

	// var markersList []ReplayedMarker = make([]ReplayedMarker, len(rawReplayedMarkersList))

	// for i := 0; i < len(rawReplayedMarkersList); i++ {
		
	// 	// conversion ReplayedMarker from interface{}
	// 	rawReplayedMarker := rawReplayedMarkersList[i].(map[string]interface{})
		
	// 	markersList[i] = ReplayedMarker{uint64(rawReplayedMarker["startMillis"].(float64)), rawReplayedMarker["intensityScoreNormalized"].(float64)}
	// } // get list of replayed markers

	// point1, point2, getMostReplayedSegmentErr := GetMostReplayedSegment(replayedMarkersList)
	// PanicIfError(getMostReplayedSegmentErr);

	// io.Copy(os.Stdout, resp.Body)


}
