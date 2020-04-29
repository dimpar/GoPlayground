package processes

// NewCommandInput wraps parameters for a unix process execution.
// It is required to provide a command name. Other params are optional.
func NewCommandInput(name string) CommandInput {
	return &commandInput{
		name: name,
		arguments: []string{},
		environmentVariables: []string{},
		isSetpgid: false,
		numberOfRetries: 0,
	}
}

type commandInput struct {
	name string
	arguments []string
	environmentVariables []string
	isSetpgid bool
	numberOfRetries uint
}

func(ci *commandInput) Name() string {
	return ci.name;
}

func(ci *commandInput) Arguments() []string {
	return ci.arguments;
}

func(ci *commandInput) SetArguments(arguments []string) {
	ci.arguments = arguments
}

func(ci *commandInput) EnvironmentVariables() []string {
	return ci.environmentVariables
}

func(ci *commandInput) SetEnvironmentVariables(env []string) {
	ci.environmentVariables = env
}

func(ci *commandInput) Pgid() bool {
	return ci.isSetpgid;
}

func(ci *commandInput) SetPgid(isSet bool) {
	ci.isSetpgid = isSet
}

// Number of retries for process execution upon failure.
func(ci *commandInput) NumberOfRetries() uint {
	return ci.numberOfRetries;
}

func(ci *commandInput) SetNumberOfRetries(retries uint) {
	ci.numberOfRetries = retries
}
