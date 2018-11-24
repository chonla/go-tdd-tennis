package tennis

import "testing"

func TestResetMatchShouldReturn0to0(t *testing.T) {
	expected := "0-0"
	scorer := new(Scorer)

	result := scorer.Reset()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointAfterResetSomeScoreShouldReturn15to0(t *testing.T) {
	expected := "15-0"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.Reset()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointAfterResetShouldReturn15to0(t *testing.T) {
	expected := "15-0"
	scorer := new(Scorer)

	scorer.Reset()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointTwiceShouldReturn30to0(t *testing.T) {
	expected := "30-0"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointTripleShouldReturn40to0(t *testing.T) {
	expected := "40-0"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetPointAfterResetShouldReturn0to15(t *testing.T) {
	expected := "0-15"
	scorer := new(Scorer)

	scorer.Reset()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetPointTwiceShouldReturn0to30(t *testing.T) {
	expected := "0-30"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.BGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetPointTripleShouldReturn0to40(t *testing.T) {
	expected := "0-40"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.BGetPoint()
	scorer.BGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetPointAfterAGetSomePointShouldReturnCurrentAPointto15(t *testing.T) {
	expected := "30-15"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetAdvantage(t *testing.T) {
	expected := "A-40"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetAdvantage(t *testing.T) {
	expected := "40-A"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAWinsWithNoDeuce(t *testing.T) {
	expected := "A Wins"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBWinsWithNoDeuce(t *testing.T) {
	expected := "B Wins"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAWinsWithDeuce(t *testing.T) {
	expected := "A Wins"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.AGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBWinsWithDeuce(t *testing.T) {
	expected := "B Wins"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointAfterBHasAdvantageShouldReturnTo40To40(t *testing.T) {
	expected := "40-40"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.AGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetPointAfterAHasAdvantageShouldReturnTo40To40(t *testing.T) {
	expected := "40-40"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	scorer.BGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointAfterGameEndShouldReturnError(t *testing.T) {
	expected := "Error"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestBGetPointAfterGameEndShouldReturnError(t *testing.T) {
	expected := "Error"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	result := scorer.BGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}

func TestAGetPointAfterGameEndAndResettShouldReturn15to0(t *testing.T) {
	expected := "15-0"
	scorer := new(Scorer)

	scorer.Reset()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.AGetPoint()
	scorer.Reset()
	result := scorer.AGetPoint()

	if result != expected {
		t.Errorf("expect result to be %v but %v", expected, result)
	}
}
