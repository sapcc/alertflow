package handlers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/sapcc/alertflow/pkg/clients"
)

type WebHookMsg struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}

// Alert is a single alert.
type Alert struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    string            `json:"startsAt,omitempty"`
	EndsAt      string            `json:"EndsAt,omitempty"`
	Fingerprint string            `json:"fingerprint"`
}

func (msg *WebHookMsg) CheckValid() error {
	if msg == nil {
		return fmt.Errorf("invalid webhook: nil")
	}

	if msg.Receiver == "" {
		return fmt.Errorf("receiver value missing")
	}

	if msg.Alerts == nil || len(msg.Alerts) == 0 {
		return fmt.Errorf("alerts list is empty or invalid")
	}

	// TODO: any more?
	return nil
}

func (alert *Alert) CheckValid() error {
	if alert.Status == "" {
		return fmt.Errorf("Status value missing")
	}

	projectId := alert.GetProjectId()
	if projectId == "" {
		return fmt.Errorf("label project_id is missing or invalid")
	}

	// TODO: any more?
	return nil
}

func (alert *Alert) GetProjectId() string {
	if alert.Labels == nil || len(alert.Labels) == 0 {
		return ""
	}

	for k, _ := range alert.Labels {
		if k == "project_id" {
			return alert.Labels[k]
		}
	}

	return ""
}

func (alert *Alert) GetFingerprint() string {
	return alert.Fingerprint
}

func AlertWebHookHandler(projectClient *clients.ProjectClient) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Printf("error: invalid request method received")
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Printf("error: error reading request body")
			http.Error(w, "error reading request body", http.StatusInternalServerError)
			return
		}

		var whMsg WebHookMsg
		if err := json.Unmarshal(body, &whMsg); err != nil {
			logger.Printf("error: failed to unmarshal body object")
			http.Error(w, "failed to unmarshal body object", http.StatusInternalServerError)
			return
		}

		// check for valid alerts
		err = whMsg.CheckValid()
		if err != nil {
			logger.Printf("error: invalid alert webhook payload received, err: %s", err)
			http.Error(w, "invalid alert payload", http.StatusInternalServerError)
			return
		}

		for _, alert := range whMsg.Alerts {
			err = alert.CheckValid()
			if err != nil {
				logger.Printf("error: no email sent. invalid alert found, err: %s", err)
				continue
			}

			// valid alert -> fetch alert attributes
			// TODO: optimize to fetch from cache
			projectId := alert.GetProjectId()
			fingerprint := alert.GetFingerprint()

			// get project
			project, err := projectClient.GetProject(projectId)
			if err != nil {
				logger.Printf("error: failed to fetch project with id: %s, err: %+v", projectId, err)
				continue
			}

			// fetch contact infos
			PrimaryContactEmail := project.ResponsiblePrimaryContactEmail
			OperatorEmail := project.ResponsibleOperatorEmail

			logger.Printf("sending email to primary contact:%s, operator:%s for alert, fingerprint: %s, project_id: %s\n", PrimaryContactEmail, OperatorEmail, fingerprint, projectId)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("status: ok"))
	})
}
