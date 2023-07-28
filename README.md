# ledwatcher
Watches a led and sends notifications to a MQTT broker when it changes

## Installation
  1. Build the application using ./build.sh
  2. Copy it to /home/pi/pi-light-sensor
  3. Copy the file init.d/pi-light-sensor to /etc/init.d
  4. Edit /etc/init.d/pi-light-sensor and change the notify URL:s to what you need
  5. Execute sudo update-rc.d pi-light-sensor defaults 2345