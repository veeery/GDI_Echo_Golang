package db

import "gitlab.com/veeery/gdi_echo_golang.git/model"

type Migrate struct {
	Table interface{}
}

func MigrateTable() []Migrate {
	return []Migrate{
		{Table: model.User{}},
	}
}
