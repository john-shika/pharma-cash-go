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
	Iterations int           `mapstructure:"iterations" json:"iterations"`
	Duration   time.Duration `mapstructure:"duration" json:"duration"`
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

type DependsOnTaskConfigImpl[T any] interface {
	GetTarget() T
	GetWaiter() string
	GetParams() DependsOnConfigParamsImpl
}

// DependsOnTaskConfig struct, T any can be string or *Config itself.
// Keep in mind, parsing by viper config file, declare T must be string.
type DependsOnTaskConfig[T any] struct {
	Target T                      `mapstructure:"target" json:"target"`
	Waiter string                 `mapstructure:"waiter" json:"waiter"`
	Params *DependsOnConfigParams `mapstructure:"params" json:"params"`
}

func NewDependsOnTaskConfig[T any](target T, waiter string, params DependsOnConfigParamsImpl) DependsOnTaskConfigImpl[T] {
	return &DependsOnTaskConfig[T]{
		Target: target,
		Waiter: waiter,
		Params: &DependsOnConfigParams{
			Iterations: params.GetIterations(),
			Duration:   params.GetDuration(),
		},
	}
}

func (d *DependsOnTaskConfig[T]) GetTarget() T {
	return d.Target
}

func (d *DependsOnTaskConfig[T]) GetWaiter() string {
	return d.Waiter
}

