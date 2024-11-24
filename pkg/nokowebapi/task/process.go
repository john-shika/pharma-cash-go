package task

import (
	"errors"
	"io"
	"nokowebapi/nokocore"
	"nokowebapi/xterm"
	"os/exec"
	"strings"
)

type StateImpl interface {
	Exited() bool
	ExitCode() int
	Pid() int
}

type ProcessStateOverlay struct {
	exited   bool
	exitCode int
	pid      int
}

func NewProcessStateOverlay(exited bool, exitCode int, pid int) StateImpl {
	return &ProcessStateOverlay{
		exited:   exited,
		exitCode: exitCode,
		pid:      pid,
	}
}

func (s *ProcessStateOverlay) Exited() bool {
	return s.exited
}

func (s *ProcessStateOverlay) ExitCode() int {
	return s.exitCode
}

func (s *ProcessStateOverlay) Pid() int {
	return s.pid
}

func MakeProcess(path string, args []string, environ []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (executor *exec.Cmd, err error) {
	if path, err = exec.LookPath(path); err != nil {
		return nil, err
	}

	// create command executor
	executor = exec.Command(path, args...)

	// binding all parameters
	executor.Stdin = stdin
	executor.Stdout = stdout
	executor.Stderr = stderr
	executor.Env = environ

	// return command executor
	return executor, nil
}

func GetProcessState(cmd *exec.Cmd) (StateImpl, error) {
	if cmd.ProcessState != nil {
		return cmd.ProcessState, nil
	}
	if cmd.Process != nil {
		return NewProcessStateOverlay(false, 0, cmd.Process.Pid), nil
	}
	return nil, errors.New("invalid process state")
}

type ProcessImpl interface {
	IsRunning() bool
	Environ() []string
	SetEnviron([]string)
	HasEnv(key string) bool
	GetEnv(key string) string
	SetEnv(key string, value string) bool
	DelEnv(key string) bool
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
	SetStdin(stdin io.Reader)
	SetStdout(stdout io.Writer)
	SetStderr(stderr io.Writer)
	Wait() (StateImpl, error)
	State() (StateImpl, error)
	Start() error
	Run() error
}

type Process struct {
	path     string
	args     []string
	executor *exec.Cmd
	environ  []string
	stdin    io.Reader
	stdout   io.Writer
	stderr   io.Writer
}

func NewProcess(path string, args ...string) ProcessImpl {
	var environ []string
	return &Process{
		path:     path,
		args:     args,
		executor: nil,
		stdin:    xterm.Stdin,
		stdout:   xterm.Stdout,
		stderr:   xterm.Stderr,
		environ:  environ,
	}
}

func (p *Process) IsRunning() bool {
	var err error
	var state StateImpl
	nokocore.KeepVoid(err, state)

	if state, err = p.State(); err != nil {
		return false
	}

	return state != nil && !state.Exited()
}

func (p *Process) Environ() []string {
	return p.environ
}

func (p *Process) SetEnviron(environ []string) {
	p.environ = environ
}

func (p *Process) HasEnv(key string) bool {
	for i, env := range p.environ {
		nokocore.KeepVoid(i)

		name := key + "="
		if strings.HasPrefix(env, name) {
			return true
		}
	}
	return false
}

func (p *Process) GetEnv(key string) string {
	for i, env := range p.environ {
		nokocore.KeepVoid(i)

		name := key + "="
		if strings.HasPrefix(env, name) {
			return env
		}
	}
	panic("invalid key")
}

func (p *Process) SetEnv(key, value string) bool {
	if p.HasEnv(key) {

		for i, env := range p.environ {
			nokocore.KeepVoid(i)

			name := key + "="
			if strings.HasPrefix(env, name) {
				p.environ[i] = key + "=" + value
				return true
			}
		}

	} else {
		p.environ = append(p.environ, key+"="+value)
		return true
	}

	return false
}

func (p *Process) DelEnv(key string) bool {
	var found bool
	var temp []string
	nokocore.KeepVoid(found, temp)
	
	if p.HasEnv(key) {
		for i, env := range p.environ {
			nokocore.KeepVoid(i)

			name := key + "="
			if strings.HasPrefix(env, name) {
				found = true
				continue
			}

			temp = append(temp, env)
		}
	}

	if found {
		p.environ = temp
		return true
	}
	return false
}

func (p *Process) Stdin() io.Reader {
	return p.stdin
}

func (p *Process) Stdout() io.Writer {
	return p.stdout
}

func (p *Process) Stderr() io.Writer {
	return p.stderr
}

func (p *Process) SetStdin(stdin io.Reader) {
	p.stdin = stdin
}

func (p *Process) SetStdout(stdout io.Writer) {
	p.stdout = stdout
}

func (p *Process) SetStderr(stderr io.Writer) {
	p.stderr = stderr
}

func (p *Process) Wait() (StateImpl, error) {
	var err error
	var state StateImpl
	nokocore.KeepVoid(err, state)

	if p.executor == nil {
		if p.executor, err = MakeProcess(p.path, p.args, p.environ, p.stdin, p.stdout, p.stderr); err != nil {
			return nil, err
		}
		if err = p.executor.Start(); err != nil {
			return nil, err
		}
	}

	err = p.executor.Wait()
	if state, err = GetProcessState(p.executor); err != nil {
		return nil, err
	}

	return state, err
}

func (p *Process) State() (StateImpl, error) {
	if p.executor == nil {
		return nil, errors.New("process not started")
	}
	return GetProcessState(p.executor)
}

func (p *Process) Start() error {
	var err error
	nokocore.KeepVoid(err)

	if p.executor == nil {
		if p.executor, err = MakeProcess(p.path, p.args, p.environ, p.stdin, p.stdout, p.stderr); err != nil {
			return err
		}
		if err = p.executor.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Process) Run() error {
	var err error
	nokocore.KeepVoid(err)

	if p.executor == nil {
		if p.executor, err = MakeProcess(p.path, p.args, p.environ, p.stdin, p.stdout, p.stderr); err != nil {
			return err
		}
		if err = p.executor.Run(); err != nil {
			return err
		}
	}
	return nil
}
