package configuration

// Server is the configuration object for the server mode.
type Server struct {
	Config      `yaml:",inline"`
	LookupdAddr string `yaml:"nsq_lookupd"`
	Channel     string `yaml:"nsq_channel"`

	// Repositories maps GitHub repositories full names their corresponding
	// NSQ topic.
	Repositories map[string]string `yaml:"repositories"`

	// CommonActions defines the triggers and operations which apply to every configured repository.
	CommonActions []Action `yaml:"common_configuration"`
}

// Action is the definition of an action: it descrbibes operations to apply when any of the
// associated triggers are met.
type Action struct {
	// Triggers is the collection of GitHub events that should trigger the action. The keys must be
	// valid GitHub event types (e.g., "pull_request"), and the value must be a list of alid values
	// for the action field of the GitHub paylost (e.g., "created").
	Triggers Trigger `yaml:"triggers"`

	// Operations to apply to all repositories when any trigger is met.
	Operations []OperationConfiguration `yaml:"operations"`
}

// StringSlice is a slice of strings.
type StringSlice []string

// Contains returns whether the StringSlice contains a given item.
func (s StringSlice) Contains(item string) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

// Trigger associates a GitHub event type (e.g., "issues", or "pull request") with a collection of
// corresponding actions (e.g., [ "opened", "reopened" ]).
type Trigger map[string]StringSlice

// Contains returns whether the triggers contains the specified (event, action) tuple.
func (t Trigger) Contains(githubEvent, githubAction string) bool {
	if actions, ok := t[githubEvent]; ok {
		return actions.Contains(githubAction)
	}
	return false
}
