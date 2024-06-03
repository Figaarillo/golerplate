package validation

import "fmt"

type Rule func(key string, value interface{}) error

type Validator struct {
	rules []Rule
}

func (v *Validator) Add(rule Rule) {
	v.rules = append(v.rules, rule)
}

func (v *Validator) Validate(data map[string]interface{}) []error {
	var errors []error

	for _, rule := range v.rules {
		for key, value := range data {
			if err := rule(key, value); err != nil {
				errors = append(errors, err)
			}
		}
	}

	return errors
}

func ValdatePrecense(key string, value interface{}) error {
	if value == "" {
		return fmt.Errorf("%s is required", key)
	}

	return nil
}

func ValdateEmail(key string, value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string", key)
	}

	if str == "" || str[:1] == "@" || str[len(str)-1:] == "@" {
		return fmt.Errorf("%s must be a valid email", key)
	}

	return nil
}
