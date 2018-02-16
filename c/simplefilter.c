#include <stdio.h>
#include <pcap.h>
#include <netinet/in.h>
#include <netinet/if_ether.h>
#include <inttypes.h>

/* Ethernet addresses are 6 bytes */
#define ETHER_ADDR_LEN	6

	/* Ethernet header */
	struct sniff_ethernet {
		u_char ether_dhost[ETHER_ADDR_LEN]; /* Destination host address */
		u_char ether_shost[ETHER_ADDR_LEN]; /* Source host address */
		u_short ether_type; /* IP? ARP? RARP? etc */
	};

	/* IP header */
	struct sniff_ip {
		u_char ip_vhl;		/* version << 4 | header length >> 2 */
		u_char ip_tos;		/* type of service */
		u_short ip_len;		/* total length */
		u_short ip_id;		/* identification */
		u_short ip_off;		/* fragment offset field */
	#define IP_RF 0x8000		/* reserved fragment flag */
	#define IP_DF 0x4000		/* dont fragment flag */
	#define IP_MF 0x2000		/* more fragments flag */
	#define IP_OFFMASK 0x1fff	/* mask for fragmenting bits */
		u_char ip_ttl;		/* time to live */
		u_char ip_p;		/* protocol */
		u_short ip_sum;		/* checksum */
		struct in_addr ip_src,ip_dst; /* source and dest address */
	};
	#define IP_HL(ip)		(((ip)->ip_vhl) & 0x0f)
	#define IP_V(ip)		(((ip)->ip_vhl) >> 4)

	/* TCP header */
	typedef u_int tcp_seq;

	struct sniff_tcp {
		u_short th_sport;	/* source port */
		u_short th_dport;	/* destination port */
		tcp_seq th_seq;		/* sequence number */
		tcp_seq th_ack;		/* acknowledgement number */
		u_char th_offx2;	/* data offset, rsvd */
	#define TH_OFF(th)	(((th)->th_offx2 & 0xf0) >> 4)
		u_char th_flags;
	#define TH_FIN 0x01
	#define TH_SYN 0x02
	#define TH_RST 0x04
	#define TH_PUSH 0x08
	#define TH_ACK 0x10
	#define TH_URG 0x20
	#define TH_ECE 0x40
	#define TH_CWR 0x80
	#define TH_FLAGS (TH_FIN|TH_SYN|TH_RST|TH_ACK|TH_URG|TH_ECE|TH_CWR)
		u_short th_win;		/* window */
		u_short th_sum;		/* checksum */
		u_short th_urp;		/* urgent pointer */
	};
	/* ethernet headers are always exactly 14 bytes */
#define SIZE_ETHERNET 14

	const struct sniff_ethernet *ethernet; /* The ethernet header */
	const struct sniff_ip *ip; /* The IP header */
	const struct sniff_tcp *tcp; /* The TCP header */
	const char *payload; /* Packet payload */

	u_int size_ip;
	u_int size_tcp;

static void detEndian() 
{
   unsigned int i = 1;
   char *c = (char*)&i;
   if (*c)    
       printf("Little endian\n");
   else
       printf("Big endian\n");
}

int main(int argc, char **argv) {

	char *device; /* Name of device (e.g. eth0, wlan0) */
	pcap_t *handle;		/* Session handle */
	char errbuf[PCAP_ERRBUF_SIZE];	/* Error string */
	struct bpf_program fp;		/* The compiled filter expression */
	char filter_exp[] = "tcp";	/* The filter expression */
	bpf_u_int32 mask;		/* The netmask of our sniffing device */
	bpf_u_int32 net;		/* The IP of our sniffing device */

	device = pcap_lookupdev(errbuf);
    if (device == NULL) {
        printf("Error finding device: %s\n", errbuf);
        return 1;
    }

    printf("Network device found: %s. Sniffing!\n", device);

	if (pcap_lookupnet(device, &net, &mask, errbuf) == -1) {
		fprintf(stderr, "Can't get netmask for device %s\n", device);
		net = 0;
		mask = 0;
	}

	handle = pcap_open_live(device, BUFSIZ, 1, 1000, errbuf);

	 if (handle == NULL) {
		 fprintf(stderr, "Couldn't open device %s: %s\n", device, errbuf);
		 return(2);
	 }
	 if (pcap_compile(handle, &fp, filter_exp, 0, net) == -1) {
		 fprintf(stderr, "Couldn't parse filter %s: %s\n", filter_exp, pcap_geterr(handle));
		 return(2);
	 }
	 if (pcap_setfilter(handle, &fp) == -1) {
		 fprintf(stderr, "Couldn't install filter %s: %s\n", filter_exp, pcap_geterr(handle));
		 return(2);
	 }

	/* Grab a packet */
	const u_char *packet;
	struct pcap_pkthdr header;
	for(int i = 0 ;i<20;i++){
		packet = pcap_next(handle, &header);
		//printf("Jacked a packet with length of [%d]\n", header.len);
		if(header.len>0){
			printf("i = %d Jacked a packet with length of [%d]\n", i,header.len);
			break;
		}
	/* And close the session */
	}
	
	pcap_close(handle);

	ethernet = (struct sniff_ethernet*)(packet);
	ip = (struct sniff_ip*)(packet + SIZE_ETHERNET);
	size_ip = IP_HL(ip)*4;
	if (size_ip < 20) {
		printf("   * Invalid IP header length: %u bytes\n", size_ip);
		return -1;
	}
	tcp = (struct sniff_tcp*)(packet + SIZE_ETHERNET + size_ip);
	size_tcp = TH_OFF(tcp)*4;
	if (size_tcp < 20) {
		printf("   * Invalid TCP header length: %u bytes\n", size_tcp);
		return -1;
	}
	//payload = (u_char *)(packet + SIZE_ETHERNET + size_ip + size_tcp); struct in_addr ip_src,ip_dst; /* source and dest address */
	detEndian();

	printf("TCP src port is %d\n",ntohs(tcp->th_sport));
	printf("TCP dst port is %d\n",ntohs(tcp->th_dport));
	struct in_addr srcIP = ip->ip_src;
	struct in_addr dstIP = ip->ip_dst;

	uint32_t hostSrc = ntohl(srcIP.s_addr);
	uint32_t hostDst = ntohl(dstIP.s_addr);

	printf("%" PRIu32 "\n",hostSrc);
	printf("%" PRIu32 "\n",hostDst);

	printf("Main finished \n");
	return(0);
}

