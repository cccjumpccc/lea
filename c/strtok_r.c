#include <stdio.h>
#include <string.h>

void f1() {
    // char *str = "hello world a b c";
    char str[] = "hello world a b c";
    char *delim = " ";
    // we can also use long x instead of saveptr
    char *saveptr;
    char *token = strtok_r(str, delim, &saveptr);
    while (token) {
        printf("%s\n", token);
        token = strtok_r(NULL, delim, &saveptr);
    }
}

void f2() {
    // char *str = "hello world a b c";
    char str[] = "hello world a b c";
    char *delim = " ";
    // we can also use long x instead of saveptr
    char *saveptr = str;
    char *token;
    do {
        token = strtok_r(NULL, delim, &saveptr);
        printf("%s\n", token);
    } while (*saveptr);
}

int main() {
    f1();
    f2();
}
