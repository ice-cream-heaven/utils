package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/ice-cream-heaven/log"
	"github.com/ice-cream-heaven/utils/common"
)

var validate = validator.New()

func Validate(m interface{}) (err error) {
	err = validate.Struct(m)
	if err != nil {
		log.Errorf("err:%v", err)
		return common.NewErrorWithError(err)
	}
	return nil
}
