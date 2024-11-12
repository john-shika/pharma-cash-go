package task

import (
	"errors"
	"fmt"
	"io"
	"nokowebapi/nokocore"
	"nokowebapi/xterm"
	"os"
	"strings"
	"sync"
	"time"
)

type DependsOnConfigParamsImpl interface {
	GetIterations() int
	GetDuration() time.Duration
}

type DependsOnConfigParams struct {
	Iterations int           `mapstructure:"iterations" json:"iterations" yaml:"iterations"`
	Duration   time.Duration `mapstructure:"duration" json:"duration" yaml:"duration"`
}

func NewDependsOnConfigParams(iterations int, duration time.Duration) *DependsOnConfigParams {
	return &DependsOnConfigParams{
		Iterations: iterations,
		Duration:   duration,
	}
}

func (d *DependsOnConfigParams) GetIterations() int {
	return d.Iterations
}

func (d *DependsOnConfigParams) GetDuration() time.Duration {
	return d.Duration
}

type DependsOnConfigImpl interface {
	GetTarget() string
	GetWaiter() string
	GetParams() *DependsOnConfigParams
}

type DependsOnConfig struct {
	Target string                 `mapstructure:"target" json:"target" yaml:"target"`
	Waiter string                 `mapstructure:"waiter" json:"waiter" yaml:"waiter"`
	Params *DependsOnConfigParams `mapstructure:"params" json:"params" yaml:"params"`
}

func NewDependsOnConfig(target string, waiter string) DependsOnConfigImpl {
	return &DependsOnConfig{
		Target: target,
		Waiter: waiter,
	}
}

func (d *DependsOnConfig) GetTarget() string {
	return d.Target
}

func (d *DependsOnConfig) GetWaiter() string {
	return d.Waiter
}

func (d *DependsOnConfig) GetParams() *DependsOnConfigParams {
	return d.Params
}

