package time_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	start := time.Now()
	time.Sleep(1 * time.Second)
	t.Log(time.Since(start).Seconds())
	pTime, err := time.Parse(time.RFC3339, "2002-10-02T10:00:00Z")
	assert.NoError(t, err)
	t.Log(pTime)
	{

		pTime, err := time.Parse(time.RFC3339, "2002-10-02T10:00:00+08:00")
		assert.NoError(t, err)
		t.Log(pTime)
	}
}

func TestTimeUnix(t *testing.T) {
	{
		unixTime := 1542190917636
		ti := time.Unix(int64(unixTime), 0)
		t.Log(ti)
	}
}

func TestSince(t *testing.T) {
	start := time.Now()
	time.Sleep(2)
	t.Log(time.Since(start))
}

type UtcTime time.Time

func (t *UtcTime) UnmarshalJSON(data []byte) error {
	var err error
	err = (*time.Time)(t).UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*t = UtcTime((*time.Time)(t).UTC())
	return nil
}

func (t UtcTime) MarshalJSON() ([]byte, error) {
	return time.Time(t).UTC().MarshalJSON()
}

func (t UtcTime) String() string {
	return time.Time(t).String()
}

type EmbedUtcTime struct {
	time.Time
}

func TestTimeLoc(t *testing.T) {
	time.Local = time.UTC
	date := time.Now()
	t.Log(date)
	date, err := time.Parse(time.RFC3339, "2019-01-05T18:31:27+08:00")
	require.NoError(t, err)
	t.Log(date)
	// unix test
	t.Logf("unix Local %d UTC %d", date.Unix(), date.UTC().Unix())
	assert.Equal(t, date.Unix(), date.UTC().Unix())
	// with json
	{

		type TestData struct {
			Time UtcTime `json:"Time"`
		}
		var data TestData
		err = json.Unmarshal([]byte(`{"Time": "2019-01-05T18:31:27+08:00"}`), &data)
		require.NoError(t, err)
		t.Logf("%s\n", data)

		var iface interface{} = data.Time
		_, ok := iface.(time.Time)
		t.Log("Utc time to time.Time", ok)
		iface = EmbedUtcTime{time.Now()}
		_, ok = iface.(time.Time)
		t.Log("embed Utc time to time.Time", ok)
	}

}

func TestTimeNowCache(t *testing.T) {
	timestamp := time.Now().Unix()
	t.Log(timestamp)
}

func BenchmarkGetTime(b *testing.B) {

	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func TestGetUseTime(t *testing.T) {
	defer func(now time.Time) { t.Log("use time", time.Since(now)) }(time.Now())
	time.Sleep(1 * time.Second)
}

func TestDurationParse(t *testing.T) {
	{
		ti, err := time.ParseDuration("1s")
		require.NoError(t, err)
		t.Log(ti)
	}
	{
		ti, err := time.ParseDuration("1.5s")
		require.NoError(t, err)
		t.Log(ti)
	}
	{
		ti, err := time.ParseDuration("0.05s")
		require.NoError(t, err)
		t.Log(ti)
	}
	{
		d := 100*time.Microsecond + 100*time.Millisecond + 20*time.Second + 1*time.Hour
		t.Logf(d.String())
	}
}
