package model

type InvalidQualityError struct{}

func (e InvalidQualityError) Error() string {
    return "invalid quality identifier"
}
