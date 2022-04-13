#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdint.h>


#include "py_gen.h"

unsigned char DECRYPTED[sizeof(SFLAG)];

//-------------------------------Galois Field math

#define PRIM_POLY  0x1002d
#define GF_FIELD_WIDTH 16

uint32_t gf_shift_multiply (uint32_t a, uint32_t b)
{
    uint32_t product, i, pp;
    pp = PRIM_POLY;
    product = 0;

    for (i = 0; i < GF_FIELD_WIDTH; i++) {
        if (a & (1 << i)) product ^= (b << i);
    }
    for (i = (GF_FIELD_WIDTH*2-2); i >= GF_FIELD_WIDTH; i--) {
        if (product & (1 << i)) product ^= (pp << (i-GF_FIELD_WIDTH));
    }
    return product;
}


uint32_t gf_euclid (uint32_t b)
{
    uint32_t e_i, e_im1, e_ip1;
    uint32_t d_i, d_im1, d_ip1;
    uint32_t y_i, y_im1, y_ip1;
    uint32_t c_i;

    if (b == 0) return -1;
    e_im1 = PRIM_POLY;
    e_i = b;
    d_im1 = GF_FIELD_WIDTH;
    for (d_i = d_im1; ((1 << d_i) & e_i) == 0; d_i--) ;
    y_i = 1;
    y_im1 = 0;

    while (e_i != 1) {

        e_ip1 = e_im1;
        d_ip1 = d_im1;
        c_i = 0;

        while (d_ip1 >= d_i) {
            c_i ^= (1 << (d_ip1 - d_i));
            e_ip1 ^= (e_i << (d_ip1 - d_i));
            if (e_ip1 == 0) return 0;
            while ((e_ip1 & (1 << d_ip1)) == 0) d_ip1--;
        }

        y_ip1 = y_im1 ^ gf_shift_multiply(c_i, y_i);
        y_im1 = y_i;
        y_i = y_ip1;

        e_im1 = e_i;
        d_im1 = d_i;
        e_i = e_ip1;
        d_i = d_ip1;
    }
    return y_i;
}

uint32_t gf_add(uint32_t a, uint32_t b)
{
    return a ^ b;
}

uint32_t gf_sub(uint32_t a, uint32_t b)
{
    return a ^ b;
}

uint32_t gf_div(uint32_t a, uint32_t b)
{
    uint32_t bi = gf_euclid (b);
    return gf_shift_multiply(a, bi);
}

int decrypt(char * sflag, char * decrypt, int len, uint32_t pincode)
{
    int i;
    int crc = CRC;
    for  (i=0; i<len; i+=2)
    {
        *(uint16_t *) decrypt = TRANSFORM ( *(uint16_t *) sflag, pincode ) ;
        crc = crc ^ *(uint16_t *)decrypt;
        decrypt += 2;
        sflag += 2;
    }
    return crc;
}

int main()
{
    int i = 0;
    uint16_t pincode =0;
    int crc;

    printf("Enter PIN-CODE:\n");
    scanf("%x",&pincode);
    printf("Decrypting\n");
    for (i=0; i<5; i++)
    {
        printf("+");
        fflush(stdout);
        sleep(1);
    }
    crc = decrypt(SFLAG, DECRYPTED, sizeof (SFLAG), pincode);

    if (0x00 == crc)
    {
        printf("\n\nSuccess:\n %s", DECRYPTED);
        return 0;
    }
    else{
        printf("\n\nWrong PIN-CODE\n");
        return -1;
    }
}

