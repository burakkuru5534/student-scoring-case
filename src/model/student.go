package model

type Student struct {
	ID        int64  `db:"id"`
	GroupName string `db:"group_name"`
	Number    int64  `db:"number"`
	Point     int64  `db:"point"`
}

func (s *Student) GetGroupName() string {
	return s.GroupName
}

func (s *Student) GetNumber() int64 {
	return s.Number
}

func (s *Student) GetPoint() int64 {
	return s.Point
}

func (s *Student) SetGroupName(groupName string) {
	s.GroupName = groupName
}
