package model

type Color struct {
	Id   int
	Name string
	HEX  string
}

func (c Color) ToDto() ColorDto {
	return ColorDto{Id: c.Id, Name: c.Name, HEX: c.HEX}
}

type ColorDto struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,max=20"`
	HEX  string `json:"HEX" validate:"required,hexcolor"`
}

func (d ColorDto) FromDto() Color {
	return Color{Id: d.Id, Name: d.Name, HEX: d.HEX}
}

type ColorAll []Color

type ColorAllDto struct {
	Colors []ColorDto `json:"colors" validate:"required"`
}

func (c ColorAll) ToDto() ColorAllDto {
	dto := ColorAllDto{make([]ColorDto, len(c))}

	for i, model := range c {
		dto.Colors[i] = model.ToDto()
	}
	return dto
}
