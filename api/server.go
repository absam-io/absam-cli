package api

import (
	"encoding/json"
	"net/url"

	"github.com/absam-io/absam-cli/models"
	"github.com/absam-io/absam-cli/utils"
)

func GetServers() models.Servers {
	data, err := utils.Get("http://api.absam.io/v1/cloud-server/listall")
	if err != nil {
		utils.Die(err)
	}

	var servers models.Servers
	json.Unmarshal(data, &servers)

	return servers
}

func GetServerInfo(id string) models.SingleServer {
	url := "http://api.absam.io/v1/cloud-server/info?id=" + id

	data, err := utils.Get(url)
	if err != nil {
		utils.Die(err)
	}

	var server models.SingleServer
	json.Unmarshal(data, &server)

	return server
}

func GetServerStatus(id string) models.Status {
	url := "http://api.absam.io/v1/cloud-server/status?id=" + id

	data, err := utils.Get(url)
	if err != nil {
		utils.Die(err)
	}

	var status models.Status
	json.Unmarshal(data, &status)

	return status
}

func ManageServer(action, id string) models.Result {
	reqUrl := "http://api.absam.io/v1/cloud-server/" + action

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
