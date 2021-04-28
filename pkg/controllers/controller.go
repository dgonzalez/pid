package controllers

type Controller interface {
	SetTarget(target float64)
	Update(target float64) float64
}
