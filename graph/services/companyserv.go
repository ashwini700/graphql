package services

import (
	"context"

	gormmodel "graphql/graph/gorm"
	"graphql/graph/model"

)

func (s *DbConnStruct) CreateCompany(newComp model.NewCompany) (*model.Company, error) {
	var job []*gormmodel.NewJob
	for _, sJob := range newComp.Jobs {
		nJob := gormmodel.NewJob{
			Title: sJob.Title,
		}
		job = append(job, &nJob)
	}
	comp := gormmodel.NewCompany{
		Name: newComp.Name,
		City: newComp.City,
		Jobs: job,
	}
	err := s.db.Create(&comp).Error
	if err != nil {
		return &model.Company{}, err
	}
	return &model.Company{
		CompanyID: &comp.CompanyID,
		Name:      comp.Name,
		City:      comp.City,
	}, nil
}

func (s *DbConnStruct) ViewCompanies(ctx context.Context) ([]*model.Company, error) {
	var listComp []*model.Company
	var dataComp []gormmodel.NewCompany
	tx := s.db.WithContext(ctx)
	err := tx.Find(&dataComp).Error
	if err != nil {
		return nil, err
	}

	for _, comp := range dataComp {
		gComp := model.Company{
			CompanyID: &comp.CompanyID,
			Name:      comp.Name,
			City:      comp.City,
		}
		listComp = append(listComp, &gComp)
	}
	return listComp, nil
}

func (s *DbConnStruct) FetchCompanyByID(ctx context.Context, companyId string) (*model.Company, error) {
	var compData gormmodel.NewCompany
	tx := s.db.WithContext(ctx).Where("ID = ?", companyId)
	err := tx.Find(&compData).Error
	if err != nil {
		return &model.Company{}, err
	}
	comp := model.Company{
		CompanyID: &compData.CompanyID,
		Name:      compData.Name,
		City:      compData.City,
	}

	return &comp, nil

}
