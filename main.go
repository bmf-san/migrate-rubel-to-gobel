package main

import (
	"fmt"

	"github.com/bmf-san/migrate-rubel-to-gobel/config"
	"github.com/bmf-san/migrate-rubel-to-gobel/database"
	gr "github.com/bmf-san/migrate-rubel-to-gobel/gobel/repository"
	"github.com/bmf-san/migrate-rubel-to-gobel/migration"
	rr "github.com/bmf-san/migrate-rubel-to-gobel/rubel/repository"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.NewConfig()
	conn := database.NewClient(cfg)
	defer conn.GobelConn.Close()
	defer conn.RubelConn.Close()

	rar := rr.NewRubelAdminRepository(conn.RubelConn)
	gar := gr.NewGobelAdminRepository(conn.GobelConn)
	rcr := rr.NewRubelCategoryRepository(conn.RubelConn)
	gcr := gr.NewGobelCategoryRepository(conn.GobelConn)
	rtr := rr.NewRubelTagRepository(conn.RubelConn)
	gtr := gr.NewGobelTagRepository(conn.GobelConn)
	rtpr := rr.NewRubelTagPostRepository(conn.RubelConn)
	gtpr := gr.NewGobelTagPostRepository(conn.GobelConn)
	rpr := rr.NewRubelPostRepository(conn.RubelConn)
	gpr := gr.NewGobelPostRepository(conn.GobelConn)

	mgt := migration.NewMigration(rar, gar, rcr, gcr, rtr, gtr, rtpr, gtpr, rpr, gpr)
	_, err := conn.GobelConn.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		panic(err)
	}

	var tables []string = []string{
		"admins", "categories", "tags", "tag_post", "posts",
	}
	for _, t := range tables {
		_, err = conn.GobelConn.Exec(fmt.Sprintf("TRUNCATE TABLE %s", t))
		if err != nil {
			panic(err)
		}
	}

	mgt.Run()

	_, err = conn.GobelConn.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		panic(err)
	}
}
