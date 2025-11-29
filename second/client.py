import time
import socket
import sys

# Configuration constants
SERVER_IP = '127.0.0.1'
SERVER_PORT = 12000
NUM_PINGS = 10
TIMEOUT_SEC = 1.0 # 1 second RTT timeout

# CRITICAL DELAY: The single-threaded server.py tends to stall. 
# A 2-second delay is used to give the server ample time to loop and reset.
INTER_PING_DELAY_SEC = 2.0 

def run_udp_pinger_client():
    """
    Sends 10 UDP ping messages to the server, measures RTT, and handles timeouts.
    """
    
    # Create a UDP socket (AF_INET for IPv4, SOCK_DGRAM for UDP)
    try:
        client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    except socket.error as e:
        print(f"Error creating socket: {e}")
        sys.exit(1)

    # Set the socket timeout to 1 second (assignment requirement)
    client_socket.settimeout(TIMEOUT_SEC)
    
    server_address = (SERVER_IP, SERVER_PORT)
    print(f"UDP Pinger Client started. Pinging {SERVER_IP}:{SERVER_PORT}...\n")
    
    packets_received = 0
    total_rtt_ms = 0

    for i in range(1, NUM_PINGS + 1):
        try:
            # 1. Prepare the ping message
            # Format: "Ping <sequence_number> <timestamp_ms>"
            timestamp_ms = int(time.time() * 1000)
            message = f"Ping {i} {timestamp_ms}"
            
            # Record the time just before sending (for RTT calculation)
            start_time = time.time()
            
            # 2. Send the ping message
            # The client does not need to bind a port; the OS assigns a temporary one.
            client_socket.sendto(message.encode(), server_address)
            
            # 3. Wait for the pong reply (up to 1 second due to settimeout)
            # The buffer size must be large enough (e.g., 1024 bytes)
            data, server = client_socket.recvfrom(1024)
            
            # 4. Calculate RTT
            end_time = time.time()
            rtt_sec = end_time - start_time
            rtt_ms = int(rtt_sec * 1000)
            
            # Print success message
            response_message = data.decode()
            print(f"PING {i}: Received '{response_message}', RTT: {rtt_ms} ms")
            
            packets_received += 1
            total_rtt_ms += rtt_ms

        except socket.timeout:
            # Timeout error (no reply received within 1 second)
            print(f"PING {i}: Request timed out")
        
        except socket.error as e:
            # Other socket errors (e.g., connection reset)
            print(f"PING {i}: Socket error: {e}")

        # Wait the designated delay before sending the next ping (CRITICAL for unstable server)
        time.sleep(INTER_PING_DELAY_SEC)

    client_socket.close()

    # Final summary calculation
    if packets_received > 0:
        avg_rtt = total_rtt_ms / packets_received
        print(f"\nUDP Pinger Client finished.")
        print(f"Summary: Sent {NUM_PINGS}, Received {packets_received}, Lost {NUM_PINGS - packets_received}")
        print(f"Average RTT: {avg_rtt:.2f} ms")
    else:
        print("\nUDP Pinger Client finished. No packets received.")


if __name__ == '__main__':
    run_udp_pinger_client()