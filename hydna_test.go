package hydna

import "bytes"
import "testing"

func TestSend(t *testing.T) {
    data := bytes.NewBufferString("Test data")

    err := Send("testing.hydna.net", data)
    if err != nil {
        t.Errorf("Unable to send data")
    }

    denyMessage := Send("testing.hydna.net/open-deny", data)
    if denyMessage == nil || denyMessage.Error() != "DENIED" {
        t.Errorf("Expected error DENIED, got %s", denyMessage.Error())
    }
}

func TestEmit(t *testing.T) {
    data := bytes.NewBufferString("Test data")

    err := Emit("testing.hydna.net", data)
    if err != nil {
        t.Errorf("Unable to send data")
    }

    denyMessage := Emit("testing.hydna.net/open-deny", data)
    if denyMessage == nil || denyMessage.Error() != "DENIED" {
        t.Errorf("Expected error DENIED, got %s", denyMessage.Error())
    }
}
