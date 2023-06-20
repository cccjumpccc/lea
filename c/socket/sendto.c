#include <arpa/inet.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <errno.h>
#include <stdio.h>

struct _pkt
{
    struct iphdr ip;
    char data;
} pkt;

int main() {
    struct  sockaddr_in to;
    to.sin_family = AF_INET;
    uint32_t remote_ip;
    inet_pton(AF_INET, "127.0.0.1", &remote_ip);
    to.sin_addr.s_addr = htonl(remote_ip);
    to.sin_port = htons(30001);

    pkt.ip.version = 4;
    pkt.ip.ihl = 5;
    pkt.ip.tot_len = 21;
    pkt.ip.id = htons(0x455);
    pkt.ip.ttl = 255;
    pkt.ip.protocol = 99;
    pkt.ip.daddr = remote_ip;
    pkt.ip.frag_off = htons (8190); 
    pkt.data = 'a';
    size_t psize = sizeof(struct iphdr)+1;

    int fd = socket(AF_INET, SOCK_RAW, IPPROTO_RAW);
    size_t count = sendto(fd, &pkt, psize, 0, (struct sockaddr*)&to, sizeof(struct sockaddr));
    printf("%d %ld\n", errno, count);
    return 0;
}
