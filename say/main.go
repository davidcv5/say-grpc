package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/davidcv5/say-grpc/api"
	"github.com/sirupsen/logrus"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where output will be written")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: flag.Arg(0)}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		logrus.Fatalf("could not say %s: %v", text.Text, err)
	}
	if err = ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		logrus.Fatalf("could not write to file %s: %v", *output, err)
	}
}
