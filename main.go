package main

import "math/rand"

func main() {
	println(Dota2Name())
	//
}

func Dota2Name() string {
	dota2Names := []string{"Abaddon", "Alchemist", "Ancient Apparition", "Anti-Mage", "Axe", "Bane", "Batrider", "Beastmaster", "Bloodseeker", "Broodmother", "Centaur Warrunner", "Chaos Knight", "Chen", "Clinkz", "Clockwerk", "Crystal Maiden", "Dark Seer", "Dark Willow", "Dazzle", "Death Prophet", "Disruptor", "Doom", "DragonKnight", "Drow Ranger", "Earth Spirit", "Earthshaker", "Elder Titan", "Enchantress", "Enigma", "Faceless Void", "Grimstroke", "Gyrocopter", "Huskar", "Invoker", "Io", "Juggernaut", "Keeper of the Light", "Kunkka", "Legion Commander", "Leshrac", "Lich", "Lifestealer", "Lina", "Lion", "Lone Druid", "Luna", "Lycan", "Meepo", "Mirana", "Morphling", "Naga Siren", "Nature's Prophet", "Necrophos", "Night Stalker", "Nyx Assassin", "Ogre Magi", "Omniknight", "Oracle", "Outworld Devourer", "Pangolier", "Phantom Assassin", "Phantom Lancer", "Puck", "Pudge", "Pugna", "Queen of Pain", "Razor", "Riki", "Rubick", "Sand King", "Shadow Demon", "Shadow Fiend", "Shadow Shaman", "Silencer", "Skywrath Mage", "Slardar", "Slark", "Snapfire", "Sniper", "Spectre", "Spirit Breaker", "Storm Spirit", "Sven", "Techies", "Templar Assassin", "Terrorblade", "Tidehunter", "Timbersaw", "Tinker", "Tiny", "Treant Protector", "Troll Warlord", "Tusk", "Underlord", "Undying", "Ursa", "Vengeful Spirit", "Venomancer", "Viper", "Visage", "Warlock", "Weaver", "Windranger", "Winter Wyvern", "Witch Doctor", "Wraith King", "Zeus"}
	return dota2Names[rand.Intn(len(dota2Names))]
}
