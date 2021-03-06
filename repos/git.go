package repos

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/varutil"
)

type GitRepository struct {
	path string
}

func NewGitRepository(path string) Repository {
	r := &GitRepository{}
	r.Init(path)
	return r
}

func (r *GitRepository) Init(path string) {
	varutil.FixDirPath(&path)
	r.path = path
}

func (r *GitRepository) Clone(url string) error {
	if !disk.IsDir(r.path) {
		os.MkdirAll(r.path, 0777)
	}

	cmd := exec.Command("git", "clone", url, r.path)
	cmd.Dir = r.path
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf(string(out), err)
	}

	return nil
}

func (r *GitRepository) Checkout(rev string) error {
	cmd := exec.Command("git", "checkout", rev)
	cmd.Dir = r.path
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf(string(out), err)
	}
	return nil
}

func (r *GitRepository) Pull() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = r.path
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf(string(out), err)
	}
	return nil
}

func (r *GitRepository) Uninit() error {
	return os.RemoveAll(r.path + ".git")
}
