package icity_sdk

import (
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

type Device struct {
	App      string
	Hardware string
	LastUsed time.Time

	// RemoveLink is the link to remove device.
	RemoveLink string
}

// GetDevices gets devices we have login iCity.
func (user *User) GetDevices() []Device {
	doc, err := user.getWithDoc(path.Devices)
	if err != nil {
		return nil
	}

	var devices []Device
	doc.Find(selector.Devices).Each(func(i int, s *goquery.Selection) {
		devices = append(devices, parseDevice(s))
	})

	return devices
}

// RemoveDevice removes one device.
func (user *User) RemoveDevice(device Device) bool {
	urlPath := device.RemoveLink

	postData := url.Values{}
	postData.Set(data.MethodKEY, "put")
	resp, err := user.postForm(urlPath, postData)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}