func (d *DependsOnTaskConfig[T]) GetParams() DependsOnConfigParamsImpl {
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

type ConfigImpl interface {
	GetName() string
	SetName(name string)
	GetExec() string
	SetExec(exec string)
	GetArgs() []string
	SetArgs(args []string)
	GetWorkdir() string
	SetWorkdir(workdir string)
	GetEnviron() []string
	SetEnviron(environ []string)
	GetStdin() io.Reader
	SetStdin(stdin string)
	GetStdout() io.Writer
	SetStdout(stdout string)
	GetStderr() io.Writer
	SetStderr(stderr string)
	GetNetwork() *nokocore.NetworkConfig
	SetNetwork(network *nokocore.NetworkConfig)
	GetDependsOn() []*DependsOnTaskConfig[string]
	SetDependsOn(dependsOn []*DependsOnTaskConfig[string])
}

type Config struct {
	Name      string                         `mapstructure:"name" json:"name"`
	Exec      string                         `mapstructure:"exec" json:"exec"`
	Args      []string                       `mapstructure:"args" json:"args"`
	Workdir   string                         `mapstructure:"workdir" json:"workdir"`
	Environ   []string                       `mapstructure:"environ" json:"environ"`
	Stdin     string                         `mapstructure:"stdin" json:"stdin"`
	Stdout    string                         `mapstructure:"stdout" json:"stdout"`
	Stderr    string                         `mapstructure:"stderr" json:"stderr"`
	Network   *nokocore.NetworkConfig        `mapstructure:"network" json:"network"`
	DependsOn []*DependsOnTaskConfig[string] `mapstructure:"depends_on" json:"dependsOn"`
}

func NewConfig(name string, exec string, args []string, workdir string, environ []string, stdin string, stdout string, stderr string, network *nokocore.NetworkConfig, dependsOn []*DependsOnTaskConfig[string]) ConfigImpl {
	return &Config{
		Name:      name,
		Exec:      exec,
		Args:      args,
		Workdir:   workdir,
		Environ:   environ,
		Stdin:     stdin,
		Stdout:    stdout,
		Stderr:    stderr,
		Network:   network,
		DependsOn: dependsOn,
	}
}

func (Config) GetNameType() string {
	return "Task"
}

func (w *Config) GetName() string {
	return strings.TrimSpace(w.Name)
}

func (w *Config) SetName(name string) {
	w.Name = name
}

func (w *Config) GetExec() string {
	return w.Exec
}

func (w *Config) SetExec(exec string) {
	w.Exec = exec
}

func (w *Config) GetArgs() []string {
	return w.Args
}

func (w *Config) SetArgs(args []string) {
	w.Args = args
}

func (w *Config) GetWorkdir() string {
	return w.Workdir
}

func (w *Config) SetWorkdir(workdir string) {
	w.Workdir = workdir
}

func (w *Config) GetEnviron() []string {
	return w.Environ
}

func (w *Config) SetEnviron(environ []string) {
	w.Environ = environ
}

func (w *Config) GetStdin() io.Reader {
	switch strings.ToLower(strings.TrimSpace(w.Stdin)) {
	case "console":
		return nokocore.NewSafeReader(xterm.Stdin)
	default:
		return nil
	}
}

func (w *Config) SetStdin(stdin string) {
	w.Stdin = stdin
}

func (w *Config) GetStdout() io.Writer {
	switch strings.ToLower(strings.TrimSpace(w.Stdout)) {
	case "console":
		return nokocore.NewSafeWriter(xterm.Stdout)
	default:
		return nil
	}
}

func (w *Config) SetStdout(stdout string) {
	w.Stdout = stdout
}

func (w *Config) GetStderr() io.Writer {
	switch strings.ToLower(strings.TrimSpace(w.Stderr)) {
	case "console":
		return nokocore.NewSafeWriter(xterm.Stderr)
	default:
		return nil
	}
}

func (w *Config) SetStderr(stderr string) {
	w.Stderr = stderr
}

func (w *Config) GetNetwork() *nokocore.NetworkConfig {
	return w.Network
}

func (w *Config) SetNetwork(network *nokocore.NetworkConfig) {
	w.Network = network
}

func (w *Config) GetDependsOn() []*DependsOnTaskConfig[string] {
	return w.DependsOn
}

func (w *Config) SetDependsOn(dependsOn []*DependsOnTaskConfig[string]) {
	w.DependsOn = dependsOn
}

type TasksImpl interface {
	Size() int
	Index(i int) ConfigImpl
	SetIndex(i int, config ConfigImpl)
	Append(config ConfigImpl)
	Remove(i int)
	GetNameType() string
	GetTaskConfig(name string) ConfigImpl
	GetDependsOnTaskConfig(task ConfigImpl) []DependsOnTaskConfigImpl[ConfigImpl]
}

// Tasks struct, keep it mind, parsing by viper config file
type Tasks []Config

func NewTasks() Tasks {
	var temp Tasks
	return temp
}

func (Tasks) GetNameType() string {
	return "Tasks"
}

func (w *Tasks) Size() int {
	return len(*w)
}

func (w *Tasks) Index(i int) ConfigImpl {
	return &(*w)[i]
}

func (w *Tasks) SetIndex(i int, config ConfigImpl) {
	(*w)[i] = *config.(*Config)
}

func (w *Tasks) Append(config ConfigImpl) {
	*w = append(*w, *config.(*Config))
}

func (w *Tasks) Remove(i int) {
	j := i + 1
	if j < len(*w) {
		*w = append((*w)[:i], (*w)[j:]...)
		return
	}

	*w = (*w)[:i]
}

func (w *Tasks) GetTaskConfig(name string) ConfigImpl {
	for i, task := range *w {
		nokocore.KeepVoid(i)

		name = strings.TrimSpace(name)
		if strings.EqualFold(task.GetName(), name) {
			return &task
		}
	}

	// task not found
	//panic(fmt.Sprintf("task '%s' not found", name))
	return nil
}

func (w *Tasks) GetDependsOnTaskConfig(task ConfigImpl) []DependsOnTaskConfigImpl[ConfigImpl] {
	var temp []DependsOnTaskConfigImpl[ConfigImpl]
	for i, dependsOn := range task.GetDependsOn() {
		nokocore.KeepVoid(i)

		target := dependsOn.GetTarget()
		waiter := dependsOn.GetWaiter()
		params := dependsOn.GetParams()

		config := w.GetTaskConfig(target)
		dependsOnTask := NewDependsOnTaskConfig(config, waiter, params)
		temp = append(temp, dependsOnTask)
	}

	return temp
}

type ProcessTaskImpl interface {
	GetProcess() ProcessImpl
	GetTaskConfig() ConfigImpl
	IsRunning() bool
}

type ProcessTask struct {
	TaskConfig ConfigImpl
	Process    ProcessImpl
}

func NewProcessTask(process ProcessImpl, config ConfigImpl) ProcessTaskImpl {
	return &ProcessTask{
		TaskConfig: config,
		Process:    process,
	}
}

func (p *ProcessTask) GetProcess() ProcessImpl {
	return p.Process
}

func (p *ProcessTask) GetTaskConfig() ConfigImpl {
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
	mainTaskHelper(task ConfigImpl) error
	startTaskHelper(task ConfigImpl) error
	GetProcessTask(name string) ProcessTaskImpl
	GetDependsOnProcessTask(task ConfigImpl) []ProcessTaskImpl
	StartProcessTask(pTask ProcessTaskImpl) error
	ExecuteAsync(tasks TasksImpl)
	Execute(tasks TasksImpl) error
	Wait() error
}

type ProcessTasks struct {
	//mainTaskHelper nokocore.ActionDoubleParamsReturn[ProcessTasksImpl, ConfigImpl, error]
	pTasks []ProcessTaskImpl
	locker nokocore.LockerImpl
	regis  WaitTasksImpl
}

func NewProcessTasks() ProcessTasksImpl {
	return &ProcessTasks{
		pTasks: []ProcessTaskImpl{},
		locker: nokocore.NewLocker(),
	}
}

func (p *ProcessTasks) startTaskHelper(task ConfigImpl) error {
	var err error
	var workDir nokocore.WorkingDirImpl
	nokocore.KeepVoid(err, workDir)

	// try to dial url it-self
	if network := task.GetNetwork(); network != nil {
		if nokocore.TryFetchUrl(network.GetURL()) {
			fmt.Printf("[FETCH] URL '%s' is alive.\n", network.GetURL().String())
			return nil
		}
	}

	workFunc := func(workDir nokocore.WorkingDirImpl) error {
		if err = os.Chdir(task.GetWorkdir()); err != nil {
			return err
		}

		process := NewProcess(task.GetExec(), task.GetArgs()...)

		// binding stdin, stdout, stderr
		process.SetStdin(task.GetStdin())
		process.SetStdout(task.GetStdout())
		process.SetStderr(task.GetStderr())
		process.SetEnviron(task.GetEnviron())

		pTask := NewProcessTask(process, task)
		return p.StartProcessTask(pTask)
	}

	if workDir, err = nokocore.SetWorkingDir(workFunc); err != nil {
		return err
	}

	return nil
}

func (p *ProcessTasks) mainTaskHelper(task ConfigImpl) error {
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

	// get arguments
	args = os.Args

	// binding values
	task.SetExec(exec)
	task.SetArgs(args)

	// set value for NOKOWEBAPI_SELF_RUNNING env
	nokoWebApiAutoRunEnv := fmt.Sprintf("%s=%s", NokoWebApiAutoRunEnv, "1")
	environ := append(task.GetEnviron(), nokoWebApiAutoRunEnv)
	task.SetEnviron(environ)

	return p.startTaskHelper(task)
}

func (p *ProcessTasks) applyMainTask(task ConfigImpl) error {
	var err error
	nokocore.KeepVoid(err)
	if err = p.mainTaskHelper(task); err != nil {
		return err
	}
	return nil
}

func (p *ProcessTasks) GetProcessTask(name string) ProcessTaskImpl {
	for i, pTask := range p.pTasks {
		nokocore.KeepVoid(i)
		task := pTask.GetTaskConfig()
		if strings.EqualFold(task.GetName(), name) {
			return pTask
		}
	}
	return nil
}

func (p *ProcessTasks) GetDependsOnProcessTask(task ConfigImpl) []ProcessTaskImpl {
	var temp []ProcessTaskImpl
	for i, dependsOn := range task.GetDependsOn() {
		nokocore.KeepVoid(i)

		dependsOnTask := p.GetProcessTask(dependsOn.GetTarget())
		temp = append(temp, dependsOnTask)
	}

	return temp
}

func (p *ProcessTasks) StartProcessTask(pTask ProcessTaskImpl) error {
	var err error
	nokocore.KeepVoid(err)

	p.locker.Lock(func() {
		task := pTask.GetTaskConfig()
		if p.GetProcessTask(task.GetName()) != nil {
			return
		}

		p.pTasks = append(p.pTasks, pTask)
		process := pTask.GetProcess()
		err = process.Start()
	})
	return err
}

func (p *ProcessTasks) ExecuteAsync(tasks TasksImpl) {
	var err error
	nokocore.KeepVoid(err)

	size := tasks.Size()
	waitTasks := NewWaitTasks()
	waitTasks.Add(size)

	for i := 0; i < size; i++ {
		task := tasks.Index(i)
		waitTasks.Run(func() error {
			return waitRunTask(tasks, p, task)
		})
	}

	p.regis = waitTasks
}

func (p *ProcessTasks) Execute(tasks TasksImpl) error {
	var err error
	nokocore.KeepVoid(err)

	size := tasks.Size()
	for i := 0; i < size; i++ {
		task := tasks.Index(i)
		if err = waitRun(tasks, p, task); err != nil {
			return err
		}
	}

	return nil
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
			fmt.Printf("[ERROR] Task '%s' failed.\n", config.GetName())
			return err
		}

		if state != nil {
			fmt.Printf("[STATE] PID: %d\n", state.Pid())
			fmt.Printf("[STATE] Exit Code: %d\n", state.ExitCode())
		}
	}

	return nil
}

