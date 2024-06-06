package main

import (
	"fmt"
	"math/rand"
	"time"
)

var choice int // додала її як публічну бо вважається що вона буде використовуватись всюди
// але тільки для логіки швидкого вибору

type MainCharacter struct {
	Name         string
	WeaponStatus bool
	Alive        bool
	Luck         int

	YouLucky bool

	Pet
}

type Pet struct {
	PetName string
	PetSave bool
}

func main() {
	fmt.Printf("Гра почалась!\n")

	var character MainCharacter
	character.createCharacter()

	character.Chapter1()
	character.Chapter2()

	fmt.Printf("\nГра закінчена! Дякую, %v!\nОсь твої результати:", character.Name)
	fmt.Printf("\nРівень вдачі: %v", character.Luck)

	if character.Pet.PetSave {
		fmt.Printf("\n%v з ним усе добре!", character.PetName)
	} else {
		fmt.Printf("\n%v нажаль з ним сталась трагедія", character.PetName)
	}
	if character.WeaponStatus {
		fmt.Printf("\nВ тебе є меч!")
	} else {
		fmt.Printf("\nВ тебе немає меча.")
	}
}

func (c *MainCharacter) createCharacter() {
	c.WeaponStatus = false
	c.Alive = true
	c.Luck = 6

	fmt.Println("\nТи розплющуєш очі, а довкола тебе темрява, чутно шурхотіння листви, сов, під тобою хрустять гілки.\nЗгадуєш як опинився тут і з пам'яті дістаєш, що ти вигулював свою тваринку, свого кращого друга.\nЯк же його звали...\n\nВведіть ім'я улюбленця:")

	fmt.Scan(&c.Pet.PetName)

	fmt.Printf("\nТочно! %v!\nАле як тебе звати?\n\nВведіть своє ім'я:\n", c.Pet.PetName)

	fmt.Scan(&c.Name)
}

func (c *MainCharacter) CheckLuck() {
	rand.Seed(time.Now().UnixNano()) // знайшла в інеті, чомусь іноді підкреслює жовтим
	randomLuck := rand.Intn(4)

	if c.Luck == 6 || randomLuck+c.Luck >= 6 {
		c.YouLucky = true
	} else {
		c.YouLucky = false
	}
}

func (c *MainCharacter) Chapter1() {
	foundSword := false

	fmt.Printf("\n%s згадалось, що гуляли з %s і тоді щось його налякало і він шмигнув у ліс.\nКоли ти намагався його наздогнати, то впав і покотився в прірву.", c.Name, c.PetName)
	fmt.Printf("\nТреба знайти %s, він наляканий і зовсім один в цьому страшному лісі. Рушай!\n\nОбери шлях:\n 1 - спробувати залізти назад у гору\n 2 - піти в гущу лісу\n", c.PetName)

	fmt.Scan(&choice)

	for {
		switch choice {
		case 1:
			fmt.Println("\nТи намагаєшся залізти в гору, ти майже доліз до самого верху, але земля занадто слизька.\nТи падаєш! Доведеться йти в ліс.")
			c.Luck -= 2

		case 2:
			fmt.Println("\nТи обтрушуєш речі і прямуєш прямо в чащу.")

		default:
			fmt.Printf("\nНе зрозумів...\nОбери шлях:\n 1 - спробувати залізти назад у гору\n 2 - піти в гущу лісу\n")
			fmt.Scan(&choice)
			continue
		}
		break
	}

	fmt.Println("\nТи рухаєшся вперед, але відчуття якісь дивні.\nРаптово ти помічаєш вогник, який наче рухається до тебе.")
	fmt.Printf("\nЩо будеш робити:\n 1 - піти на зустріч до вогника\n 2 - спробувати сховатись\n")

	fmt.Scan(&choice)

	for {
		switch choice {
		case 1:
			fmt.Println("\nВогник наче хоче вказати тобі дорогу і ти йдеш до нього\n Через якийсь час ти опиняєшся біля каменю з мечем")
			foundSword = true

		case 2:
			fmt.Println("\nТи ховаєшся за деревом і потім намагаєшся рухатись в іншу сторону від вогника\n Відчуття, що ти повністю заблукав.")

		default:
			fmt.Printf("\nНе зрозумів...\nЩо будеш робити:\n 1 - піти на зустріч до вогника\n 2 - спробувати сховатись\n")
			fmt.Scan(&choice)
			continue
		}
		break
	}

	if foundSword {
		c.CheckLuck()
		fmt.Printf("\nТи дивишся на цей меч і це нагадує тобі казки з дитинства.\n")
		fmt.Printf("\nЩо будеш робити:\n 1 - спробувати витягнути меч\n 2 - йти далі, найважливіше зараз знайти %v\n", c.PetName)
		fmt.Scan(&choice)

		for {
			switch choice {
			case 1:
				if c.YouLucky {
					fmt.Println("\nМеч тепер твій! Час продовжувати пошуки.")
					c.WeaponStatus = true
				} else {
					fmt.Println("\nТи не зміг витягнути меч. Час продовжувати пошуки.")
					c.Luck -= 2
				}

			case 2:
				fmt.Println("\nТи ще якийсь час блукаєш лісом, до поки не знаходиш печеру.")

			default:
				fmt.Printf("\nЩе раз:\n 1 - спробувати витягнути меч\n 2 - йти далі, найважливіше зараз знайти %v\n", c.PetName)
				fmt.Scan(&choice)
				continue
			}
			break
		}
	} else {
		fmt.Println("\nЩо далі:\n 1 - спробувати повернутсь до вогника\n 2 - йти далі")
		fmt.Scan(&choice)

		for {
			switch choice {
			case 1:
				fmt.Printf("\nНічого не вийшло, він зник.\nТи тільки більше втомився, але треба далі шукати %s.\nТи блукаєш до поки не знаходиш печеру.", c.PetName)
				c.Luck -= 2

			case 2:
				fmt.Println("\nТи впевнено крокуєш вибраним шляхом, який приводить тебе до печери.")

			default:
				fmt.Printf("\nНе зрозумів...\nЩо далі:\n 1 - спробувати повернутсь до вогника\n 2 - йти далі\n")
				fmt.Scan(&choice)
				continue
			}
			break
		}
	}
}

