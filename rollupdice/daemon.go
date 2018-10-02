package rollupdice

import (
	"github.com/richardlt/matrix/sdk-go/common"
	"github.com/richardlt/matrix/sdk-go/software"
	"github.com/sirupsen/logrus"
)

// Start the rollup dice software.
func Start(uri string) error {
	logrus.Infof("Start rollup dice for uri %s\n", uri)

	rd := &rollupDice{}

	return software.Connect(uri, rd, true)
}

type rollupDice struct {
	engine   *engine
	renderer *renderer
}

func (r *rollupDice) Init(a software.API) (err error) {
	logrus.Debug("Init rollup dice")

	r.renderer, err = newRenderer(a)
	if err != nil {
		return err
	}

	l := a.GetImageFromLocal("dice")

	a.SetConfig(software.ConnectRequest_SoftwareData_Config{
		Logo:           &l,
		MinPlayerCount: 1,
		MaxPlayerCount: 1,
	})

	return a.Ready()
}

func (r *rollupDice) Start(playerCount uint64) {
	r.engine = newEngine()
	r.engine.GenerateValues()
	r.print()
}

func (r *rollupDice) Close() { r.renderer.Clean() }

func (r *rollupDice) ActionReceived(slot uint64, cmd common.Command) {
	switch cmd {
	case common.Command_A_UP:
		r.engine.GenerateValues()
	}
	r.print()
}

func (r *rollupDice) print() { r.renderer.Print(r.engine.GetValues()) }
