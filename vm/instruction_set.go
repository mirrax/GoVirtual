//	TODO:	bytecode optimisation
//	TODO:	JIT compilation
//	TODO:	AOT compilation

package vm

import "container/vector"

type OpCode struct {
	code, a, b, c	int;
}
func (o *OpCode) Similar(p *OpCode) bool {
	return o.code == p.code
}
func (o *OpCode) Identical(p *OpCode) bool {
	return o.Similar(p) && o.a == p.a && o.b == p.b && o.c == p.c
}
func (o *OpCode) Replace(p *OpCode) {
	o.code, o.a, o.b, o.c = p.code, p.a, p.b, p.c
}

type InstructionSet struct {
	ops				vector.Vector
	tokens			map[string]int
}
func (i *InstructionSet) Init() {
	i.tokens = make(map[string]int)
}
func (i *InstructionSet) Len() int {
	return i.ops.Len()
}
func (i *InstructionSet) Define(name string, closure func (o *OpCode)) bool {
	// Ensure instruction token hasn't yet been defined
	if _, ok := i.tokens[name]; !ok {
		i.ops.Push(closure)
		i.tokens[name] = i.ops.Len() - 1
		return true
	}
	return false
}
func (i *InstructionSet) Find(name string) int {
	if op, ok := i.tokens[name]; ok {
		return op
	}
	return -1
}
func (i *InstructionSet) Invoke(o *OpCode) bool {
	if o.code < 0 || o.code >= i.ops.Len() { return false }
	i.ops.At(o.code).(func (o *OpCode))(o)
	return true
}
