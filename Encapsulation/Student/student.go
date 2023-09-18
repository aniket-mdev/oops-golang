package student

type Student struct {
	name   string
	course string
	RollNo int
}

// Setter Methods
// setter method only for name and course becouse they can not access directly from other packages
func (s *Student) SetName(name string) {
	s.name = name
}
func (s *Student) SetCourse(course string) {
	s.course = course
}

// Getter methods
func (s *Student) GetName() string {
	return s.name
}
func (s *Student) GetCourse() string {
	return s.course
}
func (s *Student) GetRollNo() int {
	return s.RollNo
}