func (c MainCharacter) Chapter2() {
	fmt.Printf("\nТи довго блукав і яке щастя що знайшов печеру.\nПоруч з нею ти бачеш кроки %v", c.PetName)
	fmt.Println("\nЩо робитимеш:\n 1 - втікти\n 2 - зайти в печеру")

	fmt.Scan(&choice)
	inCave := false

	for {
		switch choice {
		case 1:
			fmt.Printf("\n%s пішов так і не знайшовши %v\nБільше в тебе ніколи не було друзів.\n", c.Name, c.PetName)
		case 2:
			fmt.Println("\nТи заходиш в печеру.")
			inCave = true

		default:
			fmt.Printf("\nШо?:\n 1 - втікти\n 2 - зайти в печеру")
			fmt.Scan(&choice)
			continue
		}
		break
	}

	if inCave {
		fmt.Printf("\nТи бачиш, що %v оточили привиди!\nЩО Ж РОБИТИ?", c.PetName)
		fmt.Printf("\n 1 - одразу кинутись на допомогу\n 2 - спробувати придумати план\n")

		fmt.Scan(&choice)

		c.CheckLuck()

		for {
			switch choice {
			case 1:
				if c.WeaponStatus {
					fmt.Printf("\nТи зводиш меч і нападаєш на привидів.\nВиявляється вони бояться заліза і тобі вдається їх розлякати.\nТи виглядав дуже круто!\n%v підбігає до тебе. Він взахваті і безмежно радий тебе бачити!\nВи разом покидаєте печеру і більше ніколи не повертаєтесь в цей ліс.\n", c.PetName)
					c.Pet.PetSave = true
				} else if c.YouLucky {
					fmt.Printf("\nТи кидаєшся на привидів з єдиною думкою врятувати друга.\n%v слідує твоєму прикладу.\nПривиди не очікували такої зухвалості і йдуть геть.\nВи разом і з %v після цього жили довго і щасливо.\n", c.PetName, c.PetName)
					c.Pet.PetSave = true
				} else {
					fmt.Printf("\nТи кидаєшся на привидів з єдиною думкою врятувати друга.\nВони звертають на тебе увагу і вбивають тебе.\nНе сумуй %s, це всього-на-всього вигадана історія.\nТи був хоробрим і зробив все, що міг.\n", c.Name)
					c.Pet.PetSave = false
					c.Alive = false
				}

			case 2:
				if c.WeaponStatus && c.YouLucky {
					fmt.Printf("\nТи дуже довго думав. Привиди тебе помітили.\nНажаль тобі не вдалось врятуватись, але ти врятував %s\n", c.PetName)
					c.Pet.PetSave = true
				} else if !c.WeaponStatus && c.YouLucky {
					fmt.Println("\nТи подивився навколо, підібрав каміння і кинув в привидів.\nВони відволіклись. Треба погукати друга!\nВведіть його ім'я:")
					var name string
					fmt.Scan(&name)
					if name == c.PetName {
						fmt.Printf("\nТи кличеш %v і він чує і біжить до тебе.\nПоки привиди були зайняті камінням, вам вдалося втекти.\nІ після цього ти і %v жили довго і щасливо\n", c.PetName, c.PetName)
						c.Pet.PetSave = true
					} else {
						fmt.Printf("\nТи вигукнув щось нерозбірливе, і це привернуло увагу тільки привидів. Всі загинули.\n")
						c.Pet.PetSave = false
						c.Alive = false
					}
				} else {
					fmt.Printf("\nТобі дуже не щастить, привиди всіх вбили.\n")
					c.Pet.PetSave = false
					c.Alive = false
				}

			default:
				fmt.Printf("\nБудь-ласка... просто обери нормально\n 1 - одразу кинутись на допомогу\n 2 - спробувати придумати план\n")
				fmt.Scan(&choice)
				continue
			}
			break
		}
	}
}
