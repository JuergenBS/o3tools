package main

import (
	"encoding/base64"
	"flag"
	"os"

	"github.com/cryptix/go/logging"
	"github.com/o3ma/o3tools"
)

var (
	fIDpath      = flag.String("idpath", "o3id", "the id file to save the ne ID in")
	fForceCreate = flag.Bool("force", false, "overwrite an existing ID")
	fPassw       = flag.String("pass", "ABCD", "") //pw on the cmdline...
)

func main() {
	flag.Parse()
	logging.SetupLogging(nil)
	log := logging.Logger("o3createID")

	passw, err := base64.StdEncoding.DecodeString(*fPassw)
	logging.CheckFatal(err)

	if _, err := os.Stat(*fIDpath); (os.IsNotExist(err)) || (err == nil && *fForceCreate) {
		if *fForceCreate {
			log.Log("warning", "overwriting ID", "idPath", *fIDpath)
		}
		id, err := o3tools.CreateID(*fIDpath, passw)
		logging.CheckFatal(err)
		log.Log("done", "id created", "id", id, "idPath", *fIDpath)
	} else {
		log.Log("warning", "did nothing")
	}
}
