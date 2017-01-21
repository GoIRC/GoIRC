package main

import "fmt"
import "strings"

type Packet struct {
  _type string
  sub_type string
  sender string
  nickname string
  username string
  hostname string
  channel string
  message string
  hash_id string
}

func parseUserInfo(info string) Packet {
  nickname := info[:strings.Index(info, "!")][1:]
  username := info[strings.Index(info, "!"):strings.Index(info, "@")][1:]
  hostname := info[strings.Index(info, "@"):]
  return Packet{nickname : nickname, username: username, hostname : hostname, sender: nickname}
}

func parseServerInfo(info string) Packet {
  servername := strings.Split(info, " ")[0][1:]
  return Packet{sender : servername}
}

func parseMessage(params []string) Packet {
  packet := parseUserInfo(params[0])
  packet.message = strings.Join(params[2:], " ")[1:]
  packet.channel = params[2][1:]
  packet._type = params[1]
  return packet
}

func parseServerMessage(params []string) Packet {
  packet := parseServerInfo(params[0])
  packet.message = strings.Join(params[3:], " ")[1:]
  packet.sub_type = params[2]
  return packet
}

func parseColonless(line string) Packet {
  return Packet{hash_id : strings.Split(line, " ")[1][1:]}
}

func autoParse(data string) Packet{
  line := strings.Trim(data, "\r\n")
  params := strings.Split(line, " ")
  switch line[0]{
    case ':':
      packet := map[bool]func([]string)(Packet){true: parseMessage, false: parseServerMessage}[strings.Contains(params[0], "!")](params)
      packet._type = params[1]
      return packet
    case 'P':
      packet := Packet{hash_id : params[1][1:]}
      packet._type = strings.ToLower(params[0])
      return packet
  }
  return Packet{}
}

func main(){
  fmt.Printf("%+v\n", autoParse(":GoBOT!GoBOT@memes-9ACDE63E.dhcp.drexel.edu JOIN :#iffi"))
  fmt.Printf("%+v\n", autoParse(":GoBOT!GoBOT@memes-9ACDE63E.dhcp.drexel.edu PRIVMSG #iffi :Eat my shit"))
  fmt.Printf("%+v\n", autoParse(":irc.dicksout.club NOTICE AUTH :*** Looking up your hostname..."))
  fmt.Printf("%+v\n", autoParse(":irc.dicksout.club 001 GoBOT :Welcome to the Dicks Out Club IRC Network GoBOT!GoBOT@n2-77-23.dhcp.drexel.edu"))
  fmt.Printf("%+v\n", autoParse(":irc.dicksout.club 251 GoBOT :There are 1 users and 1 invisible on 1 servers"))
  fmt.Printf("%+v\n", autoParse("PING :237894"))
}