package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Clan struct
type Clan struct {
	Status string `json:"status"`
	Meta   struct {
		Count int `json:"count"`
	} `json:"meta"`
	Data struct {
		Num500201233 struct {
			MembersCount    int         `json:"members_count"`
			Name            string      `json:"name"`
			CreatorName     string      `json:"creator_name"`
			ClanID          int         `json:"clan_id"`
			CreatedAt       int         `json:"created_at"`
			UpdatedAt       int         `json:"updated_at"`
			LeaderName      string      `json:"leader_name"`
			MembersIds      []int       `json:"members_ids"`
			CreatorID       int         `json:"creator_id"`
			Tag             string      `json:"tag"`
			OldName         interface{} `json:"old_name"`
			IsClanDisbanded bool        `json:"is_clan_disbanded"`
			RenamedAt       interface{} `json:"renamed_at"`
			OldTag          interface{} `json:"old_tag"`
			LeaderID        int         `json:"leader_id"`
			Description     string      `json:"description"`
		} `json:"500201233"`
	} `json:"data"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// Home Page
func Home(c *gin.Context) {
	clan := Clan{}

	getJson("https://api.worldofwarships.eu/wows/clans/info/?application_id=2fe567692a1f2cbfba128363247b542d&clan_id=500201233", &clan)

	c.HTML(http.StatusOK, "index.html", clan)
}
