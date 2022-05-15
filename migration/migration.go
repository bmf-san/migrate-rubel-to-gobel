package migration

import (
	gar "github.com/bmf-san/migrate-rubel-to-gobel/gobel/repository"
	rar "github.com/bmf-san/migrate-rubel-to-gobel/rubel/repository"
)

type Migration struct {
	RubelAdminRepository    *rar.RubelAdminRepository
	GobelAdminRepository    *gar.GobelAdminRepository
	RubelCategoryRepository *rar.RubelCategoryRepository
	GobelCategoryRepository *gar.GobelCategoryRepository
	RubelTagRepository      *rar.RubelTagRepository
	GobelTagRepository      *gar.GobelTagRepository
	RubelTagPostRepository  *rar.RubelTagPostRepository
	GobelTagPostRepository  *gar.GobelTagPostRepository
	RubelPostRepository     *rar.RubelPostRepository
	GobelPostRepository     *gar.GobelPostRepository
}

func NewMigration(
	rar *rar.RubelAdminRepository,
	gar *gar.GobelAdminRepository,
	rcr *rar.RubelCategoryRepository,
	gcr *gar.GobelCategoryRepository,
	rtr *rar.RubelTagRepository,
	gtr *gar.GobelTagRepository,
	rtpr *rar.RubelTagPostRepository,
	gtpr *gar.GobelTagPostRepository,
	rpr *rar.RubelPostRepository,
	gpr *gar.GobelPostRepository,
) *Migration {
	return &Migration{
		RubelAdminRepository:    rar,
		GobelAdminRepository:    gar,
		RubelCategoryRepository: rcr,
		GobelCategoryRepository: gcr,
		RubelTagRepository:      rtr,
		GobelTagRepository:      gtr,
		RubelTagPostRepository:  rtpr,
		GobelTagPostRepository:  gtpr,
		RubelPostRepository:     rpr,
		GobelPostRepository:     gpr,
	}
}

func (mgt *Migration) Run() {
	migrateAdmin(mgt)
	migrateCategory(mgt)
	migrateTag(mgt)
	migrateTagPost(mgt)
	migratePost(mgt)
}

func migrateAdmin(mgt *Migration) {
	admins, err := mgt.RubelAdminRepository.Read()
	if err != nil {
		panic(err)
	}

	for _, a := range admins {
		cvt := a.Convert()
		_, err := mgt.GobelAdminRepository.Write(cvt)
		if err != nil {
			panic(err)
		}
	}
}

func migrateCategory(mgt *Migration) {
	categories, err := mgt.RubelCategoryRepository.Read()
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		cvt := c.Convert()
		_, err := mgt.GobelCategoryRepository.Write(cvt)
		if err != nil {
			panic(err)
		}
	}
}

func migrateTag(mgt *Migration) {
	tags, err := mgt.RubelTagRepository.Read()
	if err != nil {
		panic(err)
	}

	for _, c := range tags {
		cvt := c.Convert()
		_, err := mgt.GobelTagRepository.Write(cvt)
		if err != nil {
			panic(err)
		}
	}
}

func migrateTagPost(mgt *Migration) {
	tagposts, err := mgt.RubelTagPostRepository.Read()
	if err != nil {
		panic(err)
	}

	for _, tp := range tagposts {
		cvt := tp.Convert()
		_, err := mgt.GobelTagPostRepository.Write(cvt)
		if err != nil {
			panic(err)
		}
	}
}

func migratePost(mgt *Migration) {
	posts, err := mgt.RubelPostRepository.Read()
	if err != nil {
		panic(err)
	}

	for _, p := range posts {
		cvt := p.Convert()
		_, err := mgt.GobelPostRepository.Write(cvt)
		if err != nil {
			panic(err)
		}
	}
}
