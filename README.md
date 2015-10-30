# gtk-gravitysimulation
An exploration of golang and GTK2.

This project was an introduction to writing testable, robust go code. The main focus of the project was writing, testing, and tuning the physics engine code. The GTK2 visualization is unpolished as the visualizations main purpose was to get a look at the physics engine in action.

## Compilation
This project relies on [go-gtk](https://github.com/mattn/go-gtk/) for the GUI and thus one will need to run the following command before building.

```bash
go get github.com/mattn/go-gtk/gtk
```

I have run into issues with `go install`, however a simple `go build app.go` or `go run app.go` in this directory should suffice to create or launch the executable.

Dislaimer: This is a hobby project designed and tested on Debian Jessie, no guarantees are made on any other platforms compatibility.

## Program

The program simulates gravity on user-defined entities at a real-time scale. As a result the gravitational constant has been ballooned upwards so that watching the simulation is enjoyable. All other mathematical formula and constants other than the gravitational constant are unchanged.

In the images below you can see the black dots as entities of a given mass with larger entities having more mass. The red arrow indicated the velocity of an entity, and the blue arrow represent the acceleration of an entity. 

The canvas is a 640 pixel by 640 pixel grid with the origin at (320,320). There is a direct mapping between pixels and location such that x location 200 is pixel 520. The edges of the canvas are bounded such that entities reflect off of them with a bounding effect of losing velocity magnitude.

The bottom buttons control time in the simulation. Reset resets time to 0s. Each tick is 0.1s of real time. If Auto Update is depressed then a click will be triggered every time quanta signified by the slider.

The entities panel allows defining of all entity fields at time 0. Issuing a reset will take current values from the entities panel as entities in the simulation.

## Pictures
![simulation](https://cloud.githubusercontent.com/assets/5449328/10843762/11d705d0-7eb8-11e5-90b8-4e899bb34824.png)

![entitiespage](https://cloud.githubusercontent.com/assets/5449328/10843777/3719b1b2-7eb8-11e5-87dc-abbd05d49754.png)

## Example Values
* Planet orbiting a Star

| Mass | X-Pos | Y-Pos | X-Vel | Y-Vel | X-Acc | Y-Acc |
| ---: | ----: | ----: | ----: | ----: | ----: | ----: |
| 1000 | 0     | 0     | 0     | 0     | 0     | 0     |
| 3    | 200   | 0     | 0     | -60   | 0     | 0     |

* Several Planets orbiting a Star (Solar System)

| Mass | X-Pos | Y-Pos | X-Vel | Y-Vel | X-Acc | Y-Acc |
| ---: | ----: | ----: | ----: | ----: | ----: | ----: |
| 1000 | 0     | 0     | 0     | 0     | 0     | 0     |
| 10   | 200   | 0     | 0     | 60    | 0     | 0     |
| 2    | 0     | 100   | -80   | 0     | 0     | 0     |
| 3    | -150  | 0     | 0     | 65    | 0     | 0     |

* Binary Star system with orbiting Planet

| Mass | X-Pos | Y-Pos | X-Vel | Y-Vel | X-Acc | Y-Acc |
| ---: | ----: | ----: | ----: | ----: | ----: | ----: |
| 1000 | 100   | 0     | 40    | 0     | 0     | 0     |
| 1000 | -100  | 0     | -40   | 0     | 0     | 0     |
| 1    | 240   | 0     | 0     | -80   | 0     | 0     |
