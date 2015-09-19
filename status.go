package users

type Status int64

const (
	Banned Status = 1 << iota
	EmailConfirmed
)

func (s Status) Add(toAdd Status) Status {
	return s | toAdd
}

func (s Status) Remove(toRemove Status) Status {
	if s&toRemove == 0 {
		return s
	}
	return s ^ toRemove
}
