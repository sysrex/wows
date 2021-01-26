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

// Player struct
type Player struct {
	Status string `json:"status"`
	Meta   struct {
		Count  int         `json:"count"`
		Hidden interface{} `json:"hidden"`
	} `json:"meta"`
	Data struct {
		Num500827075 struct {
			LastBattleTime int         `json:"last_battle_time"`
			AccountID      int         `json:"account_id"`
			LevelingTier   int         `json:"leveling_tier"`
			CreatedAt      int         `json:"created_at"`
			LevelingPoints int         `json:"leveling_points"`
			UpdatedAt      int         `json:"updated_at"`
			Private        interface{} `json:"private"`
			HiddenProfile  bool        `json:"hidden_profile"`
			LogoutAt       int         `json:"logout_at"`
			Karma          interface{} `json:"karma"`
			Statistics     struct {
				Distance int `json:"distance"`
				Battles  int `json:"battles"`
				Pvp      struct {
					MaxXp             int `json:"max_xp"`
					DamageToBuildings int `json:"damage_to_buildings"`
					MainBattery       struct {
						MaxFragsBattle int   `json:"max_frags_battle"`
						Frags          int   `json:"frags"`
						Hits           int   `json:"hits"`
						MaxFragsShipID int64 `json:"max_frags_ship_id"`
						Shots          int   `json:"shots"`
					} `json:"main_battery"`
					MaxShipsSpottedShipID int64 `json:"max_ships_spotted_ship_id"`
					MaxDamageScouting     int   `json:"max_damage_scouting"`
					ArtAgro               int   `json:"art_agro"`
					MaxXpShipID           int64 `json:"max_xp_ship_id"`
					ShipsSpotted          int   `json:"ships_spotted"`
					SecondBattery         struct {
						MaxFragsBattle int   `json:"max_frags_battle"`
						Frags          int   `json:"frags"`
						Hits           int   `json:"hits"`
						MaxFragsShipID int64 `json:"max_frags_ship_id"`
						Shots          int   `json:"shots"`
					} `json:"second_battery"`
					MaxFragsShipID            int64       `json:"max_frags_ship_id"`
					Xp                        int         `json:"xp"`
					SurvivedBattles           int         `json:"survived_battles"`
					DroppedCapturePoints      int         `json:"dropped_capture_points"`
					MaxDamageDealtToBuildings int         `json:"max_damage_dealt_to_buildings"`
					TorpedoAgro               int         `json:"torpedo_agro"`
					Draws                     int         `json:"draws"`
					ControlCapturedPoints     int         `json:"control_captured_points"`
					BattlesSince510           int         `json:"battles_since_510"`
					MaxTotalAgroShipID        int64       `json:"max_total_agro_ship_id"`
					PlanesKilled              int         `json:"planes_killed"`
					Battles                   int         `json:"battles"`
					MaxShipsSpotted           int         `json:"max_ships_spotted"`
					MaxSuppressionsShipID     interface{} `json:"max_suppressions_ship_id"`
					SurvivedWins              int         `json:"survived_wins"`
					Frags                     int         `json:"frags"`
					DamageScouting            int         `json:"damage_scouting"`
					MaxTotalAgro              int         `json:"max_total_agro"`
					MaxFragsBattle            int         `json:"max_frags_battle"`
					CapturePoints             int         `json:"capture_points"`
					Ramming                   struct {
						MaxFragsBattle int   `json:"max_frags_battle"`
						Frags          int   `json:"frags"`
						MaxFragsShipID int64 `json:"max_frags_ship_id"`
					} `json:"ramming"`
					SuppressionsCount    int `json:"suppressions_count"`
					MaxSuppressionsCount int `json:"max_suppressions_count"`
					Torpedoes            struct {
						MaxFragsBattle int   `json:"max_frags_battle"`
						Frags          int   `json:"frags"`
						Hits           int   `json:"hits"`
						MaxFragsShipID int64 `json:"max_frags_ship_id"`
						Shots          int   `json:"shots"`
					} `json:"torpedoes"`
					MaxPlanesKilledShipID int64 `json:"max_planes_killed_ship_id"`
					Aircraft              struct {
						MaxFragsBattle int   `json:"max_frags_battle"`
						Frags          int   `json:"frags"`
						MaxFragsShipID int64 `json:"max_frags_ship_id"`
					} `json:"aircraft"`
					TeamCapturePoints               int         `json:"team_capture_points"`
					ControlDroppedPoints            int         `json:"control_dropped_points"`
					MaxDamageDealt                  int         `json:"max_damage_dealt"`
					MaxDamageDealtToBuildingsShipID interface{} `json:"max_damage_dealt_to_buildings_ship_id"`
					MaxDamageDealtShipID            int64       `json:"max_damage_dealt_ship_id"`
					Wins                            int         `json:"wins"`
					Losses                          int         `json:"losses"`
					DamageDealt                     int         `json:"damage_dealt"`
					MaxPlanesKilled                 int         `json:"max_planes_killed"`
					MaxScoutingDamageShipID         int64       `json:"max_scouting_damage_ship_id"`
					TeamDroppedCapturePoints        int         `json:"team_dropped_capture_points"`
					BattlesSince512                 int         `json:"battles_since_512"`
				} `json:"pvp"`
			} `json:"statistics"`
			Nickname       string `json:"nickname"`
			StatsUpdatedAt int    `json:"stats_updated_at"`
		} `json:"500827075"`
	} `json:"data"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
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
	player := Player{}

	getJSON("https://api.worldofwarships.eu/wows/clans/info/?application_id=2fe567692a1f2cbfba128363247b542d&clan_id=500201233", &clan)
	getJSON("https://api.worldofwarships.eu/wows/account/info/?application_id=2fe567692a1f2cbfba128363247b542d&account_id=500827075", &player)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"clan":   clan,
		"player": player,
	})
}
