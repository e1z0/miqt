package main

import (
	"log"
	"os"

	qt "github.com/e1z0/miqt/qt6"
	"github.com/e1z0/miqt/qt6/network"
)

func main() {

	qt.NewQApplication(os.Args)

	log.Printf("Looking up DNS info, please wait...")

	dns := network.NewQDnsLookup2(network.QDnsLookup__A, "google.com")
	dns.OnFinished(func() {
		dns.DeleteLater()

		if dns.Error() != network.QDnsLookup__NoError {
			log.Printf("An error occurred: %v", dns.ErrorString())
			return
		}

		results := dns.HostAddressRecords() // CanonicalNameRecords, TextRecords, ServiceRecords, ...
		log.Printf("Found %d result(s).", len(results))

		for _, record := range results {
			log.Printf("- %s", record.Value().ToString())
		}

		qt.QCoreApplication_Exit()
	})
	dns.Lookup()

	qt.QApplication_Exec()
}
