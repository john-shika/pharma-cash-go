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

const NokoWebApiAutoRunEnv string = "NOKOWEBAPI_SELF_RUNNING"

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

// DependsOnTaskConfig struct, T any can be string or *Config itself.
// Keep in mind, parsing by viper config file, declare T must be string.
type DependsOnTaskConfig[T any] struct {
	Target T                      `mapstructure:"target" json:"target" yaml:"target"`
	Waiter string                 `mapstructure:"waiter" json:"waiter" yaml:"waiter"`
	Params *DependsOnConfigParams `mapstructure:"params" json:"params" yaml:"params"`
}

func NewDependsOnTaskConfig[T any](target T, waiter string, params *DependsOnConfigParams) *DependsOnTaskConfig[T] {
	return &DependsOnTaskConfig[T]{
		Target: target,
		Waiter: waiter,
		Params: params,
	}
}

func (d *DependsOnTaskConfig[T]) GetTarget() T {
	return d.Target
}

func (d *DependsOnTaskConfig[T]) GetWaiter() string {
	return d.Waiter
}

func (d *DependsOnTaskConfig[T]) GetParams() *DependsOnConfigParams {
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

type Config struct {
	Name      string                         `mapstructure:"name" json:"name" yaml:"name"`
	Exec      string                         `mapstructure:"exec" json:"exec" yaml:"exec"`
	Args      []string                       `mapstructure:"args" json:"args" yaml:"args"`
	Workdir   string                         `mapstructure:"workdir" json:"workdir" yaml:"workdir"`
	Environ   []string                       `mapstructure:"environ" json:"environ" yaml:"environ"`
	Stdin     string                         `mapstructure:"stdin" json:"stdin" yaml:"stdin"`
	Stdout    string                         `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	Stderr    string                         `mapstructure:"stderr" json:"stderr" yaml:"stderr"`
	Network   *nokocore.NetworkConfig        `mapstructure:"network" json:"network" yaml:"network"`
	DependsOn []*DependsOnTaskConfig[string] `mapstructure:"depends_on" json:"dependsOn" yaml:"depends_on"`
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

// TasksConfig struct, keep it mind, parsing by viper config file
type TasksConfig []*Config

func NewTasksConfig() *TasksConfig {
	temp := make(TasksConfig, 0)
	return &temp
}

func (t *TasksConfig) GetNameType() string {
	return "Tasks"
}

func (t *TasksConfig) GetTaskConfig(name string) *Config {
	for i, task := range *t {
		nokocore.KeepVoid(i)

		taskName := strings.TrimSpace(task.Name)
		if strings.EqualFold(taskName, name) {
			return task
		}
	}
	return nil
}

func (t *TasksConfig) GetDependsOnTaskConfig(task *Config) []*DependsOnTaskConfig[*Config] {
	temp := make([]*DependsOnTaskConfig[*Config], 0)
	for i, dependsOn := range task.DependsOn {
		nokocore.KeepVoid(i)

		target := dependsOn.GetTarget()
		waiter := dependsOn.GetWaiter()
		params := dependsOn.GetParams()

		taskConfig := t.GetTaskConfig(target)
		dependsTask := NewDependsOnTaskConfig(taskConfig, waiter, params)
		temp = append(temp, dependsTask)
	}

	return temp
}

type ProcessTaskImpl interface {
	GetProcess() ProcessImpl
	GetTaskConfig() *Config
	IsRunning() bool
}

type ProcessTask struct {
	TaskConfig *Config
	Process    ProcessImpl
}

func NewProcessTask(process ProcessImpl, config *Config) ProcessTaskImpl {
	return &ProcessTask{
		TaskConfig: config,
		Process:    process,
	}
}

func (p *ProcessTask) GetProcess() ProcessImpl {
	return p.Process
}

func (p *ProcessTask) GetTaskConfig() *Config {
	return p.TaskConfig
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

func runTask(pTasks ProcessTasksImpl, task *Config) error {
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

		pTask := NewProcessTask(process, task)
		return pTasks.StartProcessTask(pTask)
	}

	if workDir, err = nokocore.SetWorkingDir(workFunc); err != nil {
		return err
	}

	return nil
}

var mainTask = func(pTasks ProcessTasksImpl, task *Config) error {
	var err error
	var args []string
	var exec string
	nokocore.KeepVoid(err, args, exec)

	// check if there are any command-line arguments provided.
	// if none, return an error indicating the source root directory cannot be determined.
	if len(os.Args) == 0 {
		return errors.New("can't get source root dir")
	}

	// initial execute and arguments it-self
	if exec, err = os.Executable(); err != nil {
		return err
	}

	// binding values
	task.Exec = exec
	task.Args = args

	nokoWebApiAutoRunEnv := fmt.Sprintf("%s=%s", NokoWebApiAutoRunEnv, "1")

	// set value for NOKOWEBAPI_SELF_RUNNING env
	task.Environ = append(task.Environ, nokoWebApiAutoRunEnv)
	return runTask(pTasks, task)
}

func EntryPoint(self nokocore.MainFunc, pTasksHandler nokocore.ActionSingleParam[ProcessTasksImpl]) {
	ApplyMainSelf(self)
	pTasks := NewProcessTasks()
	pTasksHandler.Call(pTasks)
}

func ApplyMainSelf(self nokocore.MainFunc) {
	var ok bool
	var nokoWebApiAutoRunEnv string
	nokocore.KeepVoid(ok, nokoWebApiAutoRunEnv)

	if nokoWebApiAutoRunEnv, ok = os.LookupEnv(NokoWebApiAutoRunEnv); ok {
		if nokocore.ParseEnvToBool(nokoWebApiAutoRunEnv) {

			// will be exited
			nokocore.ApplyMainFunc(self)
			return
		}
	}

	nokocore.NoErr(os.Setenv(NokoWebApiAutoRunEnv, "1"))
}

type ProcessTasksImpl interface {
	GetProcessTask(name string) ProcessTaskImpl
	GetDependsOnProcessTask(task *Config) []ProcessTaskImpl
	StartProcessTask(pTask ProcessTaskImpl) error
	ExecuteAsync(tasksConfig *TasksConfig)
	Execute(tasksConfig *TasksConfig) error
	Wait() error
}

type ProcessTasks struct {
	mainTask nokocore.ActionDoubleParamsReturn[ProcessTasksImpl, *Config, error]
	pTasks   []ProcessTaskImpl
	locker   nokocore.LockerImpl
	regis    WaitTasksImpl
}

func NewProcessTasks() ProcessTasksImpl {
	return &ProcessTasks{
		mainTask: mainTask,
		pTasks:   make([]ProcessTaskImpl, 0),
		locker:   nokocore.NewLocker(),
	}
}

func (p *ProcessTasks) applyMainTask(task *Config) error {
	var err error
	nokocore.KeepVoid(err)

	if err = p.mainTask.Call(p, task); err != nil {
		return err
	}

	return nil
}

func (p *ProcessTasks) GetProcessTask(name string) ProcessTaskImpl {
	for i, pTask := range p.pTasks {
		nokocore.KeepVoid(i)

		task := pTask.GetTaskConfig()
		taskName := strings.TrimSpace(task.Name)

		if strings.EqualFold(taskName, name) {
			return pTask
		}
	}
	return nil
}

func (p *ProcessTasks) GetDependsOnProcessTask(task *Config) []ProcessTaskImpl {
	temp := make([]ProcessTaskImpl, 0)
	for i, dependsOn := range task.DependsOn {
		nokocore.KeepVoid(i)

		dependsTask := p.GetProcessTask(dependsOn.GetTarget())
		temp = append(temp, dependsTask)
	}

	return temp
}

func (p *ProcessTasks) StartProcessTask(pTask ProcessTaskImpl) error {
	var err error
	nokocore.KeepVoid(err)

	p.locker.Lock(func() {
		task := pTask.GetTaskConfig()
		taskName := strings.TrimSpace(task.Name)
		if p.GetProcessTask(taskName) != nil {
			return
		}

		p.pTasks = append(p.pTasks, pTask)
		process := pTask.GetProcess()
		err = process.Start()
	})
	return err
}

func (p *ProcessTasks) ExecuteAsync(tasksConfig *TasksConfig) {
	p.regis = tasksConfig.ApplyAsync(p)
}

func (p *ProcessTasks) Execute(tasksConfig *TasksConfig) error {
	return tasksConfig.Apply(p)
}

func (p *ProcessTasks) Wait() error {
	var err error
	var state StateImpl
	nokocore.KeepVoid(err, state)

	if p.regis != nil {
		if err = p.regis.Wait(); err != nil {
			return err
		}
	}

	for i, pTask := range p.pTasks {
		nokocore.KeepVoid(i)

		process := pTask.GetProcess()
		config := pTask.GetTaskConfig()

		if state, err = process.Wait(); err != nil {
			fmt.Printf("[ERROR] Task '%s' failed.\n", config.Name)
			return err
		}

		if state != nil {
			fmt.Printf("[STATE] PID: %d\n", state.Pid())
			fmt.Printf("[STATE] Exit Code: %d\n", state.ExitCode())
		}
	}

	return nil
}

func makeProcessFromTask(pTasks ProcessTasksImpl, task *Config) error {
	var err error
	var workDir nokocore.WorkingDirImpl
	nokocore.KeepVoid(err, workDir)

	taskName := strings.TrimSpace(task.Name)
	if pTasks.GetProcessTask(taskName) != nil {
		return nil
	}

	fmt.Printf("[RUN] Task '%s' started.\n", taskName)

	if strings.EqualFold(taskName, "self") {
		if err = pTasks.(*ProcessTasks).applyMainTask(task); err != nil {
			return err
		}
		return nil
	}

	return runTask(pTasks, task)
}

func makeProcessFromTaskAsync(pTasks ProcessTasksImpl, task *Config, err chan<- error) {
	err <- makeProcessFromTask(pTasks, task)
}

func waitRun(tasksConfig *TasksConfig, pTasks ProcessTasksImpl, task *Config) error {
	var err error
	nokocore.KeepVoid(err)

	taskName := strings.TrimSpace(task.Name)

	for i, dependsOnTask := range tasksConfig.GetDependsOnTaskConfig(task) {
		nokocore.KeepVoid(i)

		target := dependsOnTask.GetTarget()
		waiter := dependsOnTask.GetWaiter()
		params := dependsOnTask.GetParams()

		targetName := strings.TrimSpace(target.Name)

		// Detect circular dependency between tasksConfig
		if strings.EqualFold(targetName, taskName) {
			return errors.New("circular dependency detected")
		}

		if err = waitRun(tasksConfig, pTasks, target); err != nil {
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

	if err = makeProcessFromTask(pTasks, task); err != nil {
		fmt.Printf("[RUN] Task '%s' failed.\n", taskName)
		return err
	}
	return nil
}

func waitRunTask(tasksConfig *TasksConfig, pTasks ProcessTasksImpl, task *Config) error {
	var ok bool
	var err error
	nokocore.KeepVoid(ok, err)

	taskName := strings.TrimSpace(task.Name)
	dependsOnTasks := tasksConfig.GetDependsOnTaskConfig(task)

	for i, dependsOnTask := range dependsOnTasks {
		nokocore.KeepVoid(i)

		target := dependsOnTask.GetTarget()
		waiter := dependsOnTask.GetWaiter()
		params := dependsOnTask.GetParams()

		targetName := strings.TrimSpace(target.Name)

		// detect circular dependency between tasksConfig
		if strings.EqualFold(targetName, taskName) {
			return errors.New("circular dependency detected")
		}

		if err = waitRunTask(tasksConfig, pTasks, target); err != nil {
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

	errorStack := make(chan error, 1)
	defer close(errorStack)

	go makeProcessFromTaskAsync(pTasks, task, errorStack)

	for {
		select {
		case err, ok = <-errorStack:
			if !ok {
				fmt.Println("[WARN] Channel closed.")
				return nil
			}
			if err != nil {
				fmt.Printf("[RUN] Task '%s' failed.\n", taskName)
				return err
			}
			return nil
		}
	}
}

func (t *TasksConfig) Apply(pTasks ProcessTasksImpl) error {
	var err error
	nokocore.KeepVoid(err)

	tasksConfig := *t
	for i, task := range tasksConfig {
		nokocore.KeepVoid(i)
		if err = waitRun(t, pTasks, task); err != nil {
			return err
		}
	}

	return nil
}

type WaitTasksImpl interface {
	Wait() error
	Add(delta int)
	Run(action nokocore.ActionReturn[error])
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
	t.err = action.Call()
}

func (t *TasksConfig) ApplyAsync(pTasks ProcessTasksImpl) *WaitTasks {
	var err error
	nokocore.KeepVoid(err)

	tasksConfig := *t
	size := len(tasksConfig)

	wt := NewWaitTasks()
	wt.Add(size)

	for i, task := range tasksConfig {
		nokocore.KeepVoid(i)

		// no need goroutine's for a wait run task, already run in goroutine.
		wt.Run(func() error {
			return waitRunTask(t, pTasks, task)
		})
	}

	return wt
}
