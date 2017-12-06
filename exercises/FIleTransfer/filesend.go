package FIleTransfer

import (
"log"

"github.com/kardianos/service"
	"io/ioutil"
	"fmt"
	"os"
	"net/http"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	files, err := ioutil.ReadDir("/Users/leonid_chetverikov/GolandProjects/GoLang-Training/exercises/FIleTransfer")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		upload(f.Name())
	}
}
	func upload(path string) {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		res, err := http.Post("http://127.0.0.1:5050/upload", "binary/octet-stream", file)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		message, _ := ioutil.ReadAll(res.Body)
		fmt.Printf(string(message))
	}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "filetransfer",
		DisplayName: "filetransfer",
		Description: "filetransfer",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}