package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"text/template"
	"time"
)

var temp = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="95" height="20">
    <linearGradient id="s" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
        <stop offset="1" stop-opacity=".1"/>
    </linearGradient>
    <clipPath id="r">
        <rect width="95" height="20" rx="3" fill="#fff"/>
    </clipPath>
    <g clip-path="url(#r)">
        <rect width="64" height="20" fill="#555"/>
        <rect x="64" width="31" height="20" fill="#97ca00"/>
        <rect width="95" height="20" fill="url(#s)"/>
    </g>
    <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif"
       text-rendering="geometricPrecision" font-size="110">
        <image x="5" y="3" width="14" height="14"
               xlink:href="data:image/svg+xml;base64,PHN2ZyBmaWxsPSJ3aGl0ZXNtb2tlIiByb2xlPSJpbWciIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48dGl0bGU+R2l0SHViIGljb248L3RpdGxlPjxwYXRoIGQ9Ik0xMiAuMjk3Yy02LjYzIDAtMTIgNS4zNzMtMTIgMTIgMCA1LjMwMyAzLjQzOCA5LjggOC4yMDUgMTEuMzg1LjYuMTEzLjgyLS4yNTguODItLjU3NyAwLS4yODUtLjAxLTEuMDQtLjAxNS0yLjA0LTMuMzM4LjcyNC00LjA0Mi0xLjYxLTQuMDQyLTEuNjFDNC40MjIgMTguMDcgMy42MzMgMTcuNyAzLjYzMyAxNy43Yy0xLjA4Ny0uNzQ0LjA4NC0uNzI5LjA4NC0uNzI5IDEuMjA1LjA4NCAxLjgzOCAxLjIzNiAxLjgzOCAxLjIzNiAxLjA3IDEuODM1IDIuODA5IDEuMzA1IDMuNDk1Ljk5OC4xMDgtLjc3Ni40MTctMS4zMDUuNzYtMS42MDUtMi42NjUtLjMtNS40NjYtMS4zMzItNS40NjYtNS45MyAwLTEuMzEuNDY1LTIuMzggMS4yMzUtMy4yMi0uMTM1LS4zMDMtLjU0LTEuNTIzLjEwNS0zLjE3NiAwIDAgMS4wMDUtLjMyMiAzLjMgMS4yMy45Ni0uMjY3IDEuOTgtLjM5OSAzLS40MDUgMS4wMi4wMDYgMi4wNC4xMzggMyAuNDA1IDIuMjgtMS41NTIgMy4yODUtMS4yMyAzLjI4NS0xLjIzLjY0NSAxLjY1My4yNCAyLjg3My4xMiAzLjE3Ni43NjUuODQgMS4yMyAxLjkxIDEuMjMgMy4yMiAwIDQuNjEtMi44MDUgNS42MjUtNS40NzUgNS45Mi40Mi4zNi44MSAxLjA5Ni44MSAyLjIyIDAgMS42MDYtLjAxNSAyLjg5Ni0uMDE1IDMuMjg2IDAgLjMxNS4yMS42OS44MjUuNTdDMjAuNTY1IDIyLjA5MiAyNCAxNy41OTIgMjQgMTIuMjk3YzAtNi42MjctNS4zNzMtMTItMTItMTIiLz48L3N2Zz4="/>

        <text x="415" y="140" transform="scale(.1)" fill="#fff" textLength="370">{{.Key}}</text>

        <text x="785" y="140" transform="scale(.1)" fill="#fff" textLength="210">{{.Value}}</text>
    </g>
</svg>`

type foo struct {
	Key   string
	Value string
}

var m map[string]int
var mu sync.Mutex

func endurance() {

	m = map[string]int{}

	var file *os.File

	file, err := os.OpenFile("./newbee.data", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(data, &m)
		if err != nil {
			panic(err)
		}
	}

	go func() {

		for {
			time.Sleep(time.Second * 1)
			mu.Lock()

			data, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			mu.Unlock()

			_, err = file.WriteAt(data, 0)
			if err != nil {
				panic(err)
			}

		}
	}()

}

func main() {
	var addr string
	var certfile string
	var keyfile string
	flag.StringVar(&addr, "addr", ":443", "")
	flag.StringVar(&certfile, "cert", "./certfile", "")
	flag.StringVar(&keyfile, "key", "./keyfile", "")

	endurance()
	e := gin.Default()
	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	e.GET("/:repo/visits", func(context *gin.Context) {

		context.Header("content-type", "text/html")

		mu.Lock()

		m[context.Param("repo")] += 1

		count := m[context.Param("repo")]

		var value string
		if count > 1000 {

			var conv float64
			conv = float64(count) / 1000

			value = strconv.FormatFloat(conv, 'f', 1, 64) + "k"
		} else {
			value = strconv.Itoa(count)
		}

		t.Execute(context.Writer, foo{Key: "visits", Value: value})

		mu.Unlock()

	})

	e.RunTLS(addr, certfile, keyfile)

}
