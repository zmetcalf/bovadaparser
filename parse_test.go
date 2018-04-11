package bovadaparser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var sportsData Schedule

func init() {
	data, err := ioutil.ReadFile("example.xml")
	if err != nil {
		fmt.Printf("Unable to open file %s", err)
	}

	sportsData, err = Unmarshal(data)
	if err != nil {
		fmt.Printf("Unable to read file %s", err)
	}
}

func TestSportsDataParse(t *testing.T) {
	if sportsData.PublishDate != "04/09/2017" {
		t.Errorf("Unable to get Publish Date %s", sportsData.PublishDate)
	}

	if sportsData.PublishTime != "08:50" {
		t.Errorf("Unable to get Publish Time %s", sportsData.PublishTime)
	}
}

func TestEventTypeParse(t *testing.T) {
	if sportsData.EventType.ID != "MLB" {
		t.Errorf("Unable to get Event Type ID %s",
			sportsData.EventType.ID)
	}
}

func TestDateParse(t *testing.T) {
	if sportsData.EventType.Date.GameDate != "Mon, Apr 10, 17" {
		t.Errorf("Unable to get Date GameDate %s",
			sportsData.EventType.Date.GameDate)
	}
}

func TestEventParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].ID != "2734398" {
		t.Errorf("Unable to get Event ID %s",
			sportsData.EventType.Date.Event[0].ID)
	}
}

func TestTimeParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Time.Start != "01:05PM EDT" {
		t.Errorf("Unable to get Event Time Start %s",
			sportsData.EventType.Date.Event[0].Time.Start)
	}
}

func TestCompetitorIdParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].ID != "2734398-10032" {
		t.Errorf("Unable to get Competitor ID %s",
			sportsData.EventType.Date.Event[0].Competitor[0].ID)
	}
}

func TestCompetitorNameParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Name != "New York Yankees" {
		t.Errorf("Unable to get Competitor Name %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Name)
	}
}

func TestLineTypeParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Type != "Moneyline" {
		t.Errorf("Unable to get Line Type %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Type)
	}
}

func TestChoiceValueParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Value != "-140" {
		t.Errorf("Unable to get Choice Value %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Value)
	}
}

func TestOddsFractionParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Fraction != "8/11" {
		t.Errorf("Unable to get Odds Fraction %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Fraction)
	}
}

func TestOddsLineParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Line != "-140" {
		t.Errorf("Unable to get Odds Line %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Line)
	}
}

func TestOddsMultiplierParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Multiplier != "1.714" {
		t.Errorf("Unable to get Odds Multiplier %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Multiplier)
	}
}

func TestOddsRiskParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Risk != "140" {
		t.Errorf("Unable to get Odds Risk %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Risk)
	}
}

func TestOddsWinParse(t *testing.T) {
	if sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Win != "100" {
		t.Errorf("Unable to get Odds Risk %s",
			sportsData.EventType.Date.Event[0].Competitor[0].Line[0].Choice[0].Odds[0].Win)
	}
}
