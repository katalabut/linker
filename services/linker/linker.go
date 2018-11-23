package linker

import (
	. "linker/router"
	"linker/tool/db"
)

type Linker struct {
	Db *db.MgoConnection
}

type Url struct {
	Short string `json:"short"`
	Long  string `json:"url"`
}

func (l *Linker) CreateRoutes() Routes {
	return Routes{
		Route{
			"root",
			"GET",
			"/",
			l.urlRoot,
		},
		Route{
			"show",
			"GET",
			"/{shorturl}",
			l.urlShow,
		},
		Route{
			"create",
			"POST",
			"/create",
			l.urlCreate,
		},
	}
}
