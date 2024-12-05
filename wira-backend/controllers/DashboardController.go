package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"wira-backend/connection"
)

type DashboardController struct {
	
}

type RankingTableStruc struct {
	Username 	string `json:"username"`
    Email 		string `json:"email"`
    Class       int    `json:"class_id"`
    Score       float64 `json:"reward_score"`
    Name        string `json:"character_name"`

}

func NewDashboardController() *DashboardController {
	return &DashboardController{
		//Inject services
	}
}

func (ctl *DashboardController) RankingTable(w http.ResponseWriter, r *http.Request)  {

	query := `
		Select charc.name as character_name, acc.username, acc.email, charc.class_id, scr.reward_score 
        From characters charc
        Join accounts acc on charc.acc_id = acc.acc_id
        Join scores scr on charc.char_id = scr.char_id
    `

    results, err := connection.GetJoinedData(query)
    if err != nil {
        log.Printf("Error querying characters with accounts: %v", err)
        return 
    }

    var collections []RankingTableStruc
    for _, result := range results {
        item := RankingTableStruc{
            Username	: result["username"].(string),
            Email		: result["email"].(string),
            Class		: int(result["class_id"].(int64)),
            Score		: float64(result["reward_score"].(float64)),
            Name		: result["character_name"].(string),
        }
        collections = append(collections, item)
    }

	log.Printf("Retrieved %d characters with accounts from the database", len(collections))

	json.NewEncoder(w).Encode(collections)
}