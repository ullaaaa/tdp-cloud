package certbot

import (
	"tdp-cloud/helper/certmagic"
	"tdp-cloud/helper/logman"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/dborm/certjob"
	"tdp-cloud/module/dborm/vendor"
)

func Daemon() {

	jobs, err := certjob.FetchAll(&certjob.FetchAllParam{})

	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		NewByJob(job)
	}

}

func NewByJob(job *dborm.Certjob) error {

	vendor, err := vendor.Fetch(&vendor.FetchParam{
		Id: job.VendorId, UserId: job.UserId,
	})

	if err != nil || vendor.Id == 0 {
		logman.Info("Certjob Ignore Domain:", job.Domain)
		return err
	}

	return certmagic.Manage(&certmagic.Params{
		Email:     job.Email,
		Domain:    job.Domain,
		CaType:    job.CaType,
		Provider:  vendor.Provider,
		SecretId:  vendor.SecretId,
		SecretKey: vendor.SecretKey,
	})

}
