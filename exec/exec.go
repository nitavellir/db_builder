package exec

func (h *Handler) Execute() int {
	//init db driver
	if err := h.MysqlDriver.Init(); err != nil {
		return h.sendError(err.Error())
	}

	//proc CSV
	if err := h.procCSV(); err != nil {
		return h.sendError(err.Error())
	}

	//create schema
	if err := h.createSQL(); err != nil {
		return h.sendError(err.Error())
	}

	//create table
	if err := h.MysqlDriver.CreateTable(h.CreateTableSchema); err != nil {
		return h.sendError(err.Error())
	}

	//insert data
	if err := h.MysqlDriver.InsertData(h.InsertSQL); err != nil {
		return h.sendError(err.Error())
	}

	return 0
}

//private

func (h *Handler) sendError(msg string) int {
	h.ErrorMsg = msg
	return 1
}
