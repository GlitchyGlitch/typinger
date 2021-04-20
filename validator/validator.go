package validator

import (
	"context"
	"reflect"

	"github.com/GlitchyGlitch/typinger/errs"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Validator struct {
	*validator.Validate
}

func New() *Validator {
	validate := validator.New()
	v := &Validator{
		validate,
	}
	return v
}

func (v Validator) CheckStruct(ctx context.Context, s interface{}, canBeNil bool) bool {
	if s == nil || (reflect.ValueOf(s).Kind() == reflect.Ptr && reflect.ValueOf(s).IsNil()) {
		return canBeNil
	}
	err := v.Struct(s)
	if err != nil {
		v.AddErrs(ctx, err)
		return false
	}
	return true
}

func (v Validator) AddErrs(ctx context.Context, err error) {
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errs.Add(ctx, errs.Validation(ctx, e.Field()))
		}
	}
}

func (v Validator) CheckUUID(ctx context.Context, u string) bool {
	_, err := uuid.Parse(u)
	if err != nil {
		errs.Add(ctx, errs.Validation(ctx, "id"))
		return false
	}
	return true
}

func (v Validator) CheckPagination(ctx context.Context, first, offset *int) bool {
	if *first < 0 {
		errs.Add(ctx, errs.Validation(ctx, "first"))
		return false
	}
	if *offset < 0 {
		errs.Add(ctx, errs.Validation(ctx, "offset"))
		return false
	}
	return true
}
