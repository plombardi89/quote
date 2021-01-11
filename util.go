// Copyright 2019 Philip Lombardi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"github.com/plombardi89/gozeug/randomzeug"
	"os"
)

var adjectives = []string{
	"adept",
	"adorable",
	"alluvial",
	"ample",
	"authentic",
	"avaricious",
	"beleaguered",
	"bewitched",
	"bitter",
	"bleak",
	"blissful",
	"bogus",
	"bouncy",
	"buoyant",
	"bubbly",
	"buttery",
	"cavernous",
	"chubby",
	"cimmerian",
	"cromulent",
	"crooked",
	"delectable",
	"deplorable",
	"dilatory",
	"disingenuous",
	"dowdy",
	"droopy",
	"ellipsoidal",
	"embiggened",
	"enlightened",
	"euphoric",
	"fabulous",
	"fearless",
	"feckless",
	"flippant",
	"frivolous",
	"frosty",
	"fuzzy",
	"gargantuan",
	"gibbous",
	"ginormous",
	"grizzled",
	"grumpy",
	"gummy",
	"harmonious",
	"hasty",
	"haunting",
	"honest",
	"hortatory",
	"humble",
	"icky",
	"idle",
	"inglorious",
	"irenic",
	"itchy",
	"janky",
	"jocular",
	"jolly",
	"jovial",
	"klutzy",
	"kooky",
	"limp",
	"livid",
	"loquacious",
	"luminous",
	"lumbering",
	"majestic",
	"meaty",
	"mellow",
	"menacing",
	"mirthful",
	"munificient",
	"mushy",
	"naughty",
	"negative",
	"nerdy",
	"nippy",
	"oddball",
	"oily",
	"perky",
	"pesky",
	"piquant",
	"poised",
	"pokable",
	"posh",
	"prickly",
	"queruluos",
	"quintessential",
	"raging",
	"ravenous",
	"rhapsodic",
	"serene",
	"slippery",
	"snippy",
	"tart",
	"tasty",
	"tender",
	"thunderous",
	"trim",
	"unctuous",
	"undulating",
	"unkempt",
	"unripe",
	"velvety",
	"vengeful",
	"voluminous",
	"warlike",
	"wiry",
	"wry",
	"yummy",
	"zany",
	"zesty",
}

var fruits = []string{
	"acai",
	"apple",
	"apricot",
	"banana",
	"blackberry",
	"blueberry",
	"cherry",
	"coconut",
	"cranberry",
	"date",
	"elderberry",
	"grape",
	"grapefruit",
	"jackfruit",
	"kiwi",
	"kumquat",
	"lemon",
	"lime",
	"mango",
	"mulberry",
	"nectarine",
	"orange",
	"papaya",
	"passionfruit",
	"pear",
	"persimmon",
	"plum",
	"pineapple",
	"pomegranate",
	"raspberry",
	"snozzberry",
	"strawberry",
	"tangerine",
}

func getServerID(hostnameFunc func() (string, error), random *randomzeug.Random) string {
	if hostnameFunc == nil {
		hostnameFunc = os.Hostname
	}

	if hostname, err := hostnameFunc(); err == nil && hostname != "" {
		return hostname
	}

	adjective := random.RandomSelectionFromStringSlice(adjectives)
	fruit := random.RandomSelectionFromStringSlice(fruits)

	return fmt.Sprintf("%s-%s-%s", adjective, fruit, random.RandomString(8))
}

func getEnv(name, fallback string) string {
	res := os.Getenv(name)
	if res == "" {
		res = fallback
	}

	return res
}
