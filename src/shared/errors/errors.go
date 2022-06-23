package errors

const (
	CATEGORY_NOT_FOUND              = "category not found"
	EMAIL_ALREADY_EXISTS            = "email already registered"
	FIELD_MUST_BE_INTEGER           = "the [field] field must be integer"
	FIELD_MUST_BE_GREATER_THAN_ZERO = "the [field] field must be greater than zero"
	GIFT_NOT_FOUND                  = "gift not found"
	GIFT_NOT_AVAILABLE              = "gift not available"
	GIFT_NOT_BELONG_CATEGORY        = "gift does not belong to this category"
	GIFT_QUANTITY_NOT_AVAILABLE     = "quantity greater than quantity available"
	INCORRECT_PASSWORD              = "incorrect password"
	INVALID_EMAIL                   = "invalid email"
	INVALID_PASSWORD                = "invalid password"
	NAME_IS_EMPTY                   = "name is empty"
	UNAUTHORIZED                    = "unauthorized"
)
