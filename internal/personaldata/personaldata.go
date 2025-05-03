package personaldata

import "fmt"

type Personal struct {
	//name - имя пользователя
	Name string
	//Weight - вес пользователя
	Weight float64
	//Height- рост пользователя
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)

}
