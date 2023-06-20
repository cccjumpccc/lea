#include <stdio.h>
#include <stdatomic.h>

int main() {
    _Atomic __uint8_t x = 255;
    _Atomic unsigned short a = 65535;
    signed short b = (signed short)a;
    printf("%ld\n", sizeof(unsigned short));
    printf("%d\n", a);
    printf("%u\n", a);
    printf("%d\n", b);
    printf("%u\n", b);
}
