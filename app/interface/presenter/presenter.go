package presenter

import (
	"encoding/json"
	"io"
	"os"
)

type IPresenter interface {
	Present(i interface{}) error
}

type StdoutPresenter struct {
	io.Writer
}

func NewStdoutPresenter() IPresenter {
	return &StdoutPresenter{os.Stdout}
}

func (presenter *StdoutPresenter) Present(i interface{}) error {
	json, err := json.Marshal(i)
	if err != nil {
		return err
	}

	_, err = presenter.Write(json)
	if err != nil {
		return err
	}
	return nil
}
