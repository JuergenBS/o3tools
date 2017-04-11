package o3tools

import (
	"github.com/o3ma/o3"
	"github.com/pkg/errors"
)

var tr o3.ThreemaRest

func CreateID(idpath string, pass []byte) (*o3.ThreemaID, error) {
	var err error
	tid, err := tr.CreateIdentity()
	if err != nil {
		return nil, errors.Wrap(err, "CreateIdentity failed")
	}

	if err := tid.SaveToFile(idpath, pass); err != nil {
		return nil, errors.Wrap(err, "SaveToFile failed")
	}
	return &tid, nil
}
