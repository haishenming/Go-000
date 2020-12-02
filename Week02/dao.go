package main

import (
	"errors"
	"fmt"
)

var Err_RecordNotFind = errors.New("record not find")

func IsRecordNotFind(err error) bool {
	return errors.Is(err, Err_RecordNotFind)
}

func GetUser() error {
	return fmt.Errorf("user not found: %w", Err_RecordNotFind)
}