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

func parseUserInfo(info string) Packet{
  nickname := strings.Split(info, "!")[0][1:]
  username := strings.Split(strings.Split(info, "@")[0], "!")[1]
  hostname := strings.Split(info, "@")[1]
  return Packet{nickname : nickname, username: username, hostname : hostname, sender: nickname}
}

func parseServerInfo(info string) Packet{
  servername := strings.Split(info, " ")[0][1:]
  return Packet{sender : servername}
}


func parse(data string){
  line := strings.Trim(data, "\r\n")
  switch line[0]{
    case ':':
      packet := Packet{}
      params := strings.Split(line, " ")
      if(strings.Contains(params[0], "!")){
        packet = parseUserInfo(params[0])
      } else{
        packet = parseServerInfo(params[0])
        packet.message = strings.Join(params[3:], " ")[1:]
        packet.sub_type = params[2]
      }
      packet._type = params[1]
      fmt.Printf("%+v\n", packet)
    case 'P':
      packet := Packet{hash_id : strings.Split(line, " ")[1][1:]}
      fmt.Printf("%+v\n", packet)
  }
}

func main(){
  parse(":GoBOT!GoBOT@memes-9ACDE63E.dhcp.drexel.edu JOIN :#iffi")
  parse(":GoBOT!GoBOT@memes-9ACDE63E.dhcp.drexel.edu PRIVMSG #iffi :Eat my shit")
  parse(":irc.dicksout.club NOTICE AUTH :*** Looking up your hostname...")
  parse(":irc.dicksout.club 001 GoBOT :Welcome to the Dicks Out Club IRC Network GoBOT!GoBOT@n2-77-23.dhcp.drexel.edu")
  parse(":irc.dicksout.club 251 GoBOT :There are 1 users and 1 invisible on 1 servers")

  parse("PING :237894")
}