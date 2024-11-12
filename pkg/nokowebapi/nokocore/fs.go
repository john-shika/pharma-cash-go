package nokocore

import (
	"errors"
	"os"
	"path/filepath"
)

func GetSourceRootDir() (string, error) {
	var err error
	var sourceRootPath string
	KeepVoid(err, sourceRootPath)

	// check if there are any command-line arguments provided.
	// if none, return an error indicating the source root directory cannot be determined.
	if len(os.Args) == 0 {
		return "", errors.New("can't get source root dir")
	}

	if sourceRootPath, err = filepath.Abs(os.Args[0]); err != nil {
		return "", err
	}
	sourceRootDir := filepath.Dir(sourceRootPath)
	return sourceRootDir, nil
}

func GetCurrentWorkingDir() (string, error) {
	return os.Getwd()
}

type WorkingDirImpl interface {
	GetSourceRootDir() string
	GetCurrentWorkingDir() string
}

type WorkingDir struct {
	sourceRootDir string
	currWorkDir   string
}

func NewWorkingDir(scriptRootDir, currentWorkingDir string) WorkingDirImpl {
	return &WorkingDir{scriptRootDir, currentWorkingDir}
}

func (s *WorkingDir) GetSourceRootDir() string {
	return s.sourceRootDir
}

func (s *WorkingDir) GetCurrentWorkingDir() string {
	return s.currWorkDir
}

type WorkingFunc func(workDir WorkingDirImpl) error

func (w WorkingFunc) Call(workDir WorkingDirImpl) error {
	return w(workDir)
}

func SetWorkingDir(cb WorkingFunc) (WorkingDirImpl, error) {
	var err error
	var sourceRootDir string
	var currWorkDir string
	KeepVoid(err, sourceRootDir, currWorkDir)

	if sourceRootDir, err = GetSourceRootDir(); err != nil {
		return nil, err
	}

	if currWorkDir, err = GetCurrentWorkingDir(); err != nil {
		return nil, err
	}

	workDir := NewWorkingDir(sourceRootDir, currWorkDir)
	if err = os.Chdir(sourceRootDir); err != nil {
		return nil, err
	}

	if err = cb.Call(workDir); err != nil {
		return workDir, err
	}

	if err = os.Chdir(currWorkDir); err != nil {
		return nil, err
	}
	return workDir, nil
}
