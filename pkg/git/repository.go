package git

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
	"os"
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
	r := &Repository{
		URL:       remoteURL,
		Directory: localDirectory,
		Entity:    nil,
	}
	os.MkdirAll(localDirectory, 0755)
	if entity, err := git.PlainClone(localDirectory, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: progress,
	}); err == git.ErrRepositoryAlreadyExists {
		r.Entity, err = git.PlainOpen(localDirectory)
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
