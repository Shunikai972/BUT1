import os
import struct
import socket
def send_packet(dest_addr, data):
    with socket.socket(socket.AF_PACKET, socket.SOCK_RAW, socket.ntohs(17)) as s:
        s.bind((socket.gethostname(), 0))
        s.sendto(data, (dest_addr, 0))
def main():
    # Get the current Wi-Fi network details
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)
    s.sendto(bytes('', socket.MSG_ALLINFOOBJECT), ('<broadcast>', 1))
    data, addr = s.recvfrom(1500)
    ssid, _, _ = struct.unpack('<3s3s', data[:10])
    print("Current Wi-Fi network:", ssid)
    # Get the MAC address of the Wi-Fi router
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)
    s.sendto(bytes('', socket.MSG_ALLINFOOBJECT), ('<broadcast>', 1))
    data, addr = s.recvfrom(1500)
    mac, _, _ = struct.unpack('<6x', data[:6])
    print("Wi-Fi router MAC address:", mac)
    # Send a packet to the Wi-Fi router to deauthenticate the client
    send_packet(mac, bytes('', socket.MSG_DONTWAIT))
    print("Wi-Fi disconnected successfully!")
if __name__ == '__main__':
    main()
