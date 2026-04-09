package jedlik

import "fmt"

func (c *Car) Drive() {
    remaining := c.battery - c.batteryDrain
    if remaining < 0 {
        return
    }
    c.battery = remaining
    c.distance += c.speed
}

func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
    maxDrives := c.battery / c.batteryDrain
    maxDistance := maxDrives * c.speed
    return maxDistance >= trackDistance
}
