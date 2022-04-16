package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"strings"
)

type Hero struct {
	CurrentHP                 int
	MaxHP                     int
	CurrentEXP                int
	MaxEXP                    int
	Money                     int
	HealingPotions            int
	Damage                    int
	Level                     int
	HealingPotionEffect       int
	HealingPotionPrice        int
	BetterSwordPrice          int
	BetterArmorPrice          int
	HealingPotionBuyedCounter int
	BetterSwordBuyedCounter   int
	BetterArmorBuyedCounter   int
}

func (h Hero) String() string {
	//Level: Level
	//HP: CurrentHP/MaxHP
	//EXP: EXP/MaxEXP
	//Money: Money
	//Damage: Damage
	return fmt.Sprintf(
		"Level: %v\nHP: %v/%v\nEXP: %v/%v\nMoney: %v\nDamage: %v",
		h.Level, h.CurrentHP, h.MaxHP,
		h.CurrentEXP, h.MaxEXP, h.Money, h.Damage)
}

func (h *Hero) CheckAndLevelUp() {
	for h.CurrentEXP >= h.MaxEXP {
		h.Level += 1
		h.CurrentEXP -= h.MaxEXP
		h.MaxHP = 10 + int(float64(h.Level*h.Level)*0.02) + h.BetterArmorBuyedCounter
		h.Damage = 1 + h.Level + int(math.Sqrt(float64(h.BetterSwordBuyedCounter*h.BetterSwordBuyedCounter*h.BetterSwordBuyedCounter))*0.33)
		h.CurrentHP = h.MaxHP
		h.HealingPotionEffect = int(float64(h.MaxHP) * float64(0.5))
		h.HealingPotionPrice += 1
		h.BetterArmorPrice += 1
		h.BetterSwordPrice += 1
		h.MaxEXP = 10 + int(math.Sqrt(float64(h.MaxEXP*h.MaxEXP*h.MaxEXP))*0.3)
	}
	return
}

type Enemy struct {
	CurrentHP           int
	MaxHP               int
	Damage              int
	EXPBounty           int
	MoneyBounty         int
	HealingPotionBounty int
}

func (e Enemy) String() string {
	//HP: CurrentHP/MaxHP
	//Damage: Damage
	return fmt.Sprintf("HP: %v/%v\nDamage: %v",
		e.CurrentHP, e.MaxHP, e.Damage)
}

func GetEnemy(level int) Enemy {
	enemy := Enemy{1 + int(float64(level*level)*0.0015), 1 + int(float64(level*level)*0.0015), 1 + int(float64(level*level)*0.02), int(float64(level) * 0.3), int(float64(level) * 0.3), 0}
	return enemy
}

func Run(inChan, outChan chan string, enemiesToKill int, winnerFlag string) {
}

