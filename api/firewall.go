package api

import (
	"encoding/json"
	"net/url"

	"github.com/absam-io/absam-cli/models"
	"github.com/absam-io/absam-cli/utils"
)

const (
	ADD_RULE    = "addrule"
	EDIT_RULE   = "editrule"
	DELETE_RULE = "delrule"
)

func GetFirewallStatus(id, product string) models.FirewallStatus {
	url := "https://api.absam.io/v1/" + product + "/firewall?id=" + id + "&action=status"

	data, err := utils.Get(url)
	if err != nil {
		utils.Die(err)
	}

	var status models.FirewallStatus
	json.Unmarshal(data, &status)

	return status
}

func GetFirewallRules(id, product string) models.FirewallRules {
	url := "https://api.absam.io/v1/" + product + "/firewall?id=" + id + "&action=list"

	data, err := utils.Get(url)
	if err != nil {
		utils.Die(err)
	}

	var rules models.FirewallRules
	json.Unmarshal(data, &rules)

	return rules
}

func ChangeFirewallStatus(id, action, product string) models.ChangeFirewallSuccess {
	reqUrl := "https://api.absam.io/v1/" + product + "/firewall"

	form := url.Values{}
	form.Set("id", id)
	form.Set("action", "change")
	form.Set("set", action)

	data, err := utils.Post(reqUrl, form)
	if err != nil {
		utils.Die(err)
	}

	var result models.ChangeFirewallSuccess
	json.Unmarshal(data, &result)

	return result
}

func AddRule(id, reqType, port, ip, proto, comment, expires, product string) models.ChangeFirewallSuccess {
	reqUrl := "https://api.absam.io/v1/" + product + "/firewall"

	form := url.Values{}
	form.Set("id", id)
	form.Set("action", ADD_RULE)
	form.Set("type", reqType)
	form.Set("port", port)
	form.Set("ip", ip)
	form.Set("proto", proto)
	form.Set("comment", comment)
	form.Set("expires", expires)

	data, err := utils.Post(reqUrl, form)
	if err != nil {
		utils.Die(err)
	}

	var result models.ChangeFirewallSuccess
	json.Unmarshal(data, &result)

	return result
}

func EditRule(id, reqType, port, ip, proto, comment, expires, position, product string) models.ChangeFirewallSuccess {
	reqUrl := "https://api.absam.io/v1/" + product + "/firewall"

	form := url.Values{}
	form.Set("id", id)
	form.Set("action", EDIT_RULE)
	form.Set("type", reqType)
	form.Set("port", port)
	form.Set("ip", ip)
	form.Set("proto", proto)
	form.Set("comment", comment)
	form.Set("expires", expires)
	form.Set("pos", position)

	data, err := utils.Post(reqUrl, form)
	if err != nil {
		utils.Die(err)
	}

	var result models.ChangeFirewallSuccess
	json.Unmarshal(data, &result)

	return result
}

func RemoveRule(id, pos, product string) models.ChangeFirewallSuccess {
	reqUrl := "https://api.absam.io/v1/" + product + "/firewall"

	form := url.Values{}
	form.Set("id", id)
	form.Set("action", DELETE_RULE)
	form.Set("pos", pos)

	data, err := utils.Post(reqUrl, form)
	if err != nil {
		utils.Die(err)
	}

	var result models.ChangeFirewallSuccess
	json.Unmarshal(data, &result)

	return result
}
