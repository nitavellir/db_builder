package exec

import (
	"os"
	"path/filepath"
)

func (h *Handler) Execute() int {
	if h.CSVPath == "" {
		return h.sendError("CSV file path is not specified.")
	} else if h.MysqlDriver.Table == "" {
		return h.sendError("Table name is not specified.")
	}

	if string([]rune(h.CSVPath)[0]) != "/" {
		wd, err := os.Getwd()
		if err != nil {
			return h.sendError("Can not get the working directory.")
		}
		h.CSVPath = filepath.Join(wd, h.CSVPath)
	}

	//proc CSV
	if err := h.procCSV(); err != nil {
		return h.sendError(err.Error())
	}

	//create schema
	if err := h.createSchema(); err != nil {
		return h.sendError(err.Error())
	}

	return 0
}

//private

func (h *Handler) sendError(msg string) int {
	h.ErrorMsg = msg
	return 1
}
