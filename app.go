package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/tkajder/gravitysimulator/physics"
	"github.com/tkajder/gravitysimulator/utils"
	"log"
	"math"
	"strconv"
	"time"
)

// Global drawing area pieces
var drawingarea *gtk.DrawingArea
var drawable *gdk.Drawable
var blackgc *gdk.GC
var redgc *gdk.GC
var bluegc *gdk.GC

// Size of the drawable area
const width int = 640
const height int = 640

// Entity limits
const entityfields int = 7
const entitylimit int = 18

// How much to damp velocity on colliding with the outside walls
const damping float64 = 0.7

// Parse the entities from user input and return a slice of valid entities
func initentities(entries [][]*gtk.Entry) []*physics.Entity {
	entities := make([]*physics.Entity, 0)

	for rownum, row := range entries {
		// Check if row is empty, skip if true
		empty := true
		for _, entry := range row {
			if entry != nil && entry.GetText() != "" {
				empty = false
			}
		}
		if empty {
			continue
		}

		// Parse mass
		mass, err := strconv.ParseFloat(row[0].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse mass in row %v: %v - skipping", rownum, err)
			continue
		}

		// Parse x position
		posx, err := strconv.ParseFloat(row[1].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse x position in row %v: %v - skipping", rownum, err)
			continue
		}

		// Parse y position
		posy, err := strconv.ParseFloat(row[2].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse y position in row %v: %v - skipping", rownum, err)
			continue
		}

		// Parse x velocity
		velx, err := strconv.ParseFloat(row[3].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse x velocity in row %v: %v - skipping", rownum, err)
			continue
		}

		// Parse y velocity
		vely, err := strconv.ParseFloat(row[4].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse y velocity in row %v: %v - skipping", rownum, err)
			continue
		}

		// Parse x acceleration
		accelx, err := strconv.ParseFloat(row[5].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse x acceleration in row %v: %v - skipping", rownum, err)
			continue
		}

		// Parse y acceleration
		accely, err := strconv.ParseFloat(row[6].GetText(), 64)
		if err != nil {
			log.Printf("Could not parse y acceleration in row %v: %v - skipping", rownum, err)
			continue
		}

		// If all entity fields parsed, append the new entity
		entities = append(entities, physics.NewEntity(mass, posx, posy, velx, vely, accelx, accely))
	}

	return entities
}

// Update all entities acceleration then velocity and position and kick off a draw
func updateentities(entities []*physics.Entity) {
	// Update all graviational accelerations
	for _, e := range entities {
		e.UpdateGravitationalAcceleration(entities)
	}

	// Once all gravitational accelerations are set move positions and velocities
	for _, e := range entities {
		bound(e)
		e.Update(0.01)
	}

	drawingarea.QueueDraw()
}

// Check that the entity does not leave the drawing area,
func bound(e *physics.Entity) {
	halfwidth := float64(width / 2)
	halfheight := float64(height / 2)
	if e.Position.X < -halfwidth && e.Velocity.X < 0 {
		e.Velocity = e.Velocity.InvertX()
		e.Velocity = e.Velocity.Scalarmul(damping)
	}
	if e.Position.X > halfwidth && e.Velocity.X > 0 {
		e.Velocity = e.Velocity.InvertX()
		e.Velocity = e.Velocity.Scalarmul(damping)
	}
	if e.Position.Y < -halfheight && e.Velocity.Y < 0 {
		e.Velocity = e.Velocity.InvertY()
		e.Velocity = e.Velocity.Scalarmul(damping)
	}
	if e.Position.Y > halfheight && e.Velocity.Y > 0 {
		e.Velocity = e.Velocity.InvertY()
		e.Velocity = e.Velocity.Scalarmul(damping)
	}
}

// Draw all entities with black for position, red for velocity, and blue for acceleration
func drawentities(entities []*physics.Entity) {
	for _, entity := range entities {
		drawposition(entity, drawable, blackgc)
		drawvelocity(entity, drawable, redgc)
		drawacceleration(entity, drawable, bluegc)
	}
}

func drawposition(entity *physics.Entity, drawable *gdk.Drawable, gc *gdk.GC) {
	radius := math.Sqrt(entity.Mass)
	startx := utils.RoundInt(centerfloatx(entity.Position.X) - (radius / 2))
	starty := utils.RoundInt(centerfloaty(entity.Position.Y) - (radius / 2))

	// Draw at least 1 pixel for position
	if radius < 1 {
		radius = 1
	}

	drawable.DrawArc(gc, true, startx, starty, int(utils.Round(radius)), int(radius), 0, 360*64)
}

func drawvelocity(entity *physics.Entity, drawable *gdk.Drawable, gc *gdk.GC) {
	startx := utils.RoundInt(entity.Position.X)
	starty := utils.RoundInt(entity.Position.Y)
	endx := utils.RoundInt(entity.Position.X + entity.Velocity.X)
	endy := utils.RoundInt(entity.Position.Y + entity.Velocity.Y)
	drawable.DrawLine(gc, centerx(startx), centery(starty), centerx(endx), centery(endy))
}

func drawacceleration(entity *physics.Entity, drawable *gdk.Drawable, gc *gdk.GC) {
	startx := utils.RoundInt(entity.Position.X)
	starty := utils.RoundInt(entity.Position.Y)
	endx := utils.RoundInt(entity.Position.X + entity.Acceleration.X)
	endy := utils.RoundInt(entity.Position.Y + entity.Acceleration.Y)
	drawable.DrawLine(gc, centerx(startx), centery(starty), centerx(endx), centery(endy))
}

func centerx(x int) int {
	return (width / 2) + x
}

