package main

type Data struct {
	line string
}

type Storer interface {
	Store(d *Data) error
}

type Puller interface {
	Pull(*Data) error
}

type PullStorer interface {
	Puller
	Storer
}

func Pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func Store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := Pull(ps, data)
		if i > 0 {
			if _, err := Store(ps, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}
