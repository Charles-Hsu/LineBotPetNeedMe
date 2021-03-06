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
	"log"

	"golang.org/x/text/encoding/unicode/utf32"
)

//PetType :
type PetType int

var BrownEmoji string

const (
	//Dog :
	Dog PetType = iota
	//Cat :
	Cat
)

func init() {
	var err error
	utf32BEIB := utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM)
	dec := utf32BEIB.NewDecoder()
	BrownEmoji, err = dec.String("\x00\x10\x00\x84")
	if err != nil {
		log.Print(err)
	}

}

//Pet :
type Pet struct {
	ID              string `json:"_id"`
	Name            string `json:"Name"`
	Sex             string `json:"Sex"`
	Type            string `json:"Type"`
	Build           string `json:"Build"`
	Age             string `json:"Age"`
	Variety         string `json:"Variety"`
	Reason          string `json:"Reason"`
	AcceptNum       string `json:"AcceptNum"`
	ChipNum         string `json:"ChipNum"`
	IsSterilization string `json:"IsSterilization"`
	HairType        string `json:"HairType"`
	Note            string `json:"Note"`
	Resettlement    string `json:"Resettlement"`
	Phone           string `json:"Phone"`
	Email           string `json:"Email"`
	ChildreAnlong   string `json:"ChildreAnlong"`
	AnimalAnlong    string `json:"AnimalAnlong"`
	Bodyweight      string `json:"Bodyweight"`
	ImageName       string `json:"ImageName"`
}

//PetType :
func (p *Pet) PetType() PetType {
	var retType PetType
	switch p.Variety {
	case "犬":
		retType = Dog
	case "貓":
		retType = Cat
	}

	return retType
}

//DisplayPet : Display single pet on chatbot
func (p *Pet) DisplayPet() string {
	if len(p.ImageName) > 0 {
		p.ImageName = getSecureImageAddress(p.ImageName)
		log.Println("img:", p.ImageName)
	}

	return fmt.Sprintf("%s 您好 \n 動物品種: %s \n 毛色: %s \n 體型: %s \n 性別: %s \n 名為: %s \n 公告收容所: %s \n 領養電話為: %s  \n 圖片位置: %s", BrownEmoji, p.Variety, p.HairType, p.Type, p.Sex, p.Name, p.Resettlement, p.Phone, p.ImageName)
}
