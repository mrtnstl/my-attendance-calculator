package cmd

import (
	"fmt"
	"strings"
	"time"
)

type DayRecord struct {
	In       time.Time `json:"in"`
	Out      time.Time `json:"out"`
	Fraction int64     `json:"fraction"` // remaining or gained minutes for that day
}

func ParseFile(fContent string) ([]DayRecord, error) {
	if fContent == "" {
		return nil, fmt.Errorf("err: %s\n", "empty file")
	}

	fByLine := strings.Split(fContent, "\n")
	if len(fByLine) < 1 {
		return nil, fmt.Errorf("err: %s\n", "empty file")
	}

	var dayRecords []DayRecord

	for i, line := range fByLine {
		var dayRec DayRecord

		lBySemiC := strings.Split(line, ";")
		if len(lBySemiC) != 3 {
			return nil, fmt.Errorf("err: %s %d\n", "malformed record on line", i+1)
		}

		date := lBySemiC[0]
		if date == "" {
			return nil, fmt.Errorf("err: %s %d\n", "empty date on line", i+1)
		}
		in := lBySemiC[1]
		if in == "" {
			return nil, fmt.Errorf("err: %s %d\n", "empty clock in time on line", i+1)
		}
		out := lBySemiC[2]
		if out == "" {
			return nil, fmt.Errorf("err: %s %d\n", "empty clock out time on line", i+1)
		}

		dateTimeStrIn := fmt.Sprintf("%s %s", date, in)
		dateTimeStrOut := fmt.Sprintf("%s %s", date, out)

		tIn, err := time.Parse("2006.1.2. 15:04", dateTimeStrIn)
		if err != nil {
			return nil, fmt.Errorf("err: %s %d\n", "failed to parse datetime 'in' on line", i+1)
		}
		tOut, err := time.Parse("2006.1.2. 15:04", dateTimeStrOut)
		if err != nil {
			return nil, fmt.Errorf("err: %s %d\n", "failed to parse datetime 'out' on line", i+1)
		}

		dayRec.In = tIn
		dayRec.Out = tOut
		dayRec.Fraction = (tOut.Unix() - tIn.Unix() - (510 * 60)) / 60
		dayRecords = append(dayRecords, dayRec)
	}

	return dayRecords, nil
}
