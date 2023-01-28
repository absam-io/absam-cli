package api

import (
	"encoding/json"
	"net/url"

	"github.com/absam-io/absam-cli/models"
	"github.com/absam-io/absam-cli/utils"
)

func GetCloudDatabase() models.Servers {
	data, err := utils.Get("https://api.absam.io/v1/cloud-database/listall")
	if err != nil {
		utils.Die(err)
	}

	var servers models.Servers
	json.Unmarshal(data, &servers)

	return servers
}

func GetCloudDatabaseInfo(id string) models.SingleServer {
	url := "https://api.absam.io/v1/cloud-database/info?id=" + id

	data, err := utils.Get(url)
	if err != nil {
		utils.Die(err)
	}

	var server models.SingleServer
	json.Unmarshal(data, &server)

	return server
}

func GetCloudDatabaseStatus(id string) models.Status {
	url := "https://api.absam.io/v1/cloud-database/status?id=" + id

	data, err := utils.Get(url)
	if err != nil {
		utils.Die(err)
	}

	var status models.Status
	json.Unmarshal(data, &status)

	return status
}

func ManageCloudDatabase(action, id string) models.Result {
	reqUrl := "https://api.absam.io/v1/cloud-database/" + action

	form := url.Values{}
	form.Set("id", id)

	data, err := utils.Post(reqUrl, form)
	if err != nil {
		utils.Die(err)
	}

	var result models.Result
	json.Unmarshal(data, &result)

	return result
}
