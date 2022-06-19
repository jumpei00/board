package domain

import "time"

type Visitor struct {
	ID               int       `gorm:"primaryKey;column:id"`
	YesterdayVisitor *int      `gorm:"column:yesterday_visitor"`
	TodayVisitor     *int      `gorm:"column:today_visitor"`
	VisitorSum       *int      `gorm:"column:visitor_sum"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

func NewVisitor(yesterday, today, sum int) *Visitor {
	return &Visitor{
		YesterdayVisitor: &yesterday,
		TodayVisitor: &today,
		VisitorSum: &sum,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (v *Visitor) GetYesterdayVisitor() int {
	return *v.YesterdayVisitor
}

func (v *Visitor) GetTodayVisitor() int {
	return *v.TodayVisitor
}

func (v *Visitor) GetVisitorSum() int {
	return *v.VisitorSum
}

func (v *Visitor) CoutupTodayVisitor() {
	*v.TodayVisitor += 1
}

func (v *Visitor) CountupSumVisitor() {
	*v.VisitorSum += 1
}

func (v *Visitor) ResetTodayVisitor(n int) {
	v.TodayVisitor = &n
}

func (v *Visitor) SetYesterdayVisitor(n int) {
	v.YesterdayVisitor = &n
}
