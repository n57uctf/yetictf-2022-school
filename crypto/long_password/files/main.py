import hashlib

SHIFRPREDL = '?'
k = []
s = len(SHIFRPREDL)
stroka = ''
for i in range(len(SHIFRPREDL)):
    t = ord(SHIFRPREDL[i]) ^ s
    s ^= t
    k.append(format(t, '#b'))
for i in range(len(k)):
    stroka += hashlib.md5(k[i].encode('utf-8')).hexdigest()
print(stroka)
#>>>1feab75348a254351398df7911420f6aece0895c75d170ebff6efca4a651961b232751f101ed242222a3ddf51f5b6324cc7f717a19f2b20b0b0c3c2741666e2659063240059123a9d62c9b6521617d3a9ac089ecfe7861ea06523d77d68f9460491d1997da13edce28aec3fd7605dce69bd598a67aceda8ba20deda8cabfbd331661b312fd1b5d946cf7a5fadeaa9736e253d4d7da4222eff46dbae46972564ab025ca0528d8a4cc584fcc7554207e184a8e21b4a06dd036473ae580452c860b50321ad872694b6243bc649a4b3d8f630c258b3c7087cddc354aa2405f42040df8d97c2f8b40def7e6bf38479cb007c40eec7ae8e907939054b4ba2f9748b94dd4ad74505b758a265b62050c4fca46be30897fb634f45f588418ca9e6e54fec2e47cd0791d122363d42f81d3d28728d32388761f25a79fa6f87c017b45464d65cc7f717a19f2b20b0b0c3c2741666e26cfa83a5307c4dda04feb89a3d142437d