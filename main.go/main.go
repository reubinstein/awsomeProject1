package main_go

import (
	"net/http"

	"net/http"

	// Import the different modules and their APIs
	"myapp/analytics"
	"myapp/compliance"
	"myapp/constituent"
	"myapp/crm"
	"myapp/donor"
	"myapp/election"
	"myapp/event"
	"myapp/financial"
	"myapp/fundraising"
	"myapp/inventory"
	"myapp/media"
	"myapp/membership"
	"myapp/policy"
	"myapp/program"
	"myapp/volunteer"
)

func main() {
	// Initialize the different modules and their APIs
	membershipAPI := membership.NewAPI()
	electionAPI := election.NewAPI()
	fundraisingAPI := fundraising.NewAPI()
	programAPI := program.NewAPI()
	mediaAPI := media.NewAPI()
	donorAPI := donor.NewAPI()
	eventAPI := event.NewAPI()
	crmAPI := crm.NewAPI()
	policyAPI := policy.NewAPI()
	financialAPI := financial.NewAPI()
	complianceAPI := compliance.NewAPI()
	inventoryAPI := inventory.NewAPI()
	volunteerAPI := volunteer.NewAPI()
	constituentAPI := constituent.NewAPI()
	communicationAPI := communication.NewAPI()
	analyticsAPI := analytics.NewAPI()

	// Set up the routes for each API
	http.HandleFunc("/membership/", membershipAPI.HandleRequest)
	http.HandleFunc("/election/", electionAPI.HandleRequest)
	http.HandleFunc("/fundraising/", fundraisingAPI.HandleRequest)
	http.HandleFunc("/program/", programAPI.HandleRequest)
	http.HandleFunc("/media/", mediaAPI.HandleRequest)
	http.HandleFunc("/donor/", donorAPI.HandleRequest)
	http.HandleFunc("/event/", eventAPI.HandleRequest)
	http.HandleFunc("/crm/", crmAPI.HandleRequest)
	http.HandleFunc("/policy/", policyAPI.HandleRequest)
	http.HandleFunc("/financial/", financialAPI.HandleRequest)
	http.HandleFunc("/compliance/", complianceAPI.HandleRequest)
	http.HandleFunc("/inventory/", inventoryAPI.HandleRequest)
	http.HandleFunc("/volunteer/", volunteerAPI.HandleRequest)
	http.HandleFunc("/constituent/", constituentAPI.HandleRequest)
	http.HandleFunc("/communication/", communicationAPI.HandleRequest)
	http.HandleFunc("/analytics/", analyticsAPI.HandleRequest)

	// Start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
