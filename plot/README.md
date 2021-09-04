# plot

## Functions

### func [ID](/plot.go#L29)

`func ID(x float64) float64`

ID is the identity function for float64. It returns what it is given.
Useful when you need a noop function.

## Types

### type [FuncF64](/plot.go#L13)

`type FuncF64 func(float64) float64`

FuncF64 is a function that takes a float64 and returns a float64

### type [Plot](/plot.go#L16)

`type Plot struct { ... }`

Plot is the intermediate plot format that will be rendered as ascii or a PNG

#### func (*Plot) [RenderASCII](/plot.go#L95)

`func (plt *Plot) RenderASCII() string`

RenderASCII renders the plot as ascii art

#### func (*Plot) [RenderImage](/plot.go#L69)

`func (plt *Plot) RenderImage() *image.Paletted`

RenderImage renders the plot as a paletted image

#### func (*Plot) [RenderPNG](/plot.go#L120)

`func (plt *Plot) RenderPNG() []byte`

RenderPNG renders the plot as a PNG image
