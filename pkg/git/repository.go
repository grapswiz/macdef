package git

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
	"os"
	"os/user"
)

type Repository struct {
	URL       string
	Directory string
	Entity    *git.Repository
}

func (r *Repository) Update() error {
	w, err := r.Entity.Worktree()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewRepository(remoteURL string, localDirectory string, progress sideband.Progress) (*Repository, error) {
	usr, err := user.Current()
	directory := usr.HomeDir + "/" + localDirectory
	r := &Repository{
		URL:       remoteURL,
		Directory: directory,
		Entity:    nil,
	}
	if err != nil {
		return nil, err
	}
	os.MkdirAll(directory, 0755)
	if entity, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: progress,
	}); err == git.ErrRepositoryAlreadyExists {
		r.Entity, err = git.PlainOpen(directory)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		r.Entity = entity
	}
	return r, nil
}
