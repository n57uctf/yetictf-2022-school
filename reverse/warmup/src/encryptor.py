from pyfinite import ffield
F = ffield.FField(16)

SECRETFLAG = "YetiCTF{d3bug3r_1s_my_b3st_fr13end}\n\n\n"
PINCODE =0xd8fe
CRC     =0x0000

def transform(x):
    y = F.Subtract( F.Divide( F.Add(F.Multiply(x, PINCODE), 5), 0xF30), PINCODE)
    return y

li = []

for i in range( 0, int(len(SECRETFLAG)), 2 ):
    #print ( SECRETFLAG[i], SECRETFLAG[i+1], "        ",  hex(ord(SECRETFLAG[i])), hex(ord(SECRETFLAG[i+1]))  )
    x = ord(SECRETFLAG[i]) | ( ord(SECRETFLAG[i+1]) << 8)
    CRC = CRC ^ x
    y = transform (x)
    li.append ( y )

strout = ""
for l in li:
    strout += hex(l & 0xFF) + ", " + hex( (l >> 8) & 0xFF) + ", "

print ("#ifndef __PYGEN_H")
print ("#define __PYGEN_H\n")
print ("unsigned char SFLAG[] = {"+ strout[:-2] + "};")
print ("int CRC=", hex(CRC), ";")
print ("#define TRANSFORM(a,pin) gf_div(gf_sub(gf_shift_multiply(gf_add(a, pin), 0xF30), 5), pin)")
print ("#endif")

