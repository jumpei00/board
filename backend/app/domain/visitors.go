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

func (v *Visitors) CoutupTodayVisitors() {
	v.todayVisitor += 1
}

func (v *Visitors) CountupSumVisitor() {
	v.sumVisitor += 1
}

func (v *Visitors) ResetTodayVisitors(n int) {
	v.todayVisitor = n
}

func (v *Visitors) SetYesterdayVisitors(n int) {
	v.yesterdayVisitor = n
}
