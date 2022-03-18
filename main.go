package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
)

var relationShipPartnerCsvHeader = [...]string{
	"Id",
	"Name",
	"Furigana",
	"Category",
	"Region",
	"Rank",
	"StartDate",
	"Work",
	"PostalCode",
	"Address",
	"PhoneNumber",
	"Url",
	"Benefit",
	"Comment",
	"IsActive",
	"CreatedAt",
	"UpdatedAt",
	"UpdatedBy",
	"ContinuationYears",
}

type RelationShipPartner struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Furigana          string `json:"furigana"`
	Category          string `json:"category"`
	Region            string `json:"region"`
	Rank              string `json:"rank"`
	StartDate         string `json:"startDate"`
	Work              string `json:"work"`
	PostalCode        string `json:"postalCode"`
	Address           string `json:"address"`
	PhoneNumber       string `json:"phoneNumber"`
	Url               string `json:"url"`
	Benefit           string `json:"benefit"`
	Comment           string `json:"comment"`
	IsActive          int    `json:"isActive"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
	UpdatedBy         int    `json:"updatedBy"`
	ContinuationYears int    `json:"continuationYears"`
}

func (r RelationShipPartner) Record() []string {
	return []string{
		fmt.Sprint(r.Id),
		r.Name,
		r.Furigana,
		r.Category,
		r.Region,
		r.Rank,
		r.StartDate,
		r.Work,
		r.PostalCode,
		r.Address,
		r.PhoneNumber,
		r.Url,
		r.Benefit,
		r.Comment,
		fmt.Sprint(r.IsActive),
		r.CreatedAt,
		r.UpdatedAt,
		fmt.Sprint(r.UpdatedBy),
		fmt.Sprint(r.ContinuationYears),
	}
}

var relationShipPartnerUrls = map[string]string{
	"food":         "https://site-api.consadole-sapporo.jp/api/partner/support/1",
	"medical":      "https://site-api.consadole-sapporo.jp/api/partner/support/2",
	"sales":        "https://site-api.consadole-sapporo.jp/api/partner/support/3",
	"construction": "https://site-api.consadole-sapporo.jp/api/partner/support/4",
	"newspaper":    "https://site-api.consadole-sapporo.jp/api/partner/support/5",
	"law":          "https://site-api.consadole-sapporo.jp/api/partner/support/6",
	"sightseeing":  "https://site-api.consadole-sapporo.jp/api/partner/support/7",
	"estate":       "https://site-api.consadole-sapporo.jp/api/partner/support/8",
	"finance":      "https://site-api.consadole-sapporo.jp/api/partner/support/9",
	"others":       "https://site-api.consadole-sapporo.jp/api/partner/support/10",
}

func main() {
	var relationShipPartners []RelationShipPartner

	for _, v := range relationShipPartnerUrls {
		resp, err := http.Get(v)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		byte, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var r []RelationShipPartner
		if err := json.Unmarshal(byte, &r); err != nil {
			log.Fatal(err)
		}
		relationShipPartners = append(relationShipPartners, r...)
	}

	sort.Slice(relationShipPartners, func(i, j int) bool { return relationShipPartners[i].Id < relationShipPartners[j].Id })

	dirName := "public"
	err := os.Mkdir(dirName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	file, err := os.OpenFile(path.Join(dirName, "relationship-partner.json"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// GitHub でレコードの差分がよく見えるように JSON 配列の中のレコードを 1 行ずつ改行させたい
	file.WriteString("[\n")
	defer file.WriteString("]")

	for _, rerelationShipPartner := range relationShipPartners {
		byte, err := json.Marshal(rerelationShipPartner)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString("  ")
		file.Write(byte)
		file.WriteString(",\n")
	}
}
