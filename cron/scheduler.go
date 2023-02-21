package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"indivest-engine/services"
	"indivest-engine/utils"
	"time"
)

type Cron struct {
	Sc         *gocron.Scheduler
	SandboxSrv *services.SandboxServiceConfig
	MfSrv      *services.MFService
}

func CreateScheduler() *gocron.Scheduler {
	//go timer()
	return gocron.NewScheduler(time.Local)
}

func timer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {

			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
}
func (cron *Cron) InitializeScheduler() {
	cron.SIPJobs()

	cron.DailyReportJobs()

	cron.RedisUpdateNAVJobs()

	//list more jobs
	cron.Sc.StartAsync()
}

// cron jobs
func (cron *Cron) DailyReportJobs() {

	job := cron.Sc.Every(1).Day().At("23:59")

	_, err := job.Do(cron.DailyReport)
	if err != nil {
		return
	}
}

func (cron *Cron) SIPJobs() {

	job := cron.Sc.Every(1).Day().At("12:08")

	_, err := job.Do(cron.SIP)
	if err != nil {
		return
	}

}

func (cron *Cron) RedisUpdateNAVJobs() {
	job := cron.Sc.Every(4).Hours()

	_, err := job.Do(cron.UpdateRedis)
	if err != nil {
		return
	}

}

func (cron *Cron) UpdateFundsSupported() {
	job := cron.Sc.Every(12).Hours()

	_, err := job.Do(cron.UpdateFundsSavvy)
	if err != nil {
		return
	}

}

//functions for jobs

func (cron *Cron) SIP() {
	//	 fetch all active SIP's where date is current date
	//	deduct wallet balance --> call service

	err := cron.SandboxSrv.ProcessSIP()
	if err != nil {
		utils.Log.Error(err)
		return
	}
}

func (cron *Cron) DailyReport() {
	//	 fetch all active SIP's where date is current date
	//	deduct wallet balance --> call service

	err := cron.SandboxSrv.DailyReport()
	if err != nil {
		utils.Log.Error(err)
		return
	}

}

func (cron *Cron) UpdateRedis() {

	allFunds, err := cron.MfSrv.SavvyRepo.ReadAllFundDetails()
	if err != nil {
		utils.Log.Error(err)
		return
	}

	for _, funds := range *allFunds {
		//	set nav to funds
		key := "nav_" + funds.AMCCode
		err = cron.SandboxSrv.RedisRep.SetKeyValue(key, funds.NAV)
		if err != nil {
			utils.Log.Error(err)
			return
		}

		//utils.Log.Warn(key, funds.NAV)
	}
	utils.Log.Warnf("NAV update Cron completed")
	return
}

func (cron *Cron) UpdateFundsSavvy() {

	err := cron.MfSrv.UpdateFunds()
	if err != nil {
		utils.Log.Error(err)
	}
}
