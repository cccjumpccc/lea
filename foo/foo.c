#include <stdio.h>

// TODO: x放在了初始化为0的那个区???
int main() {
    int x;
    switch (888) { 
        x = 4; 
        case 0: 
        printf("case 0\n"); 
        default: 
        printf("case default\n"); 
        printf("%d\n", x); 
    } 
}
