package ffmpeg_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestAacToWav(t *testing.T) {
	f, err := os.Open("input.aac")
	assert.NoError(t, err)
	//cmd := exec.Command(
	//	"ffmpeg",
	//	"-i",
	//	"/dev/stdin",
	//	"-ar", "16000",
	//	"-ac",
	//	"1",
	//	"-f",
	//	"wav",
	//	"-")
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
