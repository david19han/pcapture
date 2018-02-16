/* Compile with: gcc find_device.c -lpcap */
#include <stdio.h>
#include <pcap.h>

int main(int argc, char **argv) {
    char *device; /* Name of device (e.g. eth0, wlan0) */
    char error_buffer[PCAP_ERRBUF_SIZE]; /* Size defined in pcap.h */

    /* Find a device */
    device = pcap_lookupdev(error_buffer);
    if (device == NULL) {
        printf("Error finding device: %s\n", error_buffer);
        return 1;
    }

    printf("Network device found: %s. Sniffing!\n", device);

    pcap_t *handle;
    handle = pcap_open_live(device,BUFSIZ,1,1000,error_buffer);
    if(handle == NULL){
        fprintf(stderr,"error\n");
        return 2;
    }

    return 0;
}