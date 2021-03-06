package zeroless

import "testing"

func checkExchangedData(t *testing.T, received []string, expected []string) {
	for i, _ := range received {
		if received[i] != expected[i] {
			t.Error("Received:", received, "Expected:", expected)
		}
	}
}

func TestDoublePingThenPong(t *testing.T) {
	pairReceiver := NewServer(1054).Pair()
	pairSender := NewClient().ConnectLocal(1054).Pair()

	pings := [][]string{
		[]string{"ping1"},
		[]string{"ping2"},
		[]string{"ping11", "ping12"},
		[]string{"ping21", "ping22"},
	}
	pongs := [][]string{
		[]string{"pong1"},
		[]string{"pong2"},
		[]string{"pong11", "ping12"},
		[]string{"pong21", "pong22"},
	}

	for i, ping := range pings {
		pairSender <- ping
		pairSender <- ping
		y := 0
		for result := range pairReceiver {
			checkExchangedData(t, result, ping)

			y += 1
			if y == 2 {
				pairReceiver <- pongs[i]
				break
			}
		}

		result := <-pairSender
		checkExchangedData(t, result, pongs[i])
	}
}
