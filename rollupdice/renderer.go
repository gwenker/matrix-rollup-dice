package rollupdice

import (
	"strconv"

	"github.com/richardlt/matrix/sdk-go/common"
	"github.com/richardlt/matrix/sdk-go/software"
)

func newRenderer(a software.API) (*renderer, error) {
	backgroundLayer, err := a.NewLayer()
	if err != nil {
		return nil, err
	}
	valuesLayer, err := a.NewLayer()
	if err != nil {
		return nil, err
	}

	cd, err := valuesLayer.NewCaracterDriver(a.GetFontFromLocal("DiceValues"))
	if err != nil {
		return nil, err
	}

	white, err := a.GetColorFromRemoteThemeByName("flat", "white_1")
	if err != nil {
		return nil, err
	}

	grey, err := a.GetColorFromRemoteThemeByName("flat", "dark_grey_1")
	if err != nil {
		return nil, err
	}

	return &renderer{
		api:             a,
		backgroundLayer: backgroundLayer,
		valuesLayer:     valuesLayer,
		caracterDriver:  cd,
		backgroundColor: white,
		dotColor:        grey,
	}, nil
}

type renderer struct {
	api                          software.API
	backgroundLayer, valuesLayer software.Layer
	caracterDriver               *software.CaracterDriver
	backgroundColor, dotColor    common.Color
}

func (r *renderer) Clean() {
	r.backgroundLayer.Clean()
	r.valuesLayer.Clean()
}

func (r *renderer) Print(value1, value2 int) {
	r.backgroundLayer.Clean()
	r.valuesLayer.Clean()

	// draw background for dices
	for x := 0; x < 16; x++ {
		if x != 7 && x != 8 {
			for y := 1; y < 8; y++ {
				r.backgroundLayer.SetWithCoord(
					common.Coord{X: int64(x), Y: int64(y)},
					r.backgroundColor,
				)
			}
		}
	}

	// print values
	r.caracterDriver.Render([]rune(strconv.Itoa(value1))[0],
		common.Coord{X: 3, Y: 4}, r.dotColor, common.Color{})
	r.caracterDriver.Render([]rune(strconv.Itoa(value2))[0],
		common.Coord{X: 12, Y: 4}, r.dotColor, common.Color{})

	r.api.Print()
}
