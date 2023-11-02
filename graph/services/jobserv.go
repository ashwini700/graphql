package services

import (
	"context"
	"strconv"

	gormmodel "graphql/graph/gorm"
	"graphql/graph/model"
)

func (s *DbConnStruct) CreateByCompanyId(jobs []*model.NewJob, compId string) ([]*model.Job, error) {
	companyId, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		return nil, err
	}
	for _, j := range jobs {
		job := gormmodel.NewJob{
			Title:     j.Title,
			CompanyId: companyId,
		}
		err := s.db.Create(&job).Error
		if err != nil {
			return nil, err
		}
	}
	var mComp []*model.Job
	for _, data := range jobs {
		mCompni := model.Job{
			ID:        data.ID,
			Title:     data.Title,
			CompanyID: int(companyId),
		}
		mComp = append(mComp, &mCompni)
	}
	return mComp, nil
}

func (s *DbConnStruct) FetchJobByCompanyId(ctx context.Context, companyId string) ([]*model.Job, error) {
	var listOfJobs []*model.Job
	var dataOfJobs []gormmodel.NewJob
	tx := s.db.WithContext(ctx).Where("company_id = ?", companyId)
	err := tx.Find(&dataOfJobs).Error
	if err != nil {
		return nil, err
	}
	for _, job := range dataOfJobs {
		strId := strconv.FormatUint(uint64(job.ID), 10)
		mJob := model.Job{
			ID:        &strId,
			Title:     job.Title,
			CompanyID: int(job.CompanyId),
		}
		listOfJobs = append(listOfJobs, &mJob)
	}
	return listOfJobs, nil
}

func (s *DbConnStruct) GetJobById(ctx context.Context, jobId string) (*model.Job, error) {
	// var jobData model.Job
	var dbjobData gormmodel.NewJob
	tx := s.db.WithContext(ctx).Where("ID = ?", jobId)
	err := tx.Find(&dbjobData).Error
	if err != nil {
		return &model.Job{}, err
	}
	idStr := strconv.FormatUint(uint64(dbjobData.ID), 10)
	return &model.Job{
		ID:        &idStr,
		Title:     dbjobData.Title,
		CompanyID: int(dbjobData.CompanyId),
	}, nil
}

func (s *DbConnStruct) GetAllJobs(ctx context.Context) ([]*model.Job, error) {
	var listJobs []*model.Job
	var getJobs []gormmodel.NewJob
	tx := s.db.WithContext(ctx)
	err := tx.Find(&getJobs).Error
	if err != nil {
		return nil, err
	}
	for _, job := range getJobs {
		strId := strconv.FormatUint(uint64(job.ID), 10)
		data := model.Job{
			ID:        &strId,
			Title:     job.Title,
			CompanyID: int(job.CompanyId),
		}
		listJobs = append(listJobs, &data)
	}
	return listJobs, nil
}