func makeProcessFromTask(pTasks ProcessTasksImpl, task ConfigImpl) error {
	var err error
	var workDir nokocore.WorkingDirImpl
	nokocore.KeepVoid(err, workDir)

	if pTasks.GetProcessTask(task.GetName()) != nil {
		return nil
	}

	fmt.Printf("[RUN] Task '%s' started.\n", task.GetName())

	if strings.EqualFold(task.GetName(), "self") {
		if err = pTasks.(*ProcessTasks).applyMainTask(task); err != nil {
			return err
		}
		return nil
	}

	return pTasks.startTaskHelper(task)
}

func makeProcessFromTaskAsync(pTasks ProcessTasksImpl, task ConfigImpl, err chan<- error) {
	err <- makeProcessFromTask(pTasks, task)
}

func waitRun(tasks TasksImpl, pTasks ProcessTasksImpl, task ConfigImpl) error {
	var err error
	nokocore.KeepVoid(err)

	for i, dependsOnTask := range tasks.GetDependsOnTaskConfig(task) {
		nokocore.KeepVoid(i)

		target := dependsOnTask.GetTarget()
		waiter := dependsOnTask.GetWaiter()
		params := dependsOnTask.GetParams()

		// Detect circular dependency between tasks
		if strings.EqualFold(target.GetName(), task.GetName()) {
			return errors.New("circular dependency detected")
		}

		if err = waitRun(tasks, pTasks, target); err != nil {
			return err
		}

		switch waiter {
		case "none":
			break
		case "wait-for-http-alive":
			network := target.GetNetwork()
			iterations := params.GetIterations()
			duration := params.GetDuration()

			// wait for http alive
			network.WaitForHttpAlive(iterations, duration)
			break
		default:
			return errors.New("unknown waiter")
		}
	}

	if err = makeProcessFromTask(pTasks, task); err != nil {
		fmt.Printf("[RUN] Task '%s' failed.\n", task.GetName())
		return err
	}
	return nil
}

