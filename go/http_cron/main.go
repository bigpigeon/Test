// 请不要使用在任何商业用途上
// 使用前请安装 github.com/gin-gonic/gin
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Task struct {
	Cmd      string
	Args     []string
	Interval int
	Ticker   *time.Ticker
}

func (t *Task) Run() (string, error) {
	out, err := exec.Command(t.Cmd, t.Args...).Output()
	return string(out), err
}

type Registry map[string]*Task

// TODO 这个不是线程安全的，在context中直接使用有可能会出问题，可以使用go chan 去解决这个问题
func (r Registry) Stop(id string) error {
	if task, ok := r[id]; ok && task != nil {
		task.Ticker.Stop()
		r[id] = nil
		return nil
	} else {
		return fmt.Errorf("The task %s is not found.", id)
	}
}

var (
	registry = Registry{}
)

// TODO 这个不是线程安全的，在context中直接使用有可能会出问题，可以使用go chan 去解决这个问题
func NewTask(id string, interval int, cmd string, args ...string) (*Task, error) {
	if id == "" {
		return nil, fmt.Errorf("cannot register nil id.", id)
	}
	if val, ok := registry[id]; ok && val != nil {
		return nil, fmt.Errorf("The task %s already exists.", id)
	}
	registry[id] = &Task{
		Cmd:      cmd,
		Args:     args,
		Interval: interval,
		Ticker:   time.NewTicker(time.Duration(interval) * time.Millisecond),
	}
	go func() {
		task := registry[id]
		for {
			_, ok := <-task.Ticker.C
			if ok == false {
				return
			}
			out, err := task.Run()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(out)
		}
	}()
	return registry[id], nil
}

type RegisterTaskRequest struct {
	ID       string   `json:"id" binding:"required"`
	Cmd      string   `json:"cmd" binding:"required"`
	Args     []string `json:"args"`
	Interval int      `json:"interval" binding:"required"`
}

func RegisterTask(c *gin.Context) {
	var task RegisterTaskRequest
	if err := c.BindJSON(&task); err == nil {
		_, err := NewTask(task.ID, task.Interval, task.Cmd, task.Args...)
		if err != nil {
			c.JSON(409, gin.H{
				"ok":    false,
				"error": err.Error(),
			})

		} else {
			c.JSON(200, gin.H{
				"ok": true,
				"id": task.ID,
			})
		}
	} else {
		c.JSON(400, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
	}
}

type StopTaskRequest struct {
	ID string `json:"id" binding:"required"`
}

func StopTask(c *gin.Context) {
	var req StopTaskRequest
	if err := c.BindJSON(&req); err == nil {
		registry.Stop(req.ID)
		c.JSON(200, gin.H{
			"ok": true,
			"id": req.ID,
		})
	} else {
		c.JSON(400, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
	}

}

func main() {
	defaultPort := 4567
	var err error
	if len(os.Args) > 1 {
		defaultPort, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	r := gin.Default()
	r.POST("/", RegisterTask)
	r.DELETE("/", StopTask)
	r.Run(fmt.Sprintf(":%d", defaultPort))
}
