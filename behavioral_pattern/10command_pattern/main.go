package main

import "fmt"

type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}
func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

type Command interface {
	execute()
}

// 对docker进行封装
type CommandTreatEye struct {
	docker *Doctor
}

func (cmd *CommandTreatEye) execute() {
	cmd.docker.treatNose()
}

// 对docker进行封装
type CommandTreatNose struct {
	docker *Doctor
}

func (cmd *CommandTreatNose) execute() {
	cmd.docker.treatEye()
}

type Nursenvoker struct {
	CmdList []Command
}

func (n *Nursenvoker) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.execute() //执行病单绑定的命令(这里会调用病单已经绑定的医生的诊断方法)
	}
}

//病人
func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor)
	//治疗眼睛的病单
	cmdEye := CommandTreatEye{doctor}
	//治疗鼻子的病单
	cmdNose := CommandTreatNose{doctor}

	//护士
	nurse := new(Nursenvoker)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	//执行病单指令
	nurse.Notify()
}
