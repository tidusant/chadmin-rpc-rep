package main

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	rpch "github.com/tidusant/chadmin-repo/cuahang"
	"github.com/tidusant/chadmin-repo/models"

	"encoding/json"
	"time"

	"flag"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

const (
	defaultcampaigncode string = "XVsdAZGVmY"
)

type Arith int

func (t *Arith) Run(data string, result *string) error {
	log.Debugf("Call RPC orders args:" + data)
	*result = ""
	//parse args
	args := strings.Split(data, "|")

	if len(args) < 3 {
		return nil
	}
	var usex models.UserSession
	usex.Session = args[0]
	usex.Action = args[2]
	info := strings.Split(args[1], "[+]")
	usex.UserID = info[0]
	ShopID := info[1]
	usex.Params = ""
	if len(args) > 3 {
		usex.Params = args[3]
	}

	//check shop permission
	shop := rpch.GetShopById(usex.UserID, ShopID)
	if shop.Status == 0 {
		*result = c3mcommon.ReturnJsonMessage("-4", "Shop is disabled.", "", "")
		return nil
	}
	usex.Shop = shop

	if usex.Action == "la" {
		*result = LoadAll(usex)
	} else if usex.Action == "l3" {
		*result = Load3Month(usex)
	} else if usex.Action == "l6" {
		*result = Load6Month(usex)
	} else if usex.Action == "l9" {
		*result = Load9Month(usex)
	} else if usex.Action == "l12" {
		*result = Load12Month(usex)
	} else { //default
		*result = c3mcommon.ReturnJsonMessage("-5", "Action not found.", "", "")
	}

	return nil
}

func LoadAll(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, 0, -time.Now().Day()), time.Now().AddDate(0, 1, -time.Now().Day()))
	camp.Name = time.Now().Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}
func Load3Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, 0, -time.Now().Day()), time.Now().AddDate(0, 1, -time.Now().Day()))
	camp.Name = time.Now().Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -1, -time.Now().Day()), time.Now().AddDate(0, 0, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -2, -time.Now().Day()), time.Now().AddDate(0, -1, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func Load6Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, 0, -time.Now().Day()), time.Now().AddDate(0, 1, -time.Now().Day()))
	camp.Name = time.Now().Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -1, -time.Now().Day()), time.Now().AddDate(0, 0, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -2, -time.Now().Day()), time.Now().AddDate(0, -1, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -3, -time.Now().Day()), time.Now().AddDate(0, -2, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -3, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -4, -time.Now().Day()), time.Now().AddDate(0, -3, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -4, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -5, -time.Now().Day()), time.Now().AddDate(0, -4, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -5, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func Load9Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, 0, -time.Now().Day()), time.Now().AddDate(0, 1, -time.Now().Day()))
	camp.Name = time.Now().Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -1, -time.Now().Day()), time.Now().AddDate(0, 0, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -2, -time.Now().Day()), time.Now().AddDate(0, -1, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -3, -time.Now().Day()), time.Now().AddDate(0, -2, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -3, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -4, -time.Now().Day()), time.Now().AddDate(0, -3, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -4, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -5, -time.Now().Day()), time.Now().AddDate(0, -4, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -5, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -6, -time.Now().Day()), time.Now().AddDate(0, -5, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -6, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -7, -time.Now().Day()), time.Now().AddDate(0, -6, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -7, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -8, -time.Now().Day()), time.Now().AddDate(0, -7, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -8, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func Load12Month(usex models.UserSession) string {

	//default current month
	var camps []models.Campaign
	camp := rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, 0, -time.Now().Day()), time.Now().AddDate(0, 1, -time.Now().Day()))
	camp.Name = time.Now().Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -1, -time.Now().Day()), time.Now().AddDate(0, 0, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -1, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -2, -time.Now().Day()), time.Now().AddDate(0, -1, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -2, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -3, -time.Now().Day()), time.Now().AddDate(0, -2, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -3, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -4, -time.Now().Day()), time.Now().AddDate(0, -3, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -4, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -5, -time.Now().Day()), time.Now().AddDate(0, -4, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -5, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -6, -time.Now().Day()), time.Now().AddDate(0, -5, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -6, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -7, -time.Now().Day()), time.Now().AddDate(0, -6, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -7, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -8, -time.Now().Day()), time.Now().AddDate(0, -7, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -8, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -9, -time.Now().Day()), time.Now().AddDate(0, -8, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -9, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -10, -time.Now().Day()), time.Now().AddDate(0, -9, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -10, 0).Month().String()
	camps = append(camps, camp)

	camp = rpch.GetOrdersReportByRange(usex.Shop.ID.Hex(), time.Now().AddDate(0, -11, -time.Now().Day()), time.Now().AddDate(0, -10, -time.Now().Day()))
	camp.Name = time.Now().AddDate(0, -11, 0).Month().String()
	camps = append(camps, camp)

	info, _ := json.Marshal(camps)
	strrt := string(info)
	return c3mcommon.ReturnJsonMessage("1", "", "success", strrt)
}

func main() {
	var port int
	var debug bool
	flag.IntVar(&port, "port", 9888, "help message for flagname")
	flag.BoolVar(&debug, "debug", false, "Indicates if debug messages should be printed in log files")
	flag.Parse()

	logLevel := log.DebugLevel
	if !debug {
		logLevel = log.InfoLevel

	}

	log.SetOutputFile(fmt.Sprintf("adminReport-"+strconv.Itoa(port)), logLevel)
	defer log.CloseOutputFile()
	log.RedirectStdOut()

	//init db
	arith := new(Arith)
	rpc.Register(arith)
	log.Infof("running with port:" + strconv.Itoa(port))

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	c3mcommon.CheckError("rpc dail:", err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	c3mcommon.CheckError("rpc init listen", err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}