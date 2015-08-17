package socialista

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
)

type Platform struct {
	name          string
	statsURL      string
	parseResponse func(*http.Response) Stat
}

type Stat struct {
	platform *Platform
	data     map[string]string
}

func (platform Platform) getStatsFor(url string, stats chan <- Stat, errors chan *error) {
	response, error := http.Get(fmt.Sprintf(platform.statsURL, url))
	if error != nil {
		errors <- &error
	}

	stat := platform.parseResponse(response)
	stat.platform = &platform
	stats <- stat
}

func GetStatsForPlatform(url string, selectedPlatform string) {
	platforms := [...]Platform{
		Platform{"twitter", "http://urls.api.twitter.com/1/urls/count.json?url=%s", func(response *http.Response) (stat Stat) {
			var holder map[string]interface{}

			body, _ := ioutil.ReadAll(response.Body)
			if err := json.Unmarshal(body, &holder); err != nil { panic(err) }

			stat.data = map[string]string{"share_count": fmt.Sprintf("%.f", holder["count"])}
			return
		}},
		Platform{"linkedin", "http://www.linkedin.com/countserv/count/share?url=%s", func(response *http.Response) (stat Stat) {
			body, _ := ioutil.ReadAll(response.Body)
			json_body := strings.TrimRight(strings.TrimLeft(string(body), "IN.Tags.Share.handleCount("), ");", )

			var holder map[string]interface{}
			if err := json.Unmarshal([]byte(json_body), &holder); err != nil { panic(err) }

			stat.data = map[string]string{"share_count": fmt.Sprintf("%.f", holder["count"])}
			return stat
		}},
		Platform{"facebook", "https://api.facebook.com/method/links.getStats?format=json&urls=%s", func(response *http.Response) (Stat) {
			body, _ := ioutil.ReadAll(response.Body)
			var jsonBlob []interface{}
			if err := json.Unmarshal(body, &jsonBlob); err != nil { panic(err) }

			jsonObject := jsonBlob[0].(map[string]interface{})
			return Stat{data: map[string]string{
				"like_count": fmt.Sprintf("%.f", jsonObject["like_count"].(float64)),
				"comment_count": fmt.Sprintf("%.f", jsonObject["comment_count"].(float64)),
				"total_count": fmt.Sprintf("%.f", jsonObject["total_count"].(float64)),
				"share_count": fmt.Sprintf("%.f", jsonObject["share_count"].(float64)),
			}}
		}},
		Platform{"pintarest", "http://api.pinterest.com/v1/urls/count.json?callback=&url=%s", func(response *http.Response) (stat Stat) {
			body, _ := ioutil.ReadAll(response.Body)

			fmt.Println("pintarest", body)

			stat.data = nil;
			return;
		}},
	}

	errors := make(chan *error)
	var stats chan Stat = make(chan Stat)
	var platformsCount int = 0

	for _, platform := range platforms {
		if selectedPlatform == "" {
			go platform.getStatsFor(url, stats, errors)
			platformsCount++
		}   else {
			var selectedPlatforms = strings.Split(selectedPlatform, ",")
			var in = false
			for _, iPlatform := range selectedPlatforms {
				if strings.EqualFold(iPlatform, platform.name) {
					in = true;
				}
			}

			if in {
				go platform.getStatsFor(url, stats, errors)
				platformsCount++;
			}
		}

	}

	for {
		select {
		case stat := <-stats:
			fmt.Println(stat.platform.name)
			for k, v := range (stat.data) { fmt.Println("\t", k, " =", v) }

			platformsCount--
		case error := <-errors:
			fmt.Println("ERROR ~~> ")
			log.Fatal(error)
		default:
			if platformsCount <= 0 { return }
		}
	}
}

func GetStats(url string) {
	GetStatsForPlatform(url, "")
}
