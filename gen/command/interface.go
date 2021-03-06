package command

import "vwood/app/gen/domain/model"

type Command interface {
	Add(Command)
	Execute()
}

type GenBase struct {
	BasePath string
	Entities []*model.Entity
}
