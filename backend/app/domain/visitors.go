package domain

import "time"

type Visitors struct {
	ID               int       `gorm:"primaryKey;column:id"`
	YesterdayVisitor *int      `gorm:"column:yesterday_visitor"`
	TodayVisitor     *int      `gorm:"column:today_visitor"`
	VisitorSum       *int      `gorm:"column:visitor_sum"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

func NewVisitors(yesterday, today, sum int) *Visitors {
	return &Visitors{
		YesterdayVisitor: &yesterday,
		TodayVisitor: &today,
		VisitorSum: &sum,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (v *Visitors) GetYesterdayVisitors() int {
	return *v.YesterdayVisitor
}

func (v *Visitors) GetTodayVisitors() int {
	return *v.TodayVisitor
}

func (v *Visitors) GetVisitorSum() int {
	return *v.VisitorSum
}

func (v *Visitors) CoutupTodayVisitors() {
	*v.TodayVisitor += 1
}

func (v *Visitors) CountupSumVisitor() {
	*v.VisitorSum += 1
}

func (v *Visitors) ResetTodayVisitors(n int) {
	v.TodayVisitor = &n
}

func (v *Visitors) SetYesterdayVisitors(n int) {
	v.YesterdayVisitor = &n
}