func ServeConnection(c net.Conn, enemiesToKill int, winnerFlag string) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()

	actionStrings := make(map[string]string)
	actionStrings["welcome"] = `Welcome to text-based Roguelike PPC Task!
To win, you need to kill 1000 enemy!
(f/F) Find new enemy to fight
(s/S) Examine yourself
(h/H) Heal yourself
(b/B) Go to shop
`
	actionStrings["stats"] = `Your stats:
Level: %v
HP: %v/%v
EXP: %v/%v
Money: %v
Healing Potions: %v
Damage: %v
What are you going to do next?
(f/F) Find new enemy to fight
(s/S) Examine yourself
(h/H) Heal yourself
(b/B) Go to shop
`
	actionStrings["fightEnemyFound"] = `You joined the fight!
Enemy stats:
HP: %v/%v
Damage: %v
Your stats:
Level: %v
HP: %v/%v
EXP: %v/%v
Healing Potions: %v
Damage: %v
(a/A) Attack!
(h/H) Heal yourself!
(r/R) Run away!
`
	actionStrings["fightAttackEnemy"] = `You hit the enemy!
The enemy hit you too!
Enemy stats:
HP: %v/%v
Damage: %v
Your stats:
Level: %v
HP: %v/%v
EXP: %v/%v
Healing Potions: %v
Damage: %v
(a/A) Attack!
(h/H) Heal yourself!
(r/R) Run away!
`
	actionStrings["fightKillEnemy"] = `You hit the enemy!
The enemy died from your hit!
You received %v EXP
You received %v Healing Potion
You received %v Coin
What are you going to do next?
(f/F) Find new enemy to fight
(s/S) Examine yourself
(h/H) Heal yourself
(b/B) Go to shop
`
	actionStrings["fightHeroDeath"] = `You hit the enemy!
The enemy hit you too!
And after hit you fell a hero's death...
Game Over!
`
	actionStrings["fightDrinkPotion"] = `You drink a healing potion!
The enemy hit you too!
Enemy stats:
HP: %v/%v
Damage: %v
Your stats:
Level: %v
HP: %v/%v
EXP: %v/%v
Healing Potions: %v
Damage: %v
(a/A) Attack!
(h/H) Heal yourself!
(r/R) Run away!
`
	actionStrings["fightRunAway"] = `You running away!
But at the last moment before you run away, the enemy hits you!
Your stats:
Level: %v
HP: %v/%v
EXP: %v/%v
Healing Potions: %v
Damage: %v
What are you going to do next?
(f/F) Find new enemy to fight
(s/S) Examine yourself
(h/H) Heal yourself
(b/B) Go to shop
`
	actionStrings["shopEnter"] = `You entered the shop.
What do you want to buy?
(h/H) Healing Potion - %v coins
(s/S) Better Sword (+%v damage) - %v coins
(a/A) Better Armor (+1 HP) - %v coins
(q/Q) Nothing. Go away!
You have %v coins
`
	actionStrings["shopLeft"] = `You left the store.
What are you going to do next?
(f/F) Find new enemy to fight
(s/S) Examine yourself
(h/H) Heal yourself
(b/B) Go to shop
`
	actionStrings["shopBuyHealingPotion"] = `You bought a healing potion
What do you want to buy next?
(h/H) Healing Potion - %v coins
(s/S) Better Sword (+%v damage) - %v coins
(a/A) Better Armor (+1 HP) - %v coins
(q/Q) Nothing. Go away
You have %v coins
`
	actionStrings["shopBuySword"] = `You bought a better sword
What do you want to buy next?
(h/H) Healing Potion - %v coins
(s/S) Better Sword (+%v damage) - %v coins
(a/A) Better Armor (+1 HP) - %v coins
(q/Q) Nothing. Go away
You have %v coins
`
	actionStrings["shopBuyArmor"] = `You bought a better armor
What do you want to buy next?
(h/H) Healing Potion - %v coins
(s/S) Better Sword (+%v damage) - %v coins
(a/A) Better Armor (+1 HP) - %v coins
(q/Q) Nothing. Go away
You have %v coins
`
	actionStrings["shopNoMoney"] = `It's yours my friend. As long as you have enough rupees.
What do you want to buy?
(h/H) Healing Potion - %v coins
(s/S) Better Sword (+%v damage) - %v coins
(a/A) Better Armor (+1 HP) - %v coins
(q/Q) Nothing. Go away
You have %v coins
`
	actionStrings["healYourself"] = `You drinked a healing potion (%v left)
You restored %v hp
Your current status: %v/%v
What are you going to do next?
(f/F) Find new enemy to fight
(s/S) Examine yourself
(h/H) Heal yourself
(b/B) Go to shop
`
	actionStrings["win"] = `Congratulions! You killed all monsters!
Your win flag is: %v
`
	hero := Hero{10, 10, 0, 10, 10, 10, 1, 1, 3, 1, 1, 1, 0, 0, 0}

	enemyLevel := 1
	enemy := GetEnemy(enemyLevel)

	state := "nothing"

	c.Write([]byte(actionStrings["welcome"]))

	for {
		b := make([]byte, 4)
		_, err := c.Read(b)
		if err != nil {
			return
		}
		action := string(b)
		//log.Println(enemyLevel, action)
		if state == "nothing" && (strings.HasPrefix(action, "s") || strings.HasPrefix(action, "S")) {
			c.Write([]byte(fmt.Sprintf(
				actionStrings["stats"], hero.Level,
				hero.CurrentHP, hero.MaxHP, hero.CurrentEXP,
				hero.MaxEXP, hero.Money, hero.HealingPotions,
				hero.Damage)))
		}
		if state == "nothing" && (strings.HasPrefix(action, "f") || strings.HasPrefix(action, "F")) {
			state = "fight"
			c.Write([]byte(fmt.Sprintf(actionStrings["fightEnemyFound"], enemy.CurrentHP,
				enemy.MaxHP, enemy.Damage, hero.Level,
				hero.CurrentHP, hero.MaxHP, hero.CurrentEXP,
				hero.MaxEXP, hero.HealingPotions,
				hero.Damage)))
		}
		if state == "nothing" && (strings.HasPrefix(action, "b") || strings.HasPrefix(action, "B")) {
			state = "shop"
			c.Write([]byte(fmt.Sprintf(actionStrings["shopEnter"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
		}
		if state == "nothing" && (strings.HasPrefix(action, "h") || strings.HasPrefix(action, "H")) {
			hero.CurrentHP += hero.HealingPotionEffect
			if hero.CurrentHP > hero.MaxHP {
				hero.CurrentHP = hero.MaxHP
			}
			hero.HealingPotions -= 1
			c.Write([]byte(fmt.Sprintf(actionStrings["healYourself"], hero.HealingPotions, hero.HealingPotionEffect, hero.CurrentHP, hero.MaxHP)))
		}
		if state == "fight" && (strings.HasPrefix(action, "a") || strings.HasPrefix(action, "A")) {
			if hero.Damage >= enemy.CurrentHP {
				hero.HealingPotions += enemy.HealingPotionBounty
				hero.Money += enemy.MoneyBounty
				hero.CurrentEXP += enemy.EXPBounty
				hero.CheckAndLevelUp()
				enemyLevel += 1
				if enemyLevel > enemiesToKill {
					c.Write([]byte(fmt.Sprintf(actionStrings["win"], winnerFlag)))
					return
				}
				enemy = GetEnemy(enemyLevel)

				state = "nothing"
				c.Write([]byte(fmt.Sprintf(actionStrings["fightKillEnemy"], enemy.EXPBounty, enemy.HealingPotionBounty, enemy.MoneyBounty)))
			} else if enemy.Damage < hero.CurrentHP {
				enemy.CurrentHP -= hero.Damage
				hero.CurrentHP -= enemy.Damage
				c.Write([]byte(fmt.Sprintf(actionStrings["fightAttackEnemy"], enemy.CurrentHP, enemy.MaxHP, enemy.Damage, hero.Level, hero.CurrentHP, hero.MaxHP, hero.CurrentEXP, hero.MaxEXP, hero.HealingPotions, hero.Damage)))
			} else {
				c.Write([]byte(fmt.Sprintf(actionStrings["fightHeroDeath"])))
			}
		}
		if state == "fight" && (strings.HasPrefix(action, "h") || strings.HasPrefix(action, "H")) {
			if enemy.Damage < (hero.CurrentHP + hero.HealingPotionEffect) {
				hero.CurrentHP += hero.HealingPotionEffect
				if hero.CurrentHP > hero.MaxHP {
					hero.CurrentHP = hero.MaxHP
				}
				hero.CurrentHP -= enemy.Damage
				hero.HealingPotions -= 1
				c.Write([]byte(fmt.Sprintf(actionStrings["fightDrinkPotion"], enemy.CurrentHP, enemy.MaxHP, enemy.Damage, hero.Level, hero.CurrentHP, hero.MaxHP, hero.CurrentEXP, hero.MaxEXP, hero.HealingPotions, hero.Damage)))
			} else {
				c.Write([]byte(fmt.Sprintf(actionStrings["fightHeroDeath"])))
			}
		}
		if state == "fight" && (strings.HasPrefix(action, "r") || strings.HasPrefix(action, "R")) {
			if enemy.Damage < hero.CurrentHP {
				hero.CurrentHP -= enemy.Damage
				state = "nothing"
				c.Write([]byte(fmt.Sprintf(actionStrings["fightRunAway"], hero.Level, hero.CurrentHP, hero.MaxHP, hero.CurrentEXP, hero.MaxEXP, hero.HealingPotions, hero.Damage)))
			} else {
				c.Write([]byte(fmt.Sprintf(actionStrings["fightHeroDeath"])))
			}
		}
		if state == "shop" && (strings.HasPrefix(action, "h") || strings.HasPrefix(action, "H")) {
			if hero.Money >= hero.HealingPotionPrice {
				hero.Money -= hero.HealingPotionPrice
				hero.HealingPotions += 1
				c.Write([]byte(fmt.Sprintf(actionStrings["shopBuyHealingPotion"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
			} else {
				c.Write([]byte(fmt.Sprintf(actionStrings["shopNoMoney"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
			}
		}
		if state == "shop" && (strings.HasPrefix(action, "s") || strings.HasPrefix(action, "S")) {
			if hero.Money >= hero.BetterSwordPrice {
				hero.Money -= hero.BetterSwordPrice
				hero.BetterSwordBuyedCounter += 1
				hero.Damage += 1 + int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33)
				c.Write([]byte(fmt.Sprintf(actionStrings["shopBuySword"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
			} else {
				c.Write([]byte(fmt.Sprintf(actionStrings["shopNoMoney"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
			}
		}
		if state == "shop" && (strings.HasPrefix(action, "a") || strings.HasPrefix(action, "A")) {
			if hero.Money >= hero.BetterArmorPrice {
				hero.Money -= hero.BetterArmorPrice
				hero.BetterArmorBuyedCounter += 1
				hero.MaxHP += 1
				c.Write([]byte(fmt.Sprintf(actionStrings["shopBuyArmor"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
			} else {
				c.Write([]byte(fmt.Sprintf(actionStrings["shopNoMoney"], hero.HealingPotionPrice, 1+int(math.Sqrt(float64(hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter*hero.BetterSwordBuyedCounter))*0.33), hero.BetterSwordPrice, hero.BetterArmorPrice, hero.Money)))
			}
		}
		if state == "shop" && (strings.HasPrefix(action, "q") || strings.HasPrefix(action, "Q")) {
			state = "nothing"
			c.Write([]byte(fmt.Sprintf(actionStrings["shopLeft"])))
		}
	}
	c.Close()
	return
}

func ServeReader(inChan chan string, r io.Reader) {
	reader := bufio.NewReader(r)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		select {
		case inChan <- strings.TrimSpace(s):
		}
	}
}

func main() {
	mode := flag.String("mode", "local", "local/remote\nChoose to start a remote or a local game")
	addr := flag.String("addr", "0.0.0.0:12345", "Choose a addr to start a remote game")
	enemiesToKill := flag.Int("enemiesToKill", 1000, "How much you need to kill enemies to obtain a flag")
	winnerFlag := flag.String("flag", "YetiCTF{H0P3Y0UD1DN7P14Y8YY0Ur531FFFFFFF}", "CTF Flag")
	flag.Parse()
	if *mode == "local" {
		inChan := make(chan string, 1000)
		outChan := make(chan string, 1000)
		go Run(inChan, outChan, *enemiesToKill, *winnerFlag)
		go ServeReader(inChan, os.Stdin)
		for {
			select {
			case msg := <-outChan:
				fmt.Printf("\n%v", msg)
			default:
			}
		}
	} else if *mode == "remote" {
		log.Println("Launching server")
		ln, err := net.Listen("tcp", *addr)
		if err != nil {
			log.Fatalf("Error while reserving port: %v\n", err)
		}
		defer ln.Close()
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("Error while starting listening connection: %v\n", err)
			}
			go ServeConnection(conn, *enemiesToKill, *winnerFlag)
		}
	}
}
