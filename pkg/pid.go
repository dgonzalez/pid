package pkg

type PID struct {
	kp            float64
	ki            float64
	kd            float64
	integral      float64
	intervalMsecs int64
	previousError float64
	target        float64
}

func New(p float64, i float64, d float64, intervalMsecs int64) PID {
	return PID{kp: p, ki: i, kd: d, intervalMsecs: intervalMsecs}
}

func (p *PID) SetTarget(target float64) {
	p.target = target
}

func (p *PID) Update(target float64) (result float64) {
	proportional := p.target - target
	p.integral += proportional * float64(p.intervalMsecs)
	derivative := (proportional - p.previousError) / float64(p.intervalMsecs)
	result = p.kp*proportional + p.ki*p.integral + p.kd*derivative
	p.previousError = proportional
	return
}
