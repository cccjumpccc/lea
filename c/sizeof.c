#include <stdio.h>

int main() {
    char s1[] = "hello";
    printf("%d\n", sizeof(s1));
    printf("%d\n", sizeof(sizeof(s1)));
}