type Config struct {
	Name      string                  `mapstructure:"name" json:"name" yaml:"name"`
	Exec      string                  `mapstructure:"exec" json:"exec" yaml:"exec"`
	Args      []string                `mapstructure:"args" json:"args" yaml:"args"`
	Workdir   string                  `mapstructure:"workdir" json:"workdir" yaml:"workdir"`
	Environ   []string                `mapstructure:"environ" json:"environ" yaml:"environ"`
	Stdin     string                  `mapstructure:"stdin" json:"stdin" yaml:"stdin"`
	Stdout    string                  `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	Stderr    string                  `mapstructure:"stderr" json:"stderr" yaml:"stderr"`
	Network   *nokocore.NetworkConfig `mapstructure:"network" json:"network" yaml:"network"`
	DependsOn []*DependsOnConfig      `mapstructure:"depends_on" json:"dependsOn" yaml:"depends_on"`
}

func NewConfig() *Config {
	return &Config{}
}

func (t *Config) GetStdin() io.Reader {
	switch strings.ToLower(strings.TrimSpace(t.Stdin)) {
	case "console":
		return nokocore.NewSafeReader(xterm.Stdin)
	default:
		return nil
	}
}

func (t *Config) GetStdout() io.Writer {
	switch strings.ToLower(strings.TrimSpace(t.Stdout)) {
	case "console":
		return nokocore.NewSafeWriter(xterm.Stdout)
	default:
		return nil
	}
}

func (t *Config) GetStderr() io.Writer {
	switch strings.ToLower(strings.TrimSpace(t.Stderr)) {
	case "console":
		return nokocore.NewSafeWriter(xterm.Stderr)
	default:
		return nil
	}
}

type Tasks []*Config

func NewTasks() *Tasks {
	temp := make(Tasks, 0)
	return &temp
}

func (t *Tasks) GetTask(name string) *Config {
	for i, task := range *t {
		nokocore.KeepVoid(i)

		if strings.EqualFold(task.Name, name) {
			return task
		}
	}
	return nil
}

type DependsOnTaskImpl interface {
	GetTask() *Config
	GetWaiter() string
	GetParams() *DependsOnConfigParams
}

type DependsOnTask struct {
	Task   *Config
	Waiter string
	Params *DependsOnConfigParams
}

func NewDependsOnTask(task *Config, waiter string, params *DependsOnConfigParams) *DependsOnTask {
	return &DependsOnTask{
		Task:   task,
		Waiter: waiter,
		Params: params,
	}
}

func (d *DependsOnTask) GetTask() *Config {
	return d.Task
}

func (d *DependsOnTask) GetWaiter() string {
	return d.Waiter
}

func (d *DependsOnTask) GetParams() *DependsOnConfigParams {
	params := d.Params
	if params != nil {
		if params.Iterations == 0 {
			params.Iterations = 12
		}
		if params.Duration == 0 {
			params.Duration = time.Second
		}
	}
	return params
}

func (t *Tasks) GetDependsOnTask(task *Config) []*DependsOnTask {
	temp := make([]*DependsOnTask, 0)
	for i, dependsOn := range task.DependsOn {
		nokocore.KeepVoid(i)

		target := dependsOn.GetTarget()
		waiter := dependsOn.GetWaiter()
		params := dependsOn.GetParams()

		dependsTask := NewDependsOnTask(t.GetTask(target), waiter, params)
		temp = append(temp, dependsTask)
	}

	return temp
}

type ProcessTask struct {
	Process ProcessImpl
	Task    *Config
}

func NewProcessTask(process ProcessImpl, task *Config) *ProcessTask {
	return &ProcessTask{
		Process: process,
		Task:    task,
	}
}

func (p *ProcessTask) GetProcess() ProcessImpl {
	return p.Process
}

func (p *ProcessTask) GetTask() *Config {
	return p.Task
}

func (p *ProcessTask) IsRunning() bool {
	if p.Process != nil {
		return p.Process.IsRunning()
	}
	return false
}

func (p *ProcessTask) State() (StateImpl, error) {
	if p.Process != nil {
		return p.Process.State()
	}
	return nil, errors.New("process not started")
}

type MainTaskFunc func(*ProcessTasks, *Config) error

func (s MainTaskFunc) Call(processTasks *ProcessTasks, task *Config) error {
	return s(processTasks, task)
}

func runTask(processTasks *ProcessTasks, task *Config) error {
	var err error
	var workDir nokocore.WorkingDirImpl
	nokocore.KeepVoid(err, workDir)

	// try to dial url it-self
	if task.Network != nil {
		if nokocore.TryFetchUrl(task.Network.GetURL()) {
			return nil
		}
	}

	workFunc := func(workDir nokocore.WorkingDirImpl) error {
		if err = os.Chdir(task.Workdir); err != nil {
			return err
		}

		process := NewProcess(task.Exec, task.Args...)

		stdin := task.GetStdin()
		stdout := task.GetStdout()
		stderr := task.GetStderr()
		environ := task.Environ

		// binding stdin, stdout, stderr
		process.SetStdin(stdin)
		process.SetStdout(stdout)
		process.SetStderr(stderr)
		process.SetEnviron(environ)

		processTask := NewProcessTask(process, task)
		return processTasks.StartProcessTask(processTask)
	}

	if workDir, err = nokocore.SetWorkingDir(workFunc); err != nil {
		return err
	}

	return nil
}

var mainTask = func(p *ProcessTasks, t *Config) error {
	var err error
	var args []string
	var workDir nokocore.WorkingDirImpl
	nokocore.KeepVoid(err, args, workDir)

	// check if there are any command-line arguments provided.
	// if none, return an error indicating the source root directory cannot be determined.
	if len(os.Args) == 0 {
		return errors.New("can't get source root dir")
	}

	t.Exec = os.Args[0]
	t.Args = args

	return runTask(p, t)
}

type ProcessTasks struct {
	processes []*ProcessTask
	mainTask  MainTaskFunc
	locker    nokocore.LockerImpl
}

func NewProcessTasks(self MainTaskFunc) *ProcessTasks {
	if self == nil {
		self = mainTask
	}
	return &ProcessTasks{
		processes: make([]*ProcessTask, 0),
		mainTask:  self,
		locker:    nokocore.NewLocker(),
	}
}

func (p *ProcessTasks) RunSelf(task *Config) error {
	var err error
	nokocore.KeepVoid(err)

	task.Environ = append(task.Environ, "NOKOWEBAPI_SELF_RUNNING=1")
	if err = p.mainTask.Call(p, task); err != nil {
		return err
	}

	return nil
}

func (p *ProcessTasks) GetProcessTask(name string) *ProcessTask {
	for i, processTask := range p.processes {
		nokocore.KeepVoid(i)

		task := processTask.GetTask()
		if strings.EqualFold(task.Name, name) {
			return processTask
		}
	}
	return nil
}

func (p *ProcessTasks) StartProcessTask(processTask *ProcessTask) error {
	var err error
	nokocore.KeepVoid(err)

	p.locker.Lock(func() {
		task := processTask.GetTask()
		if p.GetProcessTask(task.Name) != nil {
			return
		}
		p.processes = append(p.processes, processTask)
		err = processTask.Process.Start()
	})
	return err
}

func (p *ProcessTasks) Wait() error {
	var err error
	var state StateImpl
	nokocore.KeepVoid(err, state)

	for i, process := range p.processes {
		nokocore.KeepVoid(i)

		if state, err = process.Process.Wait(); err != nil {
			fmt.Printf("[ERROR] Task '%s' failed.\n", process.Task.Name)
			return err
		}

		if state != nil {
			fmt.Printf("[STATE] PID: %d\n", state.Pid())
			fmt.Printf("[STATE] EXIT_CODE: %d\n", state.ExitCode())
		}
	}

	return nil
}

func (p *ProcessTasks) GetDependsOnProcessTask(task *Config) []*ProcessTask {
	temp := make([]*ProcessTask, 0)
	for i, dependsOn := range task.DependsOn {
		nokocore.KeepVoid(i)

		dependsTask := p.GetProcessTask(dependsOn.GetTarget())
		temp = append(temp, dependsTask)
	}

	return temp
}

func makeProcessFromTask(p *ProcessTasks, t *Config) error {
	var err error
	var workDir nokocore.WorkingDirImpl
	nokocore.KeepVoid(err, workDir)

	if p.GetProcessTask(t.Name) != nil {
		return nil
	}

	fmt.Printf("[RUN] Task '%s' started.\n", t.Name)

	if strings.EqualFold(t.Name, "self") {
		if err = p.RunSelf(t); err != nil {
			return err
		}
		return nil
	}

	return runTask(p, t)
}

func makeProcessFromTaskAsync(p *ProcessTasks, t *Config, err chan<- error) {
	err <- makeProcessFromTask(p, t)
}

func waitRun(tasks *Tasks, p *ProcessTasks, t *Config) error {
	var err error
	nokocore.KeepVoid(err)

	for i, dependsOnTask := range tasks.GetDependsOnTask(t) {
		nokocore.KeepVoid(i)

		target := dependsOnTask.GetTask()
		waiter := dependsOnTask.GetWaiter()
		params := dependsOnTask.GetParams()

		// Detect circular dependency between tasks
		if strings.EqualFold(target.Name, t.Name) {
			return errors.New("circular dependency detected")
		}

		if err = waitRun(tasks, p, target); err != nil {
			return err
		}

		switch waiter {
		case "none":
			break
		case "wait-for-http-alive":
			target.Network.WaitForHttpAlive(params.Iterations, params.Duration)
			break
		default:
			return errors.New("unknown waiter")
		}
	}

	if err = makeProcessFromTask(p, t); err != nil {
		fmt.Printf("[RUN] Task '%s' failed.\n", t.Name)
		return err
	}
	return nil
}

func waitRunTask(tasks *Tasks, p *ProcessTasks, t *Config) error {
	var ok bool
	var err error
	nokocore.KeepVoid(ok, err)

	dependsOnTasks := tasks.GetDependsOnTask(t)

	for i, dependsOnTask := range dependsOnTasks {
		nokocore.KeepVoid(i)

		target := dependsOnTask.GetTask()
		waiter := dependsOnTask.GetWaiter()
		params := dependsOnTask.GetParams()

		// detect circular dependency between tasks
		if strings.EqualFold(target.Name, t.Name) {
			return errors.New("circular dependency detected")
		}

		if err = waitRunTask(tasks, p, target); err != nil {
			return err
		}

		switch waiter {
		case "none":
			break
		case "wait-for-http-alive":
			target.Network.WaitForHttpAlive(params.Iterations, params.Duration)
			break
		default:
			return errors.New("unknown waiter")
		}
	}

	errorStack := make(chan error)
	defer close(errorStack)

	go makeProcessFromTaskAsync(p, t, errorStack)

	for {
		select {
		case err, ok = <-errorStack:
			if !ok {
				fmt.Println("[WARN] Channel closed.")
				return nil
			}
			if err != nil {
				fmt.Printf("[RUN] Task '%s' failed.\n", t.Name)
				return err
			}
			return nil
		}
	}
}

func (t *Tasks) Execute(processTasks *ProcessTasks) error {
	var err error
	nokocore.KeepVoid(err)

	tasks := *t
	for i, task := range tasks {
		nokocore.KeepVoid(i)
		if err = waitRun(t, processTasks, task); err != nil {
			return err
		}
	}

	return nil
}

type WaitTasks struct {
	*sync.WaitGroup
	err error
}

func NewWaitTasks() *WaitTasks {
	return &WaitTasks{
		WaitGroup: &sync.WaitGroup{},
		err:       nil,
	}
}

func (t *WaitTasks) Wait() error {
	t.WaitGroup.Wait()
	return t.err
}

func (t *WaitTasks) Add(delta int) {
	t.WaitGroup.Add(delta)
}

func (t *WaitTasks) Run(action nokocore.ActionReturn[error]) {
	defer t.WaitGroup.Done()
	t.err = action()
}

func (t *Tasks) ExecuteAsync(processTasks *ProcessTasks) *WaitTasks {
	var err error
	nokocore.KeepVoid(err)

	tasks := *t
	size := len(tasks)

	wt := NewWaitTasks()
	wt.Add(size)

	for i, task := range tasks {
		nokocore.KeepVoid(i)

		// no need goroutine's for a wait run task, already run in goroutine.
		wt.Run(func() error {
			return waitRunTask(t, processTasks, task)
		})
	}

	return wt
}
