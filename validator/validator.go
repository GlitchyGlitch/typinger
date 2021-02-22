package validator

import (
	"context"

	"github.com/GlitchyGlitch/typinger/errs"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator"
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

func (v Validator) ValidateErrs(ctx context.Context, s interface{}) bool {
	err := v.Struct(s)
	if err != nil {
		v.AddErrs(ctx, err)
		return false
	}
	return true
}

func (v *Validator) AddErrs(ctx context.Context, err error) {
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			graphql.AddError(ctx, errs.Validation(ctx, e.Field()))
		}
	}
}
