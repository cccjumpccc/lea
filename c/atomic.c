#include <stdio.h>

typedef struct dog {
    char name[20];
    int age;
} dog;

void test01() {
    _Atomic dog d1 = {
        .age = 6
    };
    d1.age = 10;
    printf("%d\n", d1.age);
}

void test02() {
    _Atomic int arr[10] = {3,6};
    arr[1] = 9;
    printf("%d %d\n", arr[0], arr[1]);
    char s1[20] = "hello";
    printf("%s\n", s1);
    _Atomic char s2[20] = "hello";
    // _Atomic char s2[20] = {'a', 'b', 'c', 0};
    s2[0] = 'a';
    s2[1] = 'b';
    printf("%s\n", s2);
}

int main() {
    printf("--- test01 ---\n");
    test01();
    printf("--- test02 ---\n");
    test02();
}
