package main

import (
	"fmt"
	"sort"
)

type Appointment struct {
	AppointmentID int
	Date          string
	VisitingHours int
	Fees          int
	DoctorID      int
}

type Patient struct {
	Name         string
	Appointments []Appointment
}

// Sample Data
type HospitalData struct {
	Patients []Patient
}

func DisplayPatientAppointmentHoursSummary(data HospitalData) {
	summary := make(map[string]int)
	for _, patient := range data.Patients {
		total := 0
		for _, appt := range patient.Appointments {
			total += appt.VisitingHours
		}
		summary[patient.Name] = total
	}

	for name, hours := range summary {
		fmt.Printf("%s — %d\n", name, hours)
	}
}

func FindPatientWithMaxAppointmentHours(data HospitalData) {
	maxHours := -1
	var patientName string

	for _, patient := range data.Patients {
		total := 0
		for _, appt := range patient.Appointments {
			total += appt.VisitingHours
		}
		if total > maxHours {
			maxHours = total
			patientName = patient.Name
		}
	}

	fmt.Printf("%s has taken maximum number of appointments — %d\n", patientName, maxHours)
}

func FetchDoctorWithMaxAppointment(data HospitalData) {
	doctorCount := make(map[int]int)

	for _, patient := range data.Patients {
		for _, appt := range patient.Appointments {
			doctorCount[appt.DoctorID]++
		}
	}

	maxAppt := -1
	var maxDoctorID int
	for id, count := range doctorCount {
		if count > maxAppt {
			maxAppt = count
			maxDoctorID = id
		}
	}

	fmt.Printf("Doctor with maximum number of appointments — %d\n", maxDoctorID)
}

func GroupByAppointmentDate(data HospitalData) {
	grouped := make(map[string][]Appointment)

	for _, patient := range data.Patients {
		for _, appt := range patient.Appointments {
			grouped[appt.Date] = append(grouped[appt.Date], appt)
		}
	}

	sortedDates := make([]string, 0, len(grouped))
	for date := range grouped {
		sortedDates = append(sortedDates, date)
	}
	sort.Strings(sortedDates)

	for _, date := range sortedDates {
		fmt.Printf("%s %v\n", date, grouped[date])
	}
}

func main() {
	data := HospitalData{
		Patients: []Patient{
			{
				Name: "Emily",
				Appointments: []Appointment{
					{AppointmentID: 1, Date: "2023-12-01", VisitingHours: 2, Fees: 50, DoctorID: 123},
					{AppointmentID: 2, Date: "2023-12-05", VisitingHours: 1, Fees: 30, DoctorID: 124},
					{AppointmentID: 3, Date: "2023-12-10", VisitingHours: 3, Fees: 75, DoctorID: 128},
				},
			},
			{
				Name: "Daniel",
				Appointments: []Appointment{
					{AppointmentID: 4, Date: "2023-12-01", VisitingHours: 1, Fees: 25, DoctorID: 125},
					{AppointmentID: 6, Date: "2023-12-10", VisitingHours: 1, Fees: 30, DoctorID: 123},
				},
			},
		},
	}

	fmt.Println("--- Patient Appointment Hours Summary ---")
	DisplayPatientAppointmentHoursSummary(data)

	fmt.Println("\n--- Patient with Max Appointment Hours ---")
	FindPatientWithMaxAppointmentHours(data)

	fmt.Println("\n--- Doctor with Max Appointments ---")
	FetchDoctorWithMaxAppointment(data)

	fmt.Println("\n--- Grouped Appointments by Date ---")
	GroupByAppointmentDate(data)
}
