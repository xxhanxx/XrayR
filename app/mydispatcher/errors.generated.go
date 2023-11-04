package mydispatcher

import "github.com/xxhanxx/Xray-core/common/errors"

func newError(values ...interface{}) *errors.Error {
	return errors.New(values...)
}
