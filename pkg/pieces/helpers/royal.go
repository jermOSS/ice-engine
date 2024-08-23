package helpers

type Royal struct{}

func (self *Royal) IsRoyal() bool {
	return true
}

type NonRoyal struct{}

func (self *NonRoyal) IsRoyal() bool {
	return false
}
