package probe

import (
	"net/url"
	"strconv"
	"time"
	"fmt"
)

type ProbeRecord struct {
	Timestamp  int64        `json:"timestamp"`
	Mac        string       `json:"mac"`
	SequenceId int          `json:"sequence_id"`
	SsId       string       `json:"ssid"`
	Rssi       int          `json:"rssi"`
	ApName     string       `json:"ap_name,omitempty"`
}

func (r *ProbeRecord) Values() (values url.Values) {
	values.Add("timestamp", strconv.FormatInt(r.Timestamp, 10))
	values.Add("mac", r.Mac)
	values.Add("sequence_id", strconv.Itoa(r.SequenceId))
	values.Add("ssid", r.SsId)
	values.Add("rssi", strconv.Itoa(r.Rssi))
	if r.ApName != "" {
		values.Add("ap_name", r.ApName)
	}
	return
}

func (r *ProbeRecord) String() string {
	ltime := time.Unix(r.Timestamp, 0).Local()
	if r.ApName != "" {
		return fmt.Sprintf("%s,%d,%s,%d,%s,%d,%s", ltime.Format(time.RFC3339), r.Timestamp, r.Mac, r.SequenceId, r.SsId, r.Rssi, r.ApName)
	}
	return fmt.Sprintf("%s,%d,%s,%d,%s,%d", ltime.Format(time.RFC3339), r.Timestamp, r.Mac, r.SequenceId, r.SsId, r.Rssi)
}
