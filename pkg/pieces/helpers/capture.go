package helpers

type Capturable struct{}

func (self *Capturable) CanBeCaptured() bool {
	return true
}

type Uncapturable struct{}

func (self *Uncapturable) CanBeCaptured() bool {
	return false
}
