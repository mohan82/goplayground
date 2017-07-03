package json

import (
	"testing"
	"strings"
	"math/rand"
	"strconv"
	"time"
)

func SerializeTestString() string {
	return New().SetString("test", "testValue").Serialize()
}

func TestSerializer_SerializeValue(t *testing.T) {
	expected := "testValue"
	t.Logf("%v ", t.Name())
	output := SerializeTestString()
	if !strings.Contains(output, expected) {
		t.Fatalf("Expected output to contain %v but got %v", expected, output)
	}

}

func TestSerializer_SerializeFieldLength(t *testing.T) {
	t.Logf("%v ", t.Name())
	expected := 2
	output := SerializeTestString()
	if len(strings.Split(output, ":")) != expected {
		t.Errorf("Expected output to contain  %v but got %v", expected, output)

	}
}

var BOOLEAN_TEST = []struct {
	expected bool
	actual   string
}{
	{true, "true"},
	{false, "false"},
}

func TestSerializer_SetBoolean(t *testing.T) {
	for _, data := range BOOLEAN_TEST {
		output := New().SetBoolean("test", data.expected).Serialize()
		if !strings.Contains(output, data.actual) {
			t.Errorf("Expect output to contain %v found %v", data.actual, output)
		}
	}

}

func TestSerializer_SetIntArray(t *testing.T) {
	a := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		testValue := rand.Int()
		a[i] = testValue
		t.Logf("Random value is %v\n",testValue)
	}
	output := New().SetIntArray("test", a).Serialize()
	t.Logf("Output is %v", output)
	for _, testVal := range a {
		if !strings.Contains(output, strconv.Itoa(testVal)) {
			t.Errorf("Expect output to contain %v found %v", testVal, output)
		}
	}
}


func TestSerializer_SetStringArray(t *testing.T) {
	a := make([]string, 10, 10)
	for i := 0; i < 10; i++ {
		testValue :=strconv.Itoa(rand.Int())+ ":test"
		a[i] = testValue
		t.Logf("Random value is %v\n",testValue)
	}
	output := New().SetStringArray("test", a).Serialize()
	t.Logf("Output is %v", output)
	for _, testVal := range a {
		if !strings.Contains(output, testVal) {
			t.Errorf("Expect output to contain %v found %v", testVal, output)
		}
	}
}

func TestSerializer_SetTimeArray(t *testing.T) {
	a := make([]time.Time, 10, 10)
	for i := 0; i < 10; i++ {
		testValue :=time.Now().Add(-time.Duration(rand.Intn(10000))* time.Second )
		a[i] = testValue
		t.Logf("Random value is %v\n",testValue.Unix())
	}
	output := New().SetTimeArray("test", a).Serialize()
	t.Logf("Output is %v", output)
	for _, testVal := range a {
		if !strings.Contains(output, strconv.FormatInt(testVal.Unix(),10)) {
			t.Errorf("Expect output to contain %v found %v", testVal.Unix(), output)
		}
	}
}
