package ussd

import "github.com/alainmucyo/ussd-go/validator"

// Form is a USSD form.
type Form struct {
	ValidationMessage  string
	Route              route
	ProcessingPosition int
	Data               FormData
	Inputs             []Input
}

// NewForm creates a new form.
func NewForm() *Form {
	return &Form{
		Data:   make(FormData),
		Inputs: make([]Input, 0),
	}
}

// Input adds an Input to USSD form.
func (f *Form) Input(name, displayName string,
	options ...Option) *Form {
	input := newInput(StrTrim(name), StrTrim(displayName))
	for _, option := range options {
		input.Options = append(input.Options, option)
	}
	f.Inputs = append(f.Inputs, input)
	return f
}

// Validate Input. See validator.Map for available validators.
func (f *Form) Validate(validatorKey string, args ...string) *Form {
	validatorKey = StrTrim(StrLower(validatorKey))
	if _, ok := validator.Map[validatorKey]; !ok {
		panic(&validatorDoesNotExistError{validatorKey})
	}
	i := len(f.Inputs) - 1
	input := f.Inputs[i]
	input.Validators = append(input.Validators, validatorData{
		Key:  validatorKey,
		Args: args,
	})
	f.Inputs[i] = input
	return f
}

// Option creates a USSD Input Option.
func (f Form) Option(value, displayValue string) Option {
	return Option{
		Value: StrTrim(value), DisplayValue: StrTrim(displayValue),
	}
}

type Input struct {
	Name, DisplayName string
	Options           []Option
	Validators        []validatorData
}

type Option struct {
	Value, DisplayValue string
}

type validatorData struct {
	Key  string
	Args []string
}

func newInput(name, displayName string) Input {
	return Input{
		Name:        name,
		DisplayName: displayName,
		Options:     make([]Option, 0),
		Validators:  make([]validatorData, 0),
	}
}

func (i Input) hasOptions() bool {
	return len(i.Options) > 0
}
