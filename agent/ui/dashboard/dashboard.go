package dashboard

import (
	"context"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/intelops/qualityTrace/agent/ui/dashboard/components"
	"github.com/intelops/qualityTrace/agent/ui/dashboard/events"
	"github.com/intelops/qualityTrace/agent/ui/dashboard/models"
	"github.com/intelops/qualityTrace/agent/ui/dashboard/pages"
	"github.com/intelops/qualityTrace/agent/ui/dashboard/sensors"
	"github.com/intelops/qualityTrace/agent/ui/dashboard/styles"
	"github.com/rivo/tview"
)

type Dashboard struct{}

func startUptimeCounter(sensor sensors.Sensor) {
	ticker := time.NewTicker(time.Second)
	start := time.Now()
	go func() {
		for {
			select {
			case <-ticker.C:
				sensor.Emit(events.TimeChanged, time.Since(start).Round(time.Second))
			}
		}
	}()
}

func StartDashboard(ctx context.Context, environment models.EnvironmentInformation, sensor sensors.Sensor) error {
	app := tview.NewApplication()
	tview.Styles.PrimitiveBackgroundColor = styles.HeaderBackgroundColor
	renderScheduler := components.NewRenderScheduler(app)
	sensor.Emit(events.EnvironmentStart, environment)

	startUptimeCounter(sensor)

	router := NewRouter()
	router.AddAndSwitchToPage("home", pages.NewTestRunPage(renderScheduler, sensor))

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlC, tcell.KeyEsc:
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(router, true).SetFocus(router).Run(); err != nil {
		return fmt.Errorf("failed to start dashboard: %w", err)
	}

	return nil
}
