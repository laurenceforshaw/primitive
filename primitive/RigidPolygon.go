package primitive

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/raster"
)

type RigidPolygonFactory struct{
	Order int
	X, Y []float64
	Radius float64
}

func(fact *RigidPolygonFactory) NewShape(worker *Worker) Shape{
	return NewRandomRigidPolygon(worker,fact)
}

//parse a rigid polygon line
func ParseRigidPoly(sp []string) (ShapeFactory,error){
	if(len(sp) == 1){
		return nil, fmt.Errorf("RigidPoly requires at least 2 arguments")
	}
	order,err := strconv.Atoi(sp[1])
	if(err != nil){
		return nil,err
	}
	if(order < 3){
		return nil, fmt.Errorf("A polygon must have at least 3 sides.")
	}
	if(len(sp) != order*2 + 2){
		return nil, fmt.Errorf("RigidPoly requires 2 arguments after the first two per vertex")
	}
	X := make([]float64,order)
	Y := make([]float64,order)
	Radius := 0.0
	for i := 0 ;i < order; i++ {
		X[i], err = strconv.ParseFloat(sp[2*i + 2],64)
		if(err != nil){
			return nil ,fmt.Errorf("Error reading user shape file: %s",err.Error())
		}
		Y[i], err = strconv.ParseFloat(sp[2*i + 3],64)
		if(err != nil){
			return nil, fmt.Errorf("Error reading user shape file: %s",err.Error())
		}
		Radius = math.Max(Radius, math.Sqrt(X[i]*X[i] + Y[i]*Y[i]))
	}
	res := RigidPolygonFactory{order,X,Y,Radius}
	return &res, nil
}

type RigidPolygon struct {
	Parent *RigidPolygonFactory
	Worker *Worker
	Order  int
	X, Y   []float64
	RootX, RootY float64
	MaxScale float64
	Scale float64
	Angle float64
}

func NewRandomRigidPolygon(worker *Worker, parent *RigidPolygonFactory) *RigidPolygon {
	rnd := worker.Rnd
	order := parent.Order
	x := make([]float64, order)
	y := make([]float64, order)
	RootX := rnd.Float64() * float64(worker.W)
	RootY := rnd.Float64() * float64(worker.H)
	MaxScale := math.Sqrt(float64(worker.W*worker.W + worker.H*worker.H))/parent.Radius
	Scale := rnd.Float64()*0.1*MaxScale
	Angle := rnd.Float64()*2*math.Pi
	for i := 1; i < order; i++ {
		x[i] = x[0] + rnd.Float64()*40 - 20
		y[i] = y[0] + rnd.Float64()*40 - 20
	}
	p := &RigidPolygon{parent,worker, order, x, y,RootX,RootY,MaxScale,Scale,Angle}
	p.Derive()
	p.Mutate()
	return p
}

//Get the vertex locations from the scale and angle
func (p *RigidPolygon) Derive(){
	for i := 0;i <p.Order;i++{
		p.X[i] = p.RootX + math.Cos(p.Angle)*p.Parent.X[i]*p.Scale - math.Sin(p.Angle)*p.Parent.Y[i]*p.Scale
		p.Y[i] = p.RootY + math.Cos(p.Angle)*p.Parent.Y[i]*p.Scale + math.Sin(p.Angle)*p.Parent.X[i]*p.Scale
	}
}

func (p *RigidPolygon) Draw(dc *gg.Context, scale float64) {
	dc.NewSubPath()
	for i := 0; i < p.Order; i++ {
		dc.LineTo(p.X[i], p.Y[i])
	}
	dc.ClosePath()
	dc.Fill()
}

func (p *RigidPolygon) SVG(attrs string) string {
	ret := fmt.Sprintf(
		"<RigidPolygon %s points=\"",
		attrs)
	points := make([]string, len(p.X))
	for i := 0; i < len(p.X); i++ {
		points[i] = fmt.Sprintf("%f,%f", p.X[i], p.Y[i])
	}

	return ret + strings.Join(points, ",") + "\" />"
}

func (p *RigidPolygon) Copy() Shape {
	a := *p
	a.X = make([]float64, p.Order)
	a.Y = make([]float64, p.Order)
	copy(a.X, p.X)
	copy(a.Y, p.Y)
	return &a
}

func (p *RigidPolygon) Mutate() {
	//const m = 16
	//w := p.Worker.W
	//h := p.Worker.H
	rnd := p.Worker.Rnd
	p.Scale = p.Scale + rnd.NormFloat64()*0.05*p.MaxScale
	p.Angle = p.Angle + rnd.NormFloat64()*0.1
	p.Derive()
}

func (p *RigidPolygon) Valid() bool {
	return true
}


func (p *RigidPolygon) Rasterize() []Scanline {
	var path raster.Path
	for i := 0; i <= p.Order; i++ {
		f := fixp(p.X[i%p.Order], p.Y[i%p.Order])
		if i == 0 {
			path.Start(f)
		} else {
			path.Add1(f)
		}
	}
	return fillPath(p.Worker, path)
}
