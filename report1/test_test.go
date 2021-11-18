package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDifference(t *testing.T) {

	p1 := RecordProxy{Amount: 10, Description: "a", Date: "2021-07-06", ID: "zoAr"}
	p2 := RecordProxy{Amount: 8, Description: "h", Date: "2021-08-01", ID: "zoir"}
	p3 := RecordProxy{Amount: 3, Description: "g", Date: "2021-07-01", ID: "zoib"}

	s1 := RecordSource{Date: "2021-07-06", ID: "zoAr", Amount: 10, Description: "a"}
	s2 := RecordSource{Date: "2021-08-01", ID: "zoir", Amount: 8, Description: "h"}
	s3 := RecordSource{Date: "2021-07-01", ID: "zoib", Amount: 3, Description: "g"}
	sampleProxy1 := []RecordProxy{
		p1, p2, p3,
	}

	sampleSource1 := []RecordSource{
		s1, s2, s3,
	}

	tests := []struct {
		name string

		proxy  []RecordProxy
		source []RecordSource

		report []Reports
		err    error
	}{
		{"blankInput", nil, nil, nil, errors.New("Blank input")},
		{"SourceNilInput", sampleProxy1, nil, nil, errors.New("Blank input")},
		{"ProxyNilInput", nil, sampleSource1, nil, errors.New("Blank input")},
		{"sameInput", sampleProxy1, sampleSource1, []Reports{}, errors.New("Blank input")},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			resultProxy, resultSource := FindDifference(item.source, item.proxy)
			assert.Equal(t, item.report, resultProxy)
			assert.Equal(t, item.report, resultSource)
		})
	}
}
