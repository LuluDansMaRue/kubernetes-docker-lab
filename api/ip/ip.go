package ip

import (
	"net"
	"log"
	"github.com/oschwald/geoip2-golang"
)

func getOutboundsIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetIPLocation() string {
	db, err := geoip2.Open("geolite/GeoLite2-City.mmdb")
  if err != nil {
		log.Fatal(err)
		return "nothing"
  }
	defer db.Close()
	
	localIP := getOutboundsIp()
	record, cErr := db.City(localIP)

	if cErr != nil {
		log.Fatal(err)
		return "nothing"
	}

	return record.Country.IsoCode
}