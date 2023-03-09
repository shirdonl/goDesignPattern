//++++++++++++++++++++++++++++++++++++++++
// 《Go语言设计模式》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 作者知乎：https://www.zhihu.com/people/shirdonl
// 仓库地址：https://gitee.com/shirdonl/goDesignPattern
// 仓库地址：https://github.com/shirdonl/goDesignPattern
// 交流咨询，请关注公众号"源码大数据"
//++++++++++++++++++++++++++++++++++++++++

package actualCombat

// 发票数据对象
type Invoice struct {
	Day    int
	Notice int
	IsSent bool
}

// 数据规格接口
type Specification interface {
	IsSatisfiedBy(Invoice) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

// 组合规格
type CompositeSpecification struct {
	Specification
}

// 检查规格
func (cs *CompositeSpecification) IsSatisfiedBy(in Invoice) bool {
	return false
}

// 规格与操作
func (cs *CompositeSpecification) And(spec Specification) Specification {
	a := &AndSpecification{
		cs.Specification, spec,
	}
	a.Relate(a)
	return a
}

// 规格或操作
func (cs *CompositeSpecification) Or(spec Specification) Specification {
	a := &OrSpecification{
		cs.Specification, spec,
	}
	a.Relate(a)
	return a
}

// 规格非操作
func (cs *CompositeSpecification) Not() Specification {
	a := &NotSpecification{
		cs.Specification,
	}
	a.Relate(a)
	return a
}

// 与规格有关
func (cs *CompositeSpecification) Relate(spec Specification) {
	cs.Specification = spec
}

// 与规格
type AndSpecification struct {
	Specification
	compare Specification
}

// 检查规格
func (as *AndSpecification) IsSatisfiedBy(in Invoice) bool {
	return as.Specification.IsSatisfiedBy(in) && as.compare.IsSatisfiedBy(in)
}

// 或规格
type OrSpecification struct {
	Specification
	compare Specification
}

// 检查规格
func (os *OrSpecification) IsSatisfiedBy(in Invoice) bool {
	return os.Specification.IsSatisfiedBy(in) || os.compare.IsSatisfiedBy(in)
}

// 非规格
type NotSpecification struct {
	Specification
}

// 检查规格
func (ns *NotSpecification) IsSatisfiedBy(in Invoice) bool {
	return ns.Specification.IsSatisfiedBy(in)
}

// 数据到期规格
type OverDueSpecification struct {
	Specification
}

// 检查规格
func (os *OverDueSpecification) IsSatisfiedBy(in Invoice) bool {
	return in.Day >= 30
}

// 创建数据到期规格
func NewOverDueSpecification() Specification {
	a := &OverDueSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}

// 通知发送规格
type NoticeSentSpecification struct {
	Specification
}

// 检查规格
func (ns *NoticeSentSpecification) IsSatisfiedBy(in Invoice) bool {
	return in.Notice >= 3
}

// 创建通知发送规格
func NewNoticeSentSpecification() Specification {
	a := &NoticeSentSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}

// 是否收到发票通知规格
type InCollectionSpecification struct {
	Specification
}

// 检查规格
func (ics *InCollectionSpecification) IsSatisfiedBy(in Invoice) bool {
	return !in.IsSent
}

// 创建是否收到发票通知规格
func NewInCollectionSpecification() Specification {
	a := &InCollectionSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}