func centery(y int) int {
	return (height / 2) + y
}

func centerfloatx(x float64) float64 {
	return float64(width/2) + x
}

func centerfloaty(y float64) float64 {
	return float64(height/2) + y
}

func main() {
	var autoupdating bool = false
	var autoticker *time.Ticker
	var entries [][]*gtk.Entry = make([][]*gtk.Entry, entitylimit)
	for i := 0; i < entitylimit; i++ {
		entries[i] = make([]*gtk.Entry, entityfields)
	}
	var entities []*physics.Entity = initentities(entries)

	// Initialize gtk
	gtk.Init(nil)

	// WINDOW
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Gravity Visualization")

	// Connect top window closing to gtk main loop closing
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	// TOP VERTICAL BOX
	topvbox := gtk.NewVBox(false, 1)

	// NOTEBOOK FOR TABS
	notebook := gtk.NewNotebook()

	// DRAWING AREA VERTICAL BOX
	davbox := gtk.NewVBox(false, 1)

	// DRAWING AREA
	drawingarea = gtk.NewDrawingArea()
	drawingarea.SetSizeRequest(width, height)
	drawingarea.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	drawingarea.Connect("expose_event", func() {
		drawentities(entities)
	})
	davbox.PackStart(drawingarea, true, true, 0)

	// TICK SPEED SLIDER
	ticksliderhbox := gtk.NewHBox(false, 1)

	ticksliderlabel := gtk.NewLabel("Time between ticks (ms)")
	ticksliderhbox.Add(ticksliderlabel)

	tickslider := gtk.NewHScaleWithRange(1, 1000, 100)
	// Default value of 10 ms
	tickslider.SetValue(10)
	ticksliderhbox.Add(tickslider)
	davbox.Add(ticksliderhbox)

	// BUTTONS
	buttons := gtk.NewHBox(false, 1)

	// RESET MENU ITEM
	resetbutton := gtk.NewButtonWithLabel("Reset")
	resetbutton.Clicked(func() {
		entities = initentities(entries)
		drawingarea.QueueDraw()
	})
	buttons.Add(resetbutton)

	// TICK MENU ITEM
	tickbutton := gtk.NewButtonWithLabel("Tick")
	tickbutton.Clicked(func() {
		updateentities(entities)
	})
	buttons.Add(tickbutton)

	// AUTOUPDATE MENU ITEM
	autotickbutton := gtk.NewToggleButtonWithLabel("AutoUpdate")
	autotickbutton.Clicked(func() {
		// Stop the previous ticker if it exists
		if autoticker != nil {
			autoticker.Stop()
		}

		if autoupdating {
			// Toggle autoupdating state
			autoupdating = false
		} else {
			// Start the ticker
			autoticker = time.NewTicker(time.Duration(tickslider.GetValue()) * time.Millisecond)

			// Spawn a goroutine that will run update entities every tick
			go func() {
				for _ = range autoticker.C {
					updateentities(entities)
				}
			}()

			// Toggle autoupdating state
			autoupdating = true
		}
	})
	buttons.Add(autotickbutton)
	davbox.Add(buttons)

	notebook.AppendPage(davbox, gtk.NewLabel("Simulation"))

	// INITIALIZE PANEL
	entitiesvbox := gtk.NewVBox(false, 1)

	// INITIALIZE LABELS FOR TABLE
	titles := gtk.NewHBox(false, 1)
	titles.Add(gtk.NewLabel("Mass"))
	titles.Add(gtk.NewLabel("X-Pos"))
	titles.Add(gtk.NewLabel("Y-Pos"))
	titles.Add(gtk.NewLabel("X-Vel"))
	titles.Add(gtk.NewLabel("Y-Vel"))
	titles.Add(gtk.NewLabel("X-Acc"))
	titles.Add(gtk.NewLabel("Y-Acc"))
	entitiesvbox.Add(titles)

	// INITIALIZE ENTRIES IN ROWS FOR TABLE
	for row := 0; row < entitylimit; row++ {
		rowbox := gtk.NewHBox(false, 1)
		for col := 0; col < entityfields; col++ {
			textfield := gtk.NewEntry()
			// Hold reference to text field in entries 2d array
			entries[row][col] = textfield
			rowbox.Add(textfield)
		}
		entitiesvbox.Add(rowbox)
	}

	// CLEAR ENTITIES BUTTON
	clearentitiesbutton := gtk.NewButtonWithLabel("Clear Entries")
	clearentitiesbutton.Clicked(func() {
		for row := 0; row < entitylimit; row++ {
			for col := 0; col < entityfields; col++ {
				entries[row][col].SetText("")
			}
		}
	})
	entitiesvbox.Add(clearentitiesbutton)

	// Limit the size of the entitiesvbox and add to notebook
	entitiesvbox.SetSizeRequest(width, height)
	notebook.AppendPage(entitiesvbox, gtk.NewLabel("Entities"))

	// FINISH PACKING COMPONENTS
	topvbox.PackStart(notebook, false, false, 0)

	// FINISH PACKING WINDOW
	window.Add(topvbox)

	// Show the GUI
	window.ShowAll()

	// Grab the drawable and initialize graphics context now that they are initialized
	drawable = drawingarea.GetWindow().GetDrawable()
	blackgc = gdk.NewGC(drawable)
	redgc = gdk.NewGC(drawable)
	redgc.SetRgbFgColor(gdk.NewColorRGB(255, 0, 0))
	bluegc = gdk.NewGC(drawable)
	bluegc.SetRgbFgColor(gdk.NewColorRGB(0, 0, 255))

	gtk.Main()
}
