#!/usr/bin/env python3
import socket, re, time

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("localhost", 12345))

data = sock.recv(4096) 
#print(data)

STATSRegexp = r"Your stats:\nLevel: ([\d]+)\nHP: ([\d]+)\/([\d]+)\nEXP: ([\d]+)\/([\d]+)\nMoney: ([\d]+)\nHealing Potions: ([\d]+)\nDamage: ([\d]+)"
class Hero:
    def __init__(self,lvl, chp, mhp, cexp, mexp, coins, potions, damage):
        self.lvl = lvl
        self.chp = chp
        self.mhp = mhp
        self.cexp = cexp
        self.mexp = mexp
        self.coins = coins
        self.potions = potions 
        self.damage = damage
    def __str__(self):
        return "%s, %s, %s, %s, %s, %s, %s, %s" % (self.lvl, self.chp, self.mhp, self.cexp, self.mexp, self.coins, self.potions, self.damage)

hero = Hero(0,0,0,0,0,0,0,0)

killed = 0

while True:
    #we need to check stats to continue making decisions
    sock.send(str.encode("s\n"))
    data = sock.recv(4096)
    match = re.search(STATSRegexp, data.decode())
    if match[0]:
        hero.lvl = int(match[1])
        hero.chp = int(match[2])
        hero.mhp = int(match[3])
        hero.cexp = int(match[4])
        hero.mexp = int(match[5])
        hero.coins = int(match[6])
        hero.potions = int(match[7])
        hero.damage = int(match[8])
    else:
        break
    if hero.coins >= hero.lvl:
        sock.send(str.encode("b\n"))
        sock.recv(4096)
        while hero.coins >= hero.lvl:
            sock.send(str.encode("s\n"))
            sock.recv(4096)
            hero.coins -= hero.lvl
        sock.send(str.encode("q\n"))
        sock.recv(4096)

    sock.send(str.encode("f\n"))
    sock.recv(4096)
    sock.send(str.encode("a\n"))
    data = sock.recv(4096).decode()
    if "Game Over" in data:
        break
    if "YetiCTF{" in data:
        print(data)
        break
    data=data.split("\n")[1]
    if "hits" in data:
        print(data)
        while True:
            sock.send(str.encode("a\n"))
            data = sock.recv(4096).decode().split("\n")[1]
            print("DATA: %s,\tHERO STATS: %s,\tKILLED: %s,\t" % (data, hero, killed))
            if "died" in data:
                break
    if "died" in data:
        killed += 1
        print("DATA: %s,\tHERO STATS: %s,\tKILLED: %s,\t" % (data, hero, killed))
