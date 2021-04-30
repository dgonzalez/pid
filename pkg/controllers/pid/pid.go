package pid

type PID struct {
	kp            float64
	ki            float64
	kd            float64
	integral      float64
	intervalSecs  float64
	previousError float64
	target        float64
}

func New(p, i, d float64, intervalSecs float64) *PID {
	return &PID{kp: p, ki: i, kd: d, intervalSecs: intervalSecs}
}

func (p *PID) SetTarget(target float64) {
	p.target = target
}

func (p *PID) Update(target float64) float64 {
	return p.UpdateAfter(target, p.intervalSecs)
}

func (p *PID) UpdateAfter(target float64, periodSecs float64) (result float64) {
	proportional := p.target - target
	p.integral += proportional * periodSecs
	derivative := (proportional - p.previousError) / periodSecs
	result = p.kp*proportional + p.ki*p.integral + p.kd*derivative
	p.previousError = proportional
	return
}
