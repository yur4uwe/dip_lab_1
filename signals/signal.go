package signals

var (
	count = 0
)

type Signal struct {
	amplitude float64
	sig_type  string
	id        int
}

func NewSignal(amplitude float64, sig_type string) *Signal {
	count++
	return &Signal{amplitude: amplitude, sig_type: sig_type, id: count}
}

func (s *Signal) Amplitude() float64 {
	return s.amplitude
}

func (s *Signal) Type() string {
	return s.sig_type
}

func (s *Signal) ID() int {
	return s.id
}
