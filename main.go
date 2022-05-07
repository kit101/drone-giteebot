package main

import (
	"encoding/json"
	"github.com/drone/go-scm/scm"
	"github.com/go-chi/chi"
	giteePullsConfig "github.com/kit101/drone-plugin-gitee-pulls/config"
	giteePullsPlugin "github.com/kit101/drone-plugin-gitee-pulls/plugins"
	"github.com/kit101/gitee-bot/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := chi.NewRouter()
	r.Post("/hook", func(w http.ResponseWriter, r *http.Request) {
		webhookRequest := parseRequest(r)
		log.WithFields(log.Fields{"webhookRequest": webhookRequest}).Debug()
		err := validateSecret(webhookRequest)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		// TODO do biz
		webhookData := webhookRequest.WebhookData
		status := getBuildStatus(webhookData)
		buildLink := getBuildLink(webhookData)
		config := giteePullsConfig.Config{
			GiteeApiUrl:       "https://gitee.com/api/v5",
			AccessToken:       "57eea8fbd86fc8f3bc107c7781ef977b",
			DroneProto:        webhookData.System.Proto,
			DroneHost:         webhookData.System.Host,
			Repo:              webhookData.Repo.Slug,
			PullRequestNumber: scm.ExtractPullRequest(webhookData.Build.Ref),
			BuildLink:         buildLink,
			BuildStatus:       status,
			CommitRef:         webhookData.Build.Ref,
			IsRunning:         webhookData.Event == "build" && webhookData.Action == "created" || status == giteePullsConfig.BuildStatusRunning,
			PluginComment: giteePullsConfig.Comment{
				Disabled: false,
			},
			PluginLabel: giteePullsConfig.Label{
				Disabled: false,
				Running:  "drone-build/running,E6A23C",
				Success:  "drone-build/success,67C23A",
				Failure:  "drone-build/failure,DB2828",
			},
			PluginTest: giteePullsConfig.Test{
				Disabled: false,
			},
		}
		plugin := giteePullsPlugin.NewPlugin(config)
		err = plugin.Exec()
		if err != nil {
			log.Errorf(err.Error())
			w.Write([]byte("failure"))
			return
		}
		w.Write([]byte("success"))
	})
	http.ListenAndServe(":32222", r)
}

func parseRequest(r *http.Request) model.WebhookRequest {
	droneEvent := r.Header.Get("X-Drone-Event")
	digest := r.Header.Get("digest")
	date := r.Header.Get("Date")
	signature := r.Header.Get("Signature")
	defer r.Body.Close()
	webhookData := model.WebhookData{}
	json.NewDecoder(r.Body).Decode(&webhookData)
	return model.WebhookRequest{
		DroneEvent:  droneEvent,
		Digest:      digest,
		Date:        date,
		Signature:   signature,
		WebhookData: webhookData,
	}
}

// validateSecret TODO
func validateSecret(request model.WebhookRequest) error {
	return nil
}

func getBuildLink(data model.WebhookData) string {
	return data.System.Link + "/" + data.Repo.Slug + "/" + strconv.FormatInt(data.Build.Number, 10)
}

/*
 	TODO all status
	StatusSkipped  = "skipped"
	StatusBlocked  = "blocked"
	StatusDeclined = "declined"
	StatusWaiting  = "waiting_on_dependencies"
	StatusPending  = "pending"
	StatusRunning  = "running"
	StatusPassing  = "success"
	StatusFailing  = "failure"
	StatusKilled   = "killed"
	StatusError    = "error"

	switch build.Status {
	case core.StatusPending, core.StatusRunning, core.StatusBlocked:
		io.WriteString(w, badgeStarted)
	case core.StatusPassing:
		io.WriteString(w, badgeSuccess)
	case core.StatusError:
		io.WriteString(w, badgeError)
	default:
		io.WriteString(w, badgeFailure)
	}
*/
func getBuildStatus(data model.WebhookData) giteePullsConfig.BuildStatus {
	return giteePullsConfig.BuildStatusOfValue(data.Build.Status)
}
