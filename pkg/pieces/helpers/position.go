package helpers

import "github.com/jermOSS/ice-engine/pkg/defs"

type Position struct {
	X int64
	Y int64
}

func (self *Position) Move(x int64, y int64) {
	self.X = x
	self.Y = y
}

func (self *Position) GetPosition() defs.Position {
	return defs.NewPosition(self.X, self.Y)
}

func (self *Position) GetHasMoved() bool {
	return true
}

func (self *Position) GetX() int64 {
	return self.X
}

func (self *Position) GetY() int64 {
	return self.Y
}

func (self *Position) GetI() int64 {
	return self.GetPosition().GetI()
}

func (self *Position) GetJ() int64 {
	return self.GetPosition().GetJ()
}

type PositionHasMoved struct {
	X        int64
	Y        int64
	HasMoved bool
}

func (self *PositionHasMoved) Move(x int64, y int64) {
	self.HasMoved = true
	self.X = x
	self.Y = y
}

func (self *PositionHasMoved) GetPosition() defs.Position {
	return defs.NewPosition(self.X, self.Y)
}

func (self *PositionHasMoved) GetX() int64 {
	return self.X
}

func (self *PositionHasMoved) GetY() int64 {
	return self.Y
}

func (self *PositionHasMoved) GetI() int64 {
	return self.GetPosition().GetI()
}

func (self *PositionHasMoved) GetJ() int64 {
	return self.GetPosition().GetJ()
}

func (self *PositionHasMoved) GetHasMoved() bool {
	return self.HasMoved
}
