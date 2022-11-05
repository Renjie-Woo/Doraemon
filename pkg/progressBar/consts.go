package progressBar

const (
	defaultTitle                         = "Progress"
	defaultGraph                         = "█"
	currentCountGreaterThanTotalOneError = "current count: %v is illegally greater than the total one: %v, please check it again."
	invalidTotalCountError               = "total count must greater than current count"
	invalidCurrentCountError             = "current count must no less than 0"
)
