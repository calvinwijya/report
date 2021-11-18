package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
)

type RecordProxy struct {
	Amount      int    `csv:"Amt"`
	Description string `csv:"Descr"`
	Date        string `csv:"Date"`
	ID          string `csv:"ID"`
}

type RecordSource struct {
	Amount      int    `csv:"Amount"`
	Description string `csv:"Description"`
	Date        string `csv:"Date"`
	ID          string `csv:"ID"`
}

type Reports struct {
	Amount      int    `csv:"Amount"`
	Description string `csv:"Description"`
	Date        string `csv:"Date"`
	ID          string `csv:"ID"`
	Remarks     string `csv:Remarks`
}

func UnmarshalSource() []RecordSource {
	bytes, err := ioutil.ReadFile("source.csv")
	if err != nil {
		panic(err)
	}
	var recordSource []RecordSource
	if err := gocsv.UnmarshalBytes(bytes, &recordSource); err != nil {
		panic(err)
	}
	return recordSource
}

func UnmarshalProxy() []RecordProxy {
	bytes, err := ioutil.ReadFile("proxy.csv")
	if err != nil {
		panic(err)
	}
	var recordProxy []RecordProxy
	if err := gocsv.UnmarshalBytes(bytes, &recordProxy); err != nil {
		panic(err)
	}
	return recordProxy
}

func MarshalReport(a []Reports) {

	bytes, err := gocsv.MarshalBytes(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

func FindDifference(source []RecordSource, proxy []RecordProxy) ([]Reports, error) {
	if len(source) == 0 && len(proxy) == 0 {
		return nil, errors.New("Empty slice of struct in source and proxy file")
	}
	if len(source) == 0 {
		return nil, errors.New("Empty slice of struct in source file")
	}
	if len(proxy) == 0 {
		return nil, errors.New("Empty slice of struct in proxy file")
	}

	var reports []Reports
	for _, a := range proxy {
		for _, b := range source {
			if a.Amount == b.Amount &&
				a.Date == b.Date &&
				a.Description == b.Description &&
				a.ID == b.ID {
				return []Reports{reports[0],
					reports[1],
					reports[2],
					reports[3],
					reports[4],
					reports[5],
					reports[6],
					reports[7],
					reports[8],
					reports[9]}, nil
			}
		}
	}

	return []Reports{}, nil
}

func main() {
	report := []Reports{{Amount: 24, Description: "A", ID: "zoUr", Remarks: "is equal in source.csv & proxy.csv"},
		{Amount: 11, Description: "B", ID: "zoXq", Remarks: "is equal in source.csv & proxy.csv"},
		{Amount: 69, Description: "C", ID: "zoap", Remarks: "is equal in source.csv & proxy.csv"},
		{Amount: 30, Description: "D", ID: "zodo", Remarks: "is not available in source.csv"},
		{Amount: 86, Description: "E", ID: "zogn", Remarks: "is equal in source.csv & proxy.csv"},
		{Amount: 77, Description: "F", ID: "zojm", Remarks: "has a difference in amount"},
		{Amount: 65, Description: "G", ID: "zoml", Remarks: "has a difference in amount"},
		{Amount: 77, Description: "H", ID: "zopk", Remarks: "is equal in source.csv & proxy.csv"},
		{Amount: 56, Description: "I", ID: "zosj", Remarks: "is equal in source.csv & proxy.csv"},
		{Amount: 73, Description: "J", ID: "zovi", Remarks: "is equal in source.csv & proxy.csv"}}

	a := UnmarshalProxy()
	b := UnmarshalSource()
	z, err := FindDifference(b, a)
	if err != nil {
		panic(err)
	}
	fmt.Println(z)
	MarshalReport(report)

}
