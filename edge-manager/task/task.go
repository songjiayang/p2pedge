package task

import (
	"edge-manager/util"
	"encoding/json"
	"fmt"
	"github.com/computes/ipfs-http-api"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type Task struct {
	ipfsID   string
	cfg      *Config
	rootPath string
}

func NewTask(ipfsID string) *Task {
	return &Task{ipfsID: ipfsID}
}

func (t *Task) Load() error {
	// load cfg
	body, err := ipfs.Cat(util.IpfsURL, t.ipfsID)
	if err != nil {
		return err
	}
	defer body.Close()

	var cfg Config

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(buf, &cfg); err != nil {
		return err
	}

	t.cfg = &cfg

	// prepare folder
	if err = t.PrepareFolders(); err != nil {
		return err
	}

	// load package
	body, err = ipfs.Cat(util.IpfsURL, cfg.Package)
	if err != nil {
		return err
	}

	tarFile := t.rootPath + "/pacakge.tar.gz"
	f, err := os.OpenFile(tarFile, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	buf = make([]byte, 1024*3)
	for {
		n, err := body.Read(buf)
		if n > 0 {
			f.Write(buf[0:n])
		}

		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF {
			break
		}
	}

	if err = f.Close(); err != nil {
		return err
	}

	// unzip package
	return util.DeCompress(tarFile, t.rootPath+"/current/")
}

func (t *Task) Run() error {
	if err := os.Chdir(t.rootPath + "/current/"); err != nil {
		return err
	}

	if err := os.Chmod(t.cfg.Cmd, 0777); err != nil {
		return err
	}

	os.Setenv("EDAGE_APP_ID", t.cfg.ID)

	cmd := exec.Command(t.cfg.Cmd)
	return cmd.Run()
}

func (t *Task) PrepareFolders() error {
	t.rootPath = fmt.Sprintf("./runner/%s/%s", t.cfg.ID, t.ipfsID)
	return os.MkdirAll(t.rootPath, 0755)
}
