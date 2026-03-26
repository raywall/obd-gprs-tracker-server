# OBD Tracker Server (GPRS)

This is a POC project to create a Go Server capable of receive data from OBD Tracker devices, exploring the possibilities of capture additional data of the car too.


## 1. Device Configuration (SMS)

To make the tracker talk with your Go server, you must send specific SMS commands to the SIM card used by the tracker.

> Note: Most trackers use `000000` or `123456` as the default password.


### Step 01: Set the APN (internet access)  

The tracker needs to know how to access the cellular data network.

- Command: `APN123456 <your_carrier_apn>`
- _Example (T-Mobile):_ `APN123456 fast.t-mobile.com`


### Step 2: Set the Server IP and Port

This tells the tracker where your Go server is located.

- Command: `adminip123456 <your_public_ip> 5001`
- _Example:_ `adminip123456 123.45.67.89 5001`


### Step 3: Set the Data Upload Interval

Tell it how often to send data (in seconds).

- Command: `upload123456 30` (Sends data every 30 seconds)


### Step 4: Switch Configuration to GPRS Mode

Ensure it is sending via data, not just SMS.

- Command: `GPRS123456`


## 2. Critical Requirements

1. **Public IP**
    Your computer running Go must have a Public IP address. If you are at home, you must set up Port Forwarding on your router to send traffic from port `5001` to your local computer's IP.
2. **Firewall**
    Ensure port `5001` is open in your Windows/Linux firewall.
3. **Protocol Manual**
    Every brand uses slightly different hex headers. Once you see the raw hex printing in your Go terminal, search for the "GPRS Protocol Document" for your model to decode the specific bytes for RPM, Fuel, etc.

> Depending of the brand or model name the communication can need an additional and specific hex-map to capture car data
# obd-gprs-tracker-server
