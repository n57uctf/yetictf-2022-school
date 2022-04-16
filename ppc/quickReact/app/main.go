package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"time"
)

func ServeConnection(c net.Conn, winnerFlag string) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	regexp, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal("error when compiling a regexp")
	}
	c.Write([]byte("React as quick as possible. You have only 10 seconds to answer correctly. When you guess all correctly, you receive a flag. Good Luck\n"))

	for i := 0; i < 20; i++ {
		taskChoose := r1.Intn(8)

		switch {
		case taskChoose == 0:
			{
				num1 := r1.Intn(1000)
				num2 := r1.Intn(1000)
				answer := strconv.Itoa(num1 + num2)
				c.Write([]byte(fmt.Sprintf("%v+%v=", num1, num2)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 1:
			{
				num1 := r1.Intn(1000) + 1000
				num2 := r1.Intn(1000)
				answer := strconv.Itoa(num1 - num2)
				c.Write([]byte(fmt.Sprintf("%v-%v=", num1, num2)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 2:
			{
				num1 := r1.Intn(1000)
				num2 := r1.Intn(1000)
				answer := strconv.Itoa(num1 * num2)
				c.Write([]byte(fmt.Sprintf("%v*%v=", num1, num2)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 3:
			{
				num1 := r1.Intn(1000)*10 + 1000
				num2 := r1.Intn(1000) + 1
				num1 -= (num1 % num2)
				answer := strconv.Itoa(num1 / num2)
				c.Write([]byte(fmt.Sprintf("%v/%v=", num1, num2)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 4:
			{
				num1 := r1.Intn(1000)
				hasher := md5.New()
				hasher.Write([]byte(strconv.Itoa(num1)))
				answer := hex.EncodeToString(hasher.Sum(nil))
				c.Write([]byte(fmt.Sprintf("md5(%v)=", num1)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 5:
			{
				num1 := r1.Intn(1000)
				hasher := sha256.New()
				hasher.Write([]byte(strconv.Itoa(num1)))
				answer := hex.EncodeToString(hasher.Sum(nil))
				c.Write([]byte(fmt.Sprintf("sha256(%v)=", num1)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 6:
			{
				num1 := r1.Intn(1000)
				hasher := sha1.New()
				hasher.Write([]byte(strconv.Itoa(num1)))
				answer := hex.EncodeToString(hasher.Sum(nil))
				c.Write([]byte(fmt.Sprintf("sha1(%v)=", num1)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		case taskChoose == 7:
			{
				num1 := r1.Intn(1000)
				hasher := sha512.New()
				hasher.Write([]byte(strconv.Itoa(num1)))
				answer := hex.EncodeToString(hasher.Sum(nil))
				c.Write([]byte(fmt.Sprintf("sha512(%v)=", num1)))
				buf := make([]byte, 1024)
				c.Read(buf)
				data := regexp.ReplaceAllString(string(buf), "")
				if data != answer {
					log.Printf(fmt.Sprintf("%v == %v\n", data, answer))
					c.Close()
					return
				}
			}
		}
	}

	c.Write([]byte(winnerFlag))
}

func main() {
	addr := flag.String("addr", "0.0.0.0:12345", "Choose a addr to start a remote game")
	winnerFlag := flag.String("flag", "YetiCTF{HEy_nicelyDONE_congrats}", "CTF Flag")
	flag.Parse()
	log.Println("Launching server")
	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("Error while reserving port: %v\n", err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error while starting listening connection: %v\n", err)
		}
		log.Printf(fmt.Sprintf("Serving connection from %v", conn.RemoteAddr().String()))
		go func() {
			for {
				select {
				case <-time.After(time.Second * 10):
					conn.Close()
					return
				}
			}
		}()

		go ServeConnection(conn, *winnerFlag)
	}

}
