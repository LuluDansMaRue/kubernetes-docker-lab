package ip

import (
	"net"
	"log"
	"errors"
	"github.com/oschwald/geoip2-golang"
)

// Get Outbounds IP 
// Retrieve the outbound IP (should be the public ip)
// Return net.IP
func getOutboundsIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Get IP Location
// Retrieve the IP location based on the GeoLite2 city database
// Return string (location)
// Return errors
func GetIPLocation() (string, error) {
	db, err := geoip2.Open("geolite/GeoLite2-City.mmdb")
  if err != nil {
		log.Fatal(err)
		return "", errors.New("Unable to access to GeoLite database")
  }
	defer db.Close()
	
	localIP := getOutboundsIp()
	if localIP == nil {
		return "", errors.New("unable to get ip")
	}

	record, cErr := db.City(localIP)

	if cErr != nil {
		log.Fatal(err)
		return "", errors.New("Unable to retrieve city")
	}

	return record.Country.IsoCode, nil
}