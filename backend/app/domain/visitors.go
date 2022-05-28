package domain

type Visitors struct {
	yesterdayVisitor int
	todayVisitor     int
	sumVisitor       int
}

func NewVisitors() *Visitors {
	return &Visitors{}
}

func (v *Visitors) YesterdayVisitors() int {
	return v.yesterdayVisitor
}

func (v *Visitors) TodayVisitors() int {
	return v.todayVisitor
}

func (v *Visitors) SumVisitor() int {
	return v.sumVisitor
}
