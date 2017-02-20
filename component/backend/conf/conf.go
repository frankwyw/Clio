package conf

var (
	Debug *bool

	Backend_Host string
	Backend_Port string

	K8sapihost  string
	K8sapiport  string
	K8sprotocol string

	UserMhost string
	UserMport string

	REGSERVERHOST string
	REGSERVERPORT string
)

type BackErr struct {
	Str string
}

func (b BackErr) Error() string {
	return b.Str
}
