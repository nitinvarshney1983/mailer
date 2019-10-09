package persistence

import (
	"sync"
)

//EmailTemplate represents the collection emailTemplate
type EmailTemplate struct {
	templateName    string
	subject         string
	templateContent string
}

//EmailTemplateRepo to perform CRUD operation against EmailTemplate
type EmailTemplateRepo Repo

var instance *EmailTemplateRepo
var once sync.Once

//GetInstance for singleton object
func GetInstance() *EmailTemplateRepo {
	once.Do(func() {
		instance = &EmailTemplateRepo{}
	})
	return instance
}

func (etRepo *EmailTemplateRepo) getTemplateByName(templateName string) (*EmailTemplate, error) {

	//etRepo.col.FindOne(context.TODO())
	return nil, nil

}
