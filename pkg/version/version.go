package version

var Latest = New("divisions", "develop", nil)

type Version interface {
	Application() string
	CommitHash() *string
	Tag() string
}

type version struct {
	application string
	commitHash  *string
	tag         string
}

func New(application, tag string, commitHash *string) Version {
	return &version{
		application: application,
		commitHash:  commitHash,
		tag:         tag,
	}
}

func (v version) Application() string {
	return v.application
}

func (v version) CommitHash() *string {
	return v.commitHash
}

func (v version) Tag() string {
	return v.tag
}
