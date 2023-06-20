#include <arpa/inet.h>
#include <linux/if_ether.h>
#include <linux/if_packet.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <errno.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>

int main() {
    struct  sockaddr_in from;
    from.sin_family = AF_INET;
    uint32_t remote;
    inet_pton(AF_INET, "127.0.0.1", &remote);
    from.sin_addr.s_addr = htonl(remote);
    from.sin_port = htons(30001);

    int fd = socket(AF_PACKET, SOCK_DGRAM, htons(ETH_P_ALL));
    fd = socket(AF_INET, SOCK_RAW, IPPROTO_TCP);
    printf("%d\n", errno);
    // fcntl(fd, F_SETFL, O_NONBLOCK);

	char buffer[65537];
	struct sockaddr_ll src_addr;
	socklen_t src_addr_len = sizeof(src_addr);
    while (1) {
        ssize_t count = recvfrom(fd, buffer, sizeof(buffer), 0, (struct sockaddr*)&src_addr, &src_addr_len);
        if (count == -1) {
            perror("recvfrom");
            exit(1);
        }
        if (count == sizeof(buffer)) {
            fprintf(stderr, "frame too large for buffer: truncated\n");
        } else if (count == 21) {
            char *ip = inet_ntoa(*(struct in_addr*)(buffer+12));
            printf("%s %s\n", ip, buffer+20);
        } else {
            printf("%x %x %x %x", buffer[0], buffer[1], buffer[2], buffer[3]);
        }
    }

    return 0;
}
