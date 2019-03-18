package ffmpeg_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func JsonFormat(v interface{}) string {
	writer := bytes.Buffer{}
	err := json.NewEncoder(&writer).Encode(v)
	if err != nil {
		panic(err)
	}
	return writer.String()
}

func TestAacToWav(t *testing.T) {
	//f, err := os.Open("input.aac")
	//assert.NoError(t, err)
	fData, err := ioutil.ReadFile("input.aac")
	require.NoError(t, err)
	f := bytes.NewBuffer(fData)
	cmd := exec.Command("ffmpeg", strings.Split("-i /dev/stdin -ar 16000 -ac 1 -f wav -", " ")...)
	cmd.Stdin = f
	var errorBuff bytes.Buffer
	cmd.Stderr = &errorBuff
	out, err := cmd.Output()
	assert.NoError(t, err)
	err = ioutil.WriteFile("output.wav", out, 0644)
	assert.NoError(t, err)
	t.Log(errorBuff.String())
}

// ffmpeg work well with file's io.Reader
func TestMp4ToWav(t *testing.T) {
	var f io.Reader
	var err error
	f, err = os.Open("test.m4a")
	require.NoError(t, err)
	cmdStr := strings.Split("-f mp4 -i /dev/stdin -ar 16000 -ac 1 -f wav -", " ")
	cmd := exec.Command("ffmpeg", cmdStr...)
	cmd.Stdin = f
	var errBuff bytes.Buffer
	cmd.Stderr = &errBuff
	d, err := cmd.Output()
	t.Log(errBuff.String())
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("output.wav", d, 0644)
	assert.NoError(t, err)

}

// but not work with bytes.Buffer
func TestMp4ToWavWithBuff(t *testing.T) {
	var data []byte
	var err error
	data, err = ioutil.ReadFile("test.m4a")
	require.NoError(t, err)

	cmdStr := strings.Split("-f mp4 -i /dev/stdin -ar 16000 -ac 1 -f wav -", " ")
	cmd := exec.Command("ffmpeg", cmdStr...)
	cmd.Stdin = bytes.NewBuffer(data)
	var errBuff bytes.Buffer
	cmd.Stderr = &errBuff
	d, err := cmd.Output()
	t.Log(errBuff.String())
	require.NoError(t, err)
	err = ioutil.WriteFile("output.wav", d, 0644)
	assert.NoError(t, err)
}

func TestMp4ToWavWithNewFile(t *testing.T) {
	// rewrite file
	{
		data, err := ioutil.ReadFile("input.m4a")
		require.NoError(t, err)
		err = ioutil.WriteFile("input_other.22", data, 0644)
		require.NoError(t, err)
	}
	f, err := os.Open("input_other.22")
	require.NoError(t, err)
	// file
	cmdStr := []string{"-f", "mp4"}
	cmdStr = append(cmdStr, strings.Split("-i /dev/stdin -ar 16000 -ac 1 -f wav -", " ")...)
	cmd := exec.Command("ffmpeg", cmdStr...)
	cmd.Stdin = f
	var errBuff bytes.Buffer
	cmd.Stderr = &errBuff
	d, err := cmd.Output()
	if err != nil {
		t.Log(errBuff.String())
		t.Error(err)
	}

	err = ioutil.WriteFile("output.wav", d, 0644)
	assert.NoError(t, err)
}
