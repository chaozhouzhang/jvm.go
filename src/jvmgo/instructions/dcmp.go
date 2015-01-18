package instructions

import "jvmgo/rtda"

// Compare double
type dcmpg struct {}
func (self *dcmpg) fetchOperands(bcr *BytecodeReader) {}
func (self *dcmpg) execute(thread *rtda.Thread) {
    dcmp(thread, true)
}

type dcmpl struct {}
func (self *dcmpl) fetchOperands(bcr *BytecodeReader) {}
func (self *dcmpl) execute(thread *rtda.Thread) {
    dcmp(thread, false)
}

func dcmp(thread *rtda.Thread, gFlag bool) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    if v1 > v2 {
        stack.PushInt(1)
    } else if v1 == v2 {
        stack.PushInt(0)
    } else if v1 < v2 {
        stack.PushInt(-1)
    } else if gFlag {
        stack.PushInt(1)
    } else {
        stack.PushInt(-1)
    }
}
