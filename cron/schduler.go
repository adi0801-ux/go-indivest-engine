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

	//list more jobs
	cron.Sc.StartAsync()
}

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
