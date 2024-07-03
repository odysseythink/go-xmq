package xmq_test

import (
	"flag"

	"mlib.com/go-xmq"
)

func ExampleConfigFlag() {
	cfg := xmq.NewConfig()
	flagSet := flag.NewFlagSet("", flag.ExitOnError)

	flagSet.Var(&xmq.ConfigFlag{cfg}, "consumer-opt", "option to pass through to xmq.Consumer (may be given multiple times)")
	flagSet.PrintDefaults()

	err := flagSet.Parse([]string{
		"--consumer-opt=heartbeat_interval,1s",
		"--consumer-opt=max_attempts,10",
	})
	if err != nil {
		panic(err.Error())
	}
	println("HeartbeatInterval", cfg.HeartbeatInterval)
	println("MaxAttempts", cfg.MaxAttempts)
}
