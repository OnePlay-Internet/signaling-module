package grpc

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OnePlay-Internet/webrtc-proxy/util/config"
	"github.com/OnePlay-Internet/webrtc-proxy/util/tool"
	"github.com/pion/webrtc/v3"
)

func TestClient(t *testing.T) {
	conf := config.GrpcConfig{
		ServerAddress: "localhost",
		Port:          8000,
	}
	client, err := InitGRPCClient(&conf,tool.GetDevice())
	if err != nil {
		t.Error(err)
	}
	shutdown_channel := make(chan bool)
	client.OnSDP(func(i *webrtc.SessionDescription) {
		json, _ := json.Marshal(i)
		fmt.Printf("%s\n", json)
		// shutdown_channel<-true;
	})
	client.OnICE(func(i *webrtc.ICECandidateInit) {
		json, _ := json.Marshal(i)
		fmt.Printf("%s\n", json)
		// shutdown_channel<-true;
	})
	client.OnDeviceSelect(func (mon tool.Monitor,soundcard tool.Soundcard,num int) error {
		return nil;
	})
	test := "test"
	var test_num uint16
	test_num = 0

	err = client.SendICE(&webrtc.ICECandidateInit{
		Candidate:     test,
		SDPMid:        &test,
		SDPMLineIndex: &test_num,
	})
	if err != nil {
		t.Error(err)
	}
	err = client.SendSDP(&webrtc.SessionDescription{
		Type: 1,
		SDP: "test",
	})
	if err != nil {
		t.Error(err)
	}

	<-shutdown_channel
}
