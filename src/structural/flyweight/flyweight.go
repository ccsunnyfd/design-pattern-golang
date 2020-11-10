package main

import "fmt"

// dress flyweight interface
type dress interface {
	getColor() string
}

// terroristDress concrete flyweight object
type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

// counterTerroristDress concrete flyweight object
type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

// dressFactory flyweight factory
const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

// player
type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

func (p *player) descPlayer() {
	fmt.Printf("dress: %s, playerType: %s, lat: %d, long: %d\n", p.dress.getColor(), p.playerType, p.lat, p.long)
}

// game
type game struct {
	playerList []*player
}

func newGame() *game {
	return &game{
		playerList: make([]*player, 0, 10),
	}
}

func (g *game) addTerrorist(lat, long int) {
	player := newPlayer("terrorist", TerroristDressType)
	player.lat = lat
	player.long = long
	g.playerList = append(g.playerList, player)
}

func (g *game) addCounterTerrorist(lat, long int) {
	player := newPlayer("counterTerrorist", CounterTerroristDressType)
	player.lat = lat
	player.long = long
	g.playerList = append(g.playerList, player)
}

// main
func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(1, 2)
	game.addTerrorist(3, 4)
	game.addTerrorist(5, 6)
	game.addTerrorist(7, 8)

	// //Add CounterTerrorist
	game.addCounterTerrorist(2, 1)
	game.addCounterTerrorist(4, 3)
	game.addCounterTerrorist(6, 5)

	// for dressType, dress := range dressFactoryInstance.dressMap {
	// 	fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	// }
	for _, player := range game.playerList {
		player.descPlayer()
	}
}
