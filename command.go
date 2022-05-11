package tmcl

const ABS byte = 0
const REL byte = 1
const COORD byte = 2

func (q *TMCL) ROR(motor byte, velocity int) error {
	_, err := q.Exec(1, 0, motor, velocity)
	return err
}

func (q *TMCL) ROL(motor byte, velocity int) error {
	_, err := q.Exec(2, 0, motor, velocity)
	return err
}

func (q *TMCL) MST(motor byte) error {
	_, err := q.Exec(3, 0, motor, 0)
	return err
}

func (q *TMCL) MVP(mode byte, motor byte, value int) error {
	_, err := q.Exec(4, mode, motor, value)
	return err
}

func (q *TMCL) SAP(index byte, motor byte, value int) error {
	_, err := q.Exec(5, index, motor, value)
	return err
}

func (q *TMCL) GAP(index byte, motor byte) (int, error) {
	return q.Exec(6, index, motor, 0)
}

func (q *TMCL) STAP(index byte, motor byte) error {
	_, err := q.Exec(7, index, motor, 0)
	return err
}

func (q *TMCL) RSAP(index byte, motor byte) error {
	_, err := q.Exec(8, index, motor, 0)
	return err
}

func (q *TMCL) SGP(index byte, bank byte, value int) error {
	_, err := q.Exec(9, index, bank, value)
	return err
}

func (q *TMCL) GGP(index byte, bank byte) (int, error) {
	return q.Exec(10, index, bank, 0)
}

func (q *TMCL) STGP(index byte, bank byte) (int, error) {
	return q.Exec(11, index, bank, 0)
}

func (q *TMCL) RSGP(index byte, bank byte) (int, error) {
	return q.Exec(12, index, bank, 0)
}

func (q *TMCL) SIO(port byte, value bool) error {
	var b int
	if value {
		b = 1
	}
	_, err := q.Exec(14, port, 2, b)
	return err
}

func (q *TMCL) GIO(port byte, bank byte) (int, error) {
	return q.Exec(15, port, bank, 0)
}
