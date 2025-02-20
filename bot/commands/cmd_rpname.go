package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Southclaws/cj/types"
)

var rpfirsts = []string{
	"CJ",
	"O.G.",
	"SAMP",
	"bay",
	"bone",
	"bulgarian",
	"capital",
	"carl",
	"evolve",
	"gay",
	"german",
	"god",
	"godfather",
	"grand",
	"halal",
	"infinity",
	"las",
	"leaked",
	"los",
	"next",
	"one",
	"payday",
	"pisd",
	"pure",
	"red",
	"role",
	"san",
	"scavenge",
	"sexy",
	"texas",
	"kungkingkang",
	"mengontol",
	"misebah",
	"Nusantara Project",
	"Corvus",
	"Mega-Chan",
	"Wadin",
	"Wagyu Adinarta",
}
var rpseconds = []string{
	"SAMP",
	"andreas",
	"area",
	"christian",
	"cops",
	"county",
	"day",
	"game",
	"gangstas",
	"ginger",
	"halal",
	"johnson",
	"larceny",
	"life",
	"one",
	"parrot",
	"pisd",
	"play",
	"survive",
	"turtle",
	"world",
	"timer",
}

func rpname() string {
	mp := []byte("00RP")
	first := rpfirsts[rand.Intn(len(rpfirsts))]
	second := rpseconds[rand.Intn(len(rpseconds))]
	mp[0] = []byte(strings.ToUpper(first))[0]
	mp[1] = []byte(strings.ToUpper(second))[0]
	return fmt.Sprintf("%s: %s %s Roleplay", string(mp), strings.Title(first), strings.Title(second))
}

func (cm *CommandManager) commandRP(
	interaction *discordgo.InteractionCreate,
	args map[string]*discordgo.ApplicationCommandInteractionDataOption,
	settings types.CommandSettings,
) (
	context bool,
	err error,
) {

	cm.replyDirectly(interaction, rpname())
	return
}
