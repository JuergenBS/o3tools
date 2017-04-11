package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"os"

	"github.com/cryptix/go/logging"
	"github.com/o3ma/o3"
)

var (
	fIDpath = flag.String("idpath", "o3id", "the id file to load the ID from")
	fABpath = flag.String("abpath", "o3ab", "the address book file to save the new ID in")
	fPassw  = flag.String("pass", "ABCD", "") // TODO: don't pw on the cmdline
	fRID    = flag.String("rid", "", "remote ID to query")
	// TODO: needed?
	// fPubNick = flag.String("nick", "alice", "publick nickname")
)

var tr o3.ThreemaRest

func main() {
	flag.Parse()
	logging.SetupLogging(nil)
	log := logging.Logger("o3rollodex")

	passw, err := base64.StdEncoding.DecodeString(*fPassw)
	logging.CheckFatal(err)

	tid, err := o3.LoadIDFromFile(*fIDpath, passw)
	logging.CheckFatal(err)

	log.Log("info", "loaded ID", "path", *fIDpath, "id", tid)

	if *fRID == "" {
		logging.CheckFatal(errors.New("flag rid can't be empty"))
	}

	// TODO: needed? tid.Nick = o3.NewPubNick(*fPubNick)
	ctx := o3.NewSessionContext(tid)

	if _, err := os.Stat(*fABpath); !os.IsNotExist(err) {
		logging.CheckFatal(ctx.ID.Contacts.ImportFrom(*fABpath))
	}

	var rid o3.ThreemaContact
	var got bool
	// check if we know the remote ID for
	// (just demonstration purposes \bc sending and receiving functions do this lookup for us)
	if rid, got = ctx.ID.Contacts.Get(*fRID); got == false {
		//retrieve the ID from Threema's servers
		ridStr := o3.NewIDString(*fRID)
		var err error
		rid, err = tr.GetContactByID(ridStr)
		logging.CheckFatal(err)

		log.Log("info", "retreived ID from directory server", "rid", ridStr)

		ctx.ID.Contacts.Add(rid)

		err = ctx.ID.Contacts.SaveTo(*fABpath)
		logging.CheckFatal(err)
		log.Log("info", "saved ID to address book")
	} else {
		log.Log("info", "ID already present in address book", "rid", rid)
	}
}
