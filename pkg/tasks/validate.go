package tasks

import (
	"errors"
	"reflect"
)

var (
	// ErrTaskMustBeFunc ...
	ErrTaskMustBeFunc = errors.New("Task must be a func type")
	// ErrTaskReturnsNoValue ...
	ErrTaskReturnsNoValue = errors.New("Tasks must return at least a single value")
	// ErrLastReturnValueMustBeError ..
	ErrLastReturnValueMustBeError = errors.New("Last return value of a task must be error")
)

// ValidateTask validates task function using reflection and makes sure
// it has a proper signature. Functions used as tasks must return at least a
// single value and the Last return type must be error
func ValidateTask(task interface{}) error {
	v := reflect.ValueOf(task)
	t := v.Type()

	// Task must be a function
	if t.Kind() != reflect.Func {
		return ErrTaskMustBeFunc
	}

	// Task must return at least a single value
	if t.NumOut() < 1 {
		return ErrTaskReturnsNoValue
	}

	// Last return value must be error
	LastReturnType := t.Out(t.NumOut() - 1)
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	if !LastReturnType.Implements(errorInterface) {
		return ErrLastReturnValueMustBeError
	}

	return nil
}
