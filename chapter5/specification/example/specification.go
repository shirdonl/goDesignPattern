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

package example

// 对象
type Object struct {
	Attribute int
}

// 规格接口
type Specification interface {
	IsSatisfiedBy(Object) bool
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
func (cs *CompositeSpecification) IsSatisfiedBy(obj Object) bool {
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
func (as *AndSpecification) IsSatisfiedBy(obj Object) bool {
	return as.Specification.IsSatisfiedBy(obj) && as.compare.IsSatisfiedBy(obj)
}

// 或规格
type OrSpecification struct {
	Specification
	compare Specification
}

// 检查规格
func (os *OrSpecification) IsSatisfiedBy(obj Object) bool {
	return os.Specification.IsSatisfiedBy(obj) || os.compare.IsSatisfiedBy(obj)
}

// 非规格
type NotSpecification struct {
	Specification
}

// 检查规格
func (ns *NotSpecification) IsSatisfiedBy(obj Object) bool {
	return ns.Specification.IsSatisfiedBy(obj)
}

// 业务规格
type BusinessSpecification struct {
	Specification
}

// 检查规格
func (bs *BusinessSpecification) IsSatisfiedBy(obj Object) bool {
	return obj.Attribute >= 8
}

// 构造函数
func NewBusinessSpecification() Specification {
	a := &BusinessSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}
