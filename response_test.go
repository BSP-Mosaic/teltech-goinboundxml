package inboundxml

import (
	"testing"
)

func TestNewResponse(t *testing.T) {
	var (
		response *Response
	)

	response = NewResponse()

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Say("Hello World", nil)

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response><Say>Hello World</Say></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Hangup()

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response><Say>Hello World</Say><Hangup /></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Reset()

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Play("http://www.example.com/audio.mp3", nil)

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response><Play>http://www.example.com/audio.mp3</Play></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Reset()
	response.Play("http://www.example.com/audio.mp3", &PlayAttrs{Loop: 1})

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response><Play loop=\"1\">http://www.example.com/audio.mp3</Play></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Reset()
	response.Record(&RecordAttrs{MaxLength: 60})

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response><Record maxLength=\"60\" /></Response>" {
		t.Errorf("Unexpected string generated by response")
	}

	response.Reset()
	response.Record(&RecordAttrs{Action: "http://www.example.com/recording-callback", MaxLength: 60})

	if response.String() != "<?xml version=\"1.0\" encoding=\"utf-8\"?><Response><Record action=\"http://www.example.com/recording-callback\" maxLength=\"60\" /></Response>" {
		t.Errorf("Unexpected string generated by response")
	}
}