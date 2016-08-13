package todos

import (
	"database/sql"
	"log"
	"encoding/json"

	"github.com/kataras/iris"

	"reactizer-go/api/utils"
)

type get struct {
	db *sql.DB
}

func (t *get) Serve(c *iris.Context) {
	T := utils.GetT(c)
	_, err := utils.Authorize(c, t.db)
	if err != nil {
		c.SetStatusCode(401)
		c.Write(T(err.Error()))
		return
	}

	rows, err := t.db.Query("SELECT * FROM todo")
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()

	todos, err := scanTodos(rows)
	if err != nil {
		log.Print(err)
		return
	}

	data, err := json.Marshal(todos)
	if err != nil {
		log.Print(err)
		return
	}

	c.Write(string(data))
}