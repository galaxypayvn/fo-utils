package qr

import (
	"encoding/json"
	"gitlab.com/goxp/cloud0/logger"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type Bank struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Code              string `json:"code"`
	Bin               string `json:"bin"`
	ShortName         string `json:"shortName"`
	Logo              string `json:"logo"`
	TransferSupported int    `json:"transferSupported"`
	LookupSupported   int    `json:"lookupSupported"`
	ShortName0        string `json:"short_name"`
	Support           int    `json:"support"`
	IsTransfer        int    `json:"isTransfer"`
	SwiftCode         string `json:"swift_code"`
}

type DataInquirySample struct {
	FullName      string `json:"fullName"`
	PhoneNumber   string `json:"phoneNumber"`
	PersonalId    string `json:"personalId"`
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"bankCode"`
}

var cacheBranch = make(map[string]Bank)
var m sync.Mutex

func init() {
	cacheBranch = LoadBank()
}

func getBinBank(brCode string) string {
	m.Lock()
	defer m.Unlock()
	if val, ok := cacheBranch[brCode]; ok {
		return val.Bin
	}
	return ""
}
func LoadBank() map[string]Bank {
	logg := logger.Tag("QRLib.LoadBank")
	logg.Logger.SetReportCaller(true)

	mydir, err := os.Getwd()
	if err != nil {
		logg.WithError(err).Error("Getwd.err")
		panic(err)
	}
	plan, err := ioutil.ReadFile(path.Join(mydir, "/vietQR_conf/bank.json"))
	if err != nil {
		logg.WithError(err).Error("ReadBankJson.err")
		panic(err)
	}
	var data []Bank
	json.Unmarshal(plan, &data)
	var mapBank = make(map[string]Bank, 0)
	m.Lock()
	defer m.Unlock()
	for _, item := range data {
		mapBank[item.Code] = item
	}
	return mapBank
}

func LoadSampleInquiry() map[string]DataInquirySample {
	logg := logger.Tag("QRLib.LoadSampleInquiry")
	logg.Logger.SetReportCaller(true)

	mydir, err := os.Getwd()
	if err != nil {
		logg.WithError(err).Error("Getwd.err")
		panic(err)
	}
	plan, err := ioutil.ReadFile(path.Join(mydir, "/vietQR_conf/data_example.json"))
	if err != nil {
		logg.WithError(err).Error("LoadSampleInquiry.err")
		panic(err)
	}
	var data []DataInquirySample
	json.Unmarshal(plan, &data)
	var mapBank = make(map[string]DataInquirySample, 0)

	for _, item := range data {
		mapBank[item.BankCode+"_"+item.AccountNumber] = item
	}
	return mapBank
}
