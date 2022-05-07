package model

type WebhookRequest struct {
	DroneEvent  string
	Digest      string
	Date        string
	Signature   string
	WebhookData WebhookData
}

type WebhookData struct {
	Event  string     `json:"event"`
	Action string     `json:"action"`
	User   User       `json:"user,omitempty"`
	Repo   Repository `json:"repo,omitempty"`
	Build  Build      `json:"build,omitempty"`
	System System     `json:"system"`
}

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Machine   bool   `json:"machine"`
	Admin     bool   `json:"admin"`
	Active    bool   `json:"active"`
	Avatar    string `json:"avatar"`
	Syncing   bool   `json:"syncing"`
	Synced    int64  `json:"synced"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
	LastLogin int64  `json:"last_login"`
	Token     string `json:"-"`
	Refresh   string `json:"-"`
	Expiry    int64  `json:"-"`
	Hash      string `json:"-"`
}

type Repository struct {
	ID            int64  `json:"id"`
	UID           string `json:"uid"`
	UserID        int64  `json:"user_id"`
	Namespace     string `json:"namespace"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	SCM           string `json:"scm"`
	HTTPURL       string `json:"git_http_url"`
	SSHURL        string `json:"git_ssh_url"`
	Link          string `json:"link"`
	Branch        string `json:"default_branch"`
	Private       bool   `json:"private"`
	Visibility    string `json:"visibility"`
	Active        bool   `json:"active"`
	Config        string `json:"config_path"`
	Trusted       bool   `json:"trusted"`
	Protected     bool   `json:"protected"`
	IgnoreForks   bool   `json:"ignore_forks"`
	IgnorePulls   bool   `json:"ignore_pull_requests"`
	CancelPulls   bool   `json:"auto_cancel_pull_requests"`
	CancelPush    bool   `json:"auto_cancel_pushes"`
	CancelRunning bool   `json:"auto_cancel_running"`
	Timeout       int64  `json:"timeout"`
	Throttle      int64  `json:"throttle,omitempty"`
	Counter       int64  `json:"counter"`
	Synced        int64  `json:"synced"`
	Created       int64  `json:"created"`
	Updated       int64  `json:"updated"`
	Version       int64  `json:"version"`
	Signer        string `json:"-"`
	Secret        string `json:"-"`
	Build         *Build `json:"build,omitempty"`
	Perms         *Perm  `json:"permissions,omitempty"`
	Archived      bool   `json:"archived"`
}
type Perm struct {
	UserID  int64  `db:"perm_user_id"  json:"-"`
	RepoUID string `db:"perm_repo_uid" json:"-"`
	Read    bool   `db:"perm_read"     json:"read"`
	Write   bool   `db:"perm_write"    json:"write"`
	Admin   bool   `db:"perm_admin"    json:"admin"`
	Synced  int64  `db:"perm_synced"   json:"-"`
	Created int64  `db:"perm_created"  json:"-"`
	Updated int64  `db:"perm_updated"  json:"-"`
}

type Build struct {
	ID           int64             `db:"build_id"             json:"id"`
	RepoID       int64             `db:"build_repo_id"        json:"repo_id"`
	Trigger      string            `db:"build_trigger"        json:"trigger"`
	Number       int64             `db:"build_number"         json:"number"`
	Parent       int64             `db:"build_parent"         json:"parent,omitempty"`
	Status       string            `db:"build_status"         json:"status"`
	Error        string            `db:"build_error"          json:"error,omitempty"`
	Event        string            `db:"build_event"          json:"event"`
	Action       string            `db:"build_action"         json:"action"`
	Link         string            `db:"build_link"           json:"link"`
	Timestamp    int64             `db:"build_timestamp"      json:"timestamp"`
	Title        string            `db:"build_title"          json:"title,omitempty"`
	Message      string            `db:"build_message"        json:"message"`
	Before       string            `db:"build_before"         json:"before"`
	After        string            `db:"build_after"          json:"after"`
	Ref          string            `db:"build_ref"            json:"ref"`
	Fork         string            `db:"build_source_repo"    json:"source_repo"`
	Source       string            `db:"build_source"         json:"source"`
	Target       string            `db:"build_target"         json:"target"`
	Author       string            `db:"build_author"         json:"author_login"`
	AuthorName   string            `db:"build_author_name"    json:"author_name"`
	AuthorEmail  string            `db:"build_author_email"   json:"author_email"`
	AuthorAvatar string            `db:"build_author_avatar"  json:"author_avatar"`
	Sender       string            `db:"build_sender"         json:"sender"`
	Params       map[string]string `db:"build_params"         json:"params,omitempty"`
	Cron         string            `db:"build_cron"           json:"cron,omitempty"`
	Deploy       string            `db:"build_deploy"         json:"deploy_to,omitempty"`
	DeployID     int64             `db:"build_deploy_id"      json:"deploy_id,omitempty"`
	Debug        bool              `db:"build_debug"          json:"debug,omitempty"`
	Started      int64             `db:"build_started"        json:"started"`
	Finished     int64             `db:"build_finished"       json:"finished"`
	Created      int64             `db:"build_created"        json:"created"`
	Updated      int64             `db:"build_updated"        json:"updated"`
	Version      int64             `db:"build_version"        json:"version"`
	Stages       []*Stage          `db:"-"                    json:"stages,omitempty"`
}
type Stage struct {
	ID        int64             `json:"id"`
	RepoID    int64             `json:"repo_id"`
	BuildID   int64             `json:"build_id"`
	Number    int               `json:"number"`
	Name      string            `json:"name"`
	Kind      string            `json:"kind,omitempty"`
	Type      string            `json:"type,omitempty"`
	Status    string            `json:"status"`
	Error     string            `json:"error,omitempty"`
	ErrIgnore bool              `json:"errignore"`
	ExitCode  int               `json:"exit_code"`
	Machine   string            `json:"machine,omitempty"`
	OS        string            `json:"os"`
	Arch      string            `json:"arch"`
	Variant   string            `json:"variant,omitempty"`
	Kernel    string            `json:"kernel,omitempty"`
	Limit     int               `json:"limit,omitempty"`
	LimitRepo int               `json:"throttle,omitempty"`
	Started   int64             `json:"started"`
	Stopped   int64             `json:"stopped"`
	Created   int64             `json:"created"`
	Updated   int64             `json:"updated"`
	Version   int64             `json:"version"`
	OnSuccess bool              `json:"on_success"`
	OnFailure bool              `json:"on_failure"`
	DependsOn []string          `json:"depends_on,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
	Steps     []*Step           `json:"steps,omitempty"`
}
type Step struct {
	ID        int64    `json:"id"`
	StageID   int64    `json:"step_id"` // this is a typo, fixing it has far reaching ramifications. It should only be attempted in a major version change
	Number    int      `json:"number"`
	Name      string   `json:"name"`
	Status    string   `json:"status"`
	Error     string   `json:"error,omitempty"`
	ErrIgnore bool     `json:"errignore,omitempty"`
	ExitCode  int      `json:"exit_code"`
	Started   int64    `json:"started,omitempty"`
	Stopped   int64    `json:"stopped,omitempty"`
	Version   int64    `json:"version"`
	DependsOn []string `json:"depends_on,omitempty"`
	Image     string   `json:"image,omitempty"`
	Detached  bool     `json:"detached,omitempty"`
	Schema    string   `json:"schema,omitempty"`
}

type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}
