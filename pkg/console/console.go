package console

import (
	"context"
	"fmt"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"

	"github.com/harvester/harvester-installer/pkg/config"
	"github.com/harvester/harvester-installer/pkg/widgets"
)

var (
	debug bool
)

func initLogs() error {
	if os.Getenv("DEBUG") == "true" {
		debug = true
		logrus.SetLevel(logrus.DebugLevel)
	}
	f, err := os.OpenFile("/var/log/console.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755) //0600)
	if err != nil {
		return err
	}
	logrus.SetOutput(f)
	return nil
}

// Console is the structure of the harvester console
type Console struct {
	context context.Context
	*gocui.Gui
	elements map[string]widgets.Element
	config   *config.HarvesterConfig
}

// RunConsole starts the console
func RunConsole() error {
	c, err := NewConsole()
	if err != nil {
		return err
	}
	if err := initLogs(); err != nil {
		return err
	}
	return c.doRun()
}

// NewConsole initialize the console
func NewConsole() (*Console, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	conf := config.NewHarvesterConfig()
	conf.SystemSettings = map[string]string{
		"auto-add-disk-paths": "/dev/sdwut",
		"log-level":           "Debug",
		"foo":                 "bar",
	}
	return &Console{
		context:  context.Background(),
		Gui:      g,
		elements: make(map[string]widgets.Element),
		config:   conf,
	}, nil
}

// GetElement gets an element by name
func (c *Console) GetElement(name string) (widgets.Element, error) {
	e, ok := c.elements[name]
	if ok {
		return e, nil
	}
	return nil, fmt.Errorf("element %q is not found", name)
}

// AddElement adds an element with name
func (c *Console) AddElement(name string, element widgets.Element) {
	c.elements[name] = element
}

func (c *Console) setContentByName(name string, content string) error {
	v, err := c.GetElement(name)
	if err != nil {
		return err
	}
	if content == "" {
		return v.Close()
	}
	if err := v.Show(); err != nil {
		return err
	}
	v.SetContent(content)
	_, err = c.Gui.SetViewOnTop(name)
	return err
}

func (c *Console) CloseElement(name string) {
	v, err := c.GetElement(name)
	if err != nil {
		return
	}
	if err = v.Close(); err != nil && err != gocui.ErrUnknownView {
		logrus.Error(err)
	}
}

func (c *Console) CloseElements(names ...string) {
	for _, name := range names {
		c.CloseElement(name)
	}
}

func (c *Console) doRun() error {
	defer c.Close()

	if hd, _ := os.LookupEnv("HARVESTER_DASHBOARD"); hd == "true" {
		c.SetManagerFunc(c.layoutDashboard)
	} else {
		c.SetManagerFunc(c.layoutInstall)
	}

	if err := setGlobalKeyBindings(c.Gui); err != nil {
		return err
	}

	if err := c.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func setGlobalKeyBindings(g *gocui.Gui) error {
	g.InputEsc = true
	if debug {
		if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