func waitRunTask(tasks TasksImpl, pTasks ProcessTasksImpl, task ConfigImpl) error {
	var ok bool
	var err error
	nokocore.KeepVoid(ok, err)

	dependsOnTasks := tasks.GetDependsOnTaskConfig(task)

	for i, dependsOnTask := range dependsOnTasks {
		nokocore.KeepVoid(i)

		target := dependsOnTask.GetTarget()
		waiter := dependsOnTask.GetWaiter()
		params := dependsOnTask.GetParams()

		// detect circular dependency between tasks
		if strings.EqualFold(target.GetName(), task.GetName()) {
			return errors.New("circular dependency detected")
		}

		if err = waitRunTask(tasks, pTasks, target); err != nil {
			return err
		}

		switch waiter {
		case "none":
			break
		case "wait-for-http-alive":
			network := target.GetNetwork()
			iterations := params.GetIterations()
			duration := params.GetDuration()

			// wait for http alive
			network.WaitForHttpAlive(iterations, duration)
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
				fmt.Printf("[RUN] Task '%s' failed.\n", task.GetName())
				return err
			}
			return nil
		}
	}
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

func (w *WaitTasks) Wait() error {
	w.WaitGroup.Wait()
	return w.err
}

func (w *WaitTasks) Add(delta int) {
	w.WaitGroup.Add(delta)
}

func (w *WaitTasks) Run(action nokocore.ActionReturn[error]) {
	defer w.WaitGroup.Done()
	w.err = action.Call()
}
